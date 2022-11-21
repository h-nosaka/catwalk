package base

import (
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"strconv"

	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type Meta struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	Count   int `json:"count"`
	Current int `json:"current"`
}

type Response struct {
	Meta   *Meta        `json:"meta"`
	Result *interface{} `json:"result"`
}

type ArrayResponse struct {
	Meta    *Meta          `json:"meta"`
	Results *[]interface{} `json:"results"`
}

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Meta    *Meta    `json:"meta"`
	Message string   `json:"message"`
	Errors  *[]Error `json:"errors"`
}

const JsonError = "{}"

func SetDefaultHeaders(f *fiber.Ctx) {
	f.Response().Header.Add("Content-Type", "application/json")
}

func ApiResult(f *fiber.Ctx, code int, result interface{}) error {
	body := ToJson(Response{
		Meta:   &Meta{Total: 1, Page: 1, Count: 1, Current: 1},
		Result: &result,
	}, JsonError)
	SetDefaultHeaders(f)
	return f.Status(code).SendString(string(body))
}

func ApiResults(f *fiber.Ctx, code int, results []interface{}, total int) error {
	page, err := strconv.Atoi(f.Query("page", "1"))
	if err != nil {
		page = 1
	}
	start, err := strconv.Atoi(f.Query("start", "1"))
	if err != nil {
		start = 1
	}
	body := ToJson(ArrayResponse{
		Meta: &Meta{
			Total:   total,
			Page:    int(math.Ceil(float64(total) / float64(page))),
			Count:   len(results),
			Current: start,
		},
		Results: &results,
	}, JsonError)
	SetDefaultHeaders(f)
	return f.Status(code).SendString(body)
}

func ErrorApi(f *fiber.Ctx, code int, errors ...Error) error {
	body, err := json.Marshal(&ErrorResponse{
		Meta:   &Meta{Total: 1, Page: 1, Count: 1, Current: 1},
		Errors: &errors,
	})
	if err != nil {
		return ErrorCodeApi(f, 500, err)
	}
	SetDefaultHeaders(f)
	return f.Status(code).SendString(string(body))
}

func ErrorCodeApi(f *fiber.Ctx, code int, detail error) error {
	if detail != nil {
		Log.Error(fmt.Sprintf("api error: %v", detail))
	}
	body := ToJson(ErrorResponse{
		Meta: &Meta{Total: 1, Page: 1, Count: 1, Current: 1},
	}, JsonError)
	SetDefaultHeaders(f)
	return f.Status(code).SendString(body)
}

func ValidateRegexp(fl validator.FieldLevel) bool {
	r := regexp.MustCompile(fl.Param())
	keys := fl.Field()
	switch keys.Type().String() {
	case "string":
		return r.MatchString(fl.Field().String())
	case "[]string":
		ok := true
		for _, key := range keys.Interface().([]string) {
			if !r.MatchString(key) {
				ok = false
				break
			}
		}
		return ok
	}
	return false
}

func ValidateNew() *validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("regexp", ValidateRegexp); err != nil {
		Log.Error(err.Error())
		return validate
	}
	return validate
}

func ValidateError(err error) ([]Error, error) {
	if err != nil {
		rs := []Error{}
		for _, item := range err.(validator.ValidationErrors) {
			rs = append(rs, Error{
				Field:   item.Field(),
				Message: item.Param(),
			})
		}
		return rs, err
	}
	return nil, err
}

func RequestParser(f *fiber.Ctx, request interface{}) interface{} {
	Log.Info(fmt.Sprintf("uri: %s, body: %s", f.Request().RequestURI(), f.Request().Body()), zap.String("request_id", f.Locals("requestid").(string)))
	if string(f.Request().Header.Method()) == fasthttp.MethodGet {
		if err := f.QueryParser(request); err != nil {
			Log.Error("request error: " + err.Error())
			return err
		}
	} else {
		if err := f.BodyParser(request); err != nil {
			Log.Error(fmt.Sprintf("request error: %s", err.Error()), zap.String("request_id", f.Locals("requestid").(string)))
			return err
		}
	}
	result := ValidateNew().Struct(request)
	if result != nil {
		if errors, err := ValidateError(result.(validator.ValidationErrors)); err != nil {
			Log.Error(fmt.Sprintf("validation error: %s", err.Error()), zap.String("request_id", f.Locals("requestid").(string)))
			return errors
		}
	}
	return nil
}

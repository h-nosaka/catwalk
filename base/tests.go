package base

import (
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"testing"

	miniredis "github.com/alicebob/miniredis/v2"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
)

var MiniRedis *miniredis.Miniredis

func Run(t *testing.T, it string, test func()) {
	// Redis初期化
	if MiniRedis != nil {
		MiniRedis.Close()
	}
	MiniRedis = miniredis.RunT(t)
	os.Setenv("REDIS_URL", MiniRedis.Addr())
	RedisInit()
	It(t, it)
	test()
}

func It(t *testing.T, it string) {
	_, file, line, _ := runtime.Caller(2) // Run or CreateApiTest内から呼ばれた場合は2
	if strings.Contains(file, "tests.go") {
		_, file, line, _ = runtime.Caller(1) // 直接CreateTestを呼んだ場合の呼び出し元は1
	}
	t.Logf("%s:%d: %s", file, line, it)
}

func CreateApiTest(t *testing.T, it string, routes func(*fiber.App)) *apitest.APITest {
	It(t, it)
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	routes(app)
	// デバグ表示するならこちら
	// return apitest.New().Debug().HandlerFunc(FiberToHandlerFunc(app))
	// CI上で動かすならこちら
	return apitest.New().HandlerFunc(FiberToHandlerFunc(app))
}

const MessageEqaul = "%s:%d: got: %v, want: %v"
const MessageNotEqaul = "%s:%d: got: %v, not_want: %v"
const MessageError = "error: %v"

func AssertEqaul(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf(MessageEqaul, file, line, got, want)
	}
}

func AssertNotEqaul(t *testing.T, got interface{}, want interface{}) {
	if got == want {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf(MessageNotEqaul, file, line, got, want)
	}
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r, 10000) // 10秒以内にレスポンスを返すこと
		if err != nil {
			panic(err)
		}

		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

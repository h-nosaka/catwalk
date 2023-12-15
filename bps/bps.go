package bps

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/h-nosaka/catwalk/money"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Bps struct {
	Raw      *big.Int              // ベーシズポイント化された実数
	Digit    *big.Int              // ベーシスポイントの乗算数
	DigitRaw DigitType             // ベーシスポイントの桁数
	Round    func(float64) float64 // 小数点の丸め用関数
	Error    error                 // 最後に発生したエラー
}

type DigitType int

const (
	DigitBps DigitType = 4
	DigitPpm DigitType = 6
	DigitPpb DigitType = 9
	DigitPpt DigitType = 12
	DigitPpq DigitType = 15
)

type BpsFormatOption int

const (
	OptionInt64 BpsFormatOption = iota
	OptionFloatPercent
	OptionFloatPerMille
	OptionFloatBasisPoint
)

var Digit = DigitBps
var Round = math.Round
var defaultDigit = DigitBps
var defaultRound = math.Round

// エラー定義
var (
	ErrorInvalidDataSource          = errors.New("invalid data source")
	ErrorInvalidArithmeticOperators = errors.New("invalid arithmetic operators")
	ErrorInvalidSourceIntegerZero   = errors.New("integer is zero")
	ErrorNilValue                   = errors.New("value is nil")
	ErrorUndefinedLanguage          = errors.New("undefined language type")
	ErrorUndefinedCurrency          = errors.New("undefined currency type")
	ErrorUndefinedCurrencyFormat    = errors.New("undefined currency format")
)

func Init(digit DigitType, round func(float64) float64) {
	Digit = digit
	Round = round
}

func Reset() {
	Digit = defaultDigit
	Round = defaultRound
}

func (p DigitType) Int() int {
	return int(p)
}

func (p DigitType) Int64() int64 {
	return int64(p)
}

func New(src interface{}) *Bps {
	rs := &Bps{
		Raw:      big.NewInt(0),
		Digit:    big.NewInt(10),
		DigitRaw: Digit,
		Round:    Round,
		Error:    nil,
	}
	rs.Digit.Exp(rs.Digit, big.NewInt(Digit.Int64()), nil)
	return rs.Set(src)
}

func (p *Bps) Set(src interface{}) *Bps {
	switch v := src.(type) {
	case uint:
		return New(int64(v))
	case uint8:
		return New(int64(v))
	case uint16:
		return New(int64(v))
	case uint32:
		return New(int64(v))
	case uint64:
		return New(int64(v))
	case int:
		return New(int64(v))
	case int8:
		return New(int64(v))
	case int16:
		return New(int64(v))
	case int32:
		return New(int64(v))
	case int64:
		p.Raw.Mul(big.NewInt(v), p.Digit)
	case float32:
		return New(float64(v))
	case float64:
		raw := big.NewFloat(v)
		raw.Mul(raw, big.NewFloat(float64(p.Digit.Int64())))
		i, acc := raw.Int64()
		f := 0.0
		if acc != big.Exact {
			a := big.NewFloat(0)
			b, _ := a.Mul(raw, big.NewFloat(10)).Int64()
			f = float64(b%10) / 10
		}
		if p.Round(f) >= 1.0 {
			i += 1
		}
		p.Raw = big.NewInt(i)
	case string:
		val := strings.Split(v, ".")
		if len(v) > 0 {
			i, err := strconv.ParseInt(val[0], 10, 64)
			if err != nil {
				p.Error = err
			}
			p.Raw.Mul(big.NewInt(i), p.Digit)
		}
		if len(val) > 1 {
			s := val[1]
			add := false
			for len(s) < int(p.DigitRaw) {
				s += "0"
			}
			if len(val[1]) > int(p.DigitRaw) {
				s = val[1][0:p.DigitRaw]
				a, err := strconv.Atoi(val[1][p.DigitRaw : p.DigitRaw+1])
				if err != nil {
					p.Error = err
				}
				if p.Round(float64(a)/10) >= 1.0 {
					add = true
				}
			}
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				p.Error = err
			}
			if add {
				i++
			}
			p.Raw.Add(p.Raw, big.NewInt(i))
		}
	case Bps:
		return &v
	case *Bps:
		return v
	default:
		p.Error = ErrorInvalidDataSource
	}
	return p
}

func (p Bps) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Bps) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = *New(value)
	return p.Error
}

func (p Bps) Value() (driver.Value, error) {
	if p.Raw == nil {
		p = *New(0)
	}
	return p.String(), p.Error
}

func (p *Bps) Scan(data interface{}) error {
	fmt.Println(data)
	*p = *New(data)
	return p.Error
}

func (p *Bps) SetRound(round func(float64) float64) *Bps {
	p.Round = round
	return p
}

func (p *Bps) ChangeDigit(digit DigitType) *Bps {
	diff := big.NewInt(10)
	if digit > p.DigitRaw {
		diff.Exp(diff, big.NewInt(digit.Int64()-p.DigitRaw.Int64()), nil)
		p.Raw.Mul(p.Raw, diff)
		p.Digit.Exp(big.NewInt(10), big.NewInt(digit.Int64()), nil)
		p.DigitRaw = digit
	} else {
		src := p.String()
		p.Digit.Exp(big.NewInt(10), big.NewInt(digit.Int64()), nil)
		p.DigitRaw = digit
		p.Set(src)
	}
	return p
}

func (p *Bps) Int64() int64 {
	add := int64(0)
	if p.Round(float64(big.NewInt(0).Div(p.Raw, big.NewInt(p.Digit.Int64()/10)).Int64()%10)/10) >= 1.0 {
		add = 1
	}
	return big.NewInt(0).Div(p.Raw, p.Digit).Int64() + add
}

func (p *Bps) Decimal(num int) string {
	v := p.Raw.Int64() % p.Digit.Int64()
	if num < p.DigitRaw.Int() {
		r := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(p.DigitRaw.Int()-num)), nil)
		add := int64(0)
		fmt.Println(v, r.Int64(), v%r.Int64())
		if p.Round(float64(v%r.Int64())/float64(r.Int64())) == 1.0 {
			add = 1
		}
		v = v/r.Int64() + add
	}
	f := fmt.Sprintf("%d", v)
	for len(f) < num {
		f += "0"
	}
	return f[0:num]
}

func (p *Bps) String() string {
	f := p.Raw.Int64() % p.Digit.Int64()
	return fmt.Sprintf("%d.%d", big.NewInt(0).Div(p.Raw, p.Digit).Int64(), f)
}

func (p *Bps) FloatString(num int) string {
	return fmt.Sprintf("%d.%s", big.NewInt(0).Div(p.Raw, p.Digit).Int64(), p.Decimal(num))
}

func (p *Bps) Float64() float64 {
	rs, err := strconv.ParseFloat(p.String(), 64)
	if err != nil {
		p.Error = err
	}
	return rs
}

func (p *Bps) Calc(arithmetic string, src interface{}) *Bps {
	val := New(src)
	switch arithmetic {
	case "+": // 加算
		p.Raw.Add(p.Raw, val.Raw)
	case "-": // 減算
		p.Raw.Sub(p.Raw, val.Raw)
	case "%": // 余り
		p.Raw.Mod(p.Raw, val.Raw)
	case "*", "/", "^":
		if val.Int64() == 0 {
			p.Error = ErrorInvalidSourceIntegerZero
			return p
		}
		switch arithmetic {
		case "*": // 乗算
			p.Raw.Mul(p.Raw, big.NewInt(val.Int64()))
		case "/": // 除算
			p.Raw.Div(p.Raw, big.NewInt(val.Int64()))
		case "^": // 累乗
			v := big.NewInt(p.Int64())
			v.Exp(v, big.NewInt(val.Int64()), nil)
			p.Set(v.Int64())
		}
	default:
		p.Error = ErrorInvalidArithmeticOperators
	}
	return p
}

func (p *Bps) Format(lang *money.Language, options ...BpsFormatOption) string {
	num := big.NewInt(0).Div(p.Raw, p.Digit).Int64()
	rs := ""
	dec := ""
	sep := "."
	useFloat := true
	for _, option := range options {
		switch option {
		case OptionInt64:
			num = p.Int64()
			useFloat = false
		case OptionFloatPercent:
			dec = p.Decimal(2)
		case OptionFloatPerMille:
			dec = p.Decimal(3)
		case OptionFloatBasisPoint:
			dec = p.Decimal(4)
		}
	}
	if lang != nil {
		msg := message.NewPrinter(language.MustParse(lang.String()))
		rs = msg.Sprintf("%d", num)
		sep = money.GetLanguage(*lang).Decimal
	}
	if len(rs) == 0 {
		rs = fmt.Sprintf("%d", num)
	}
	if useFloat {
		if len(dec) == 0 {
			dec = p.Decimal(p.DigitRaw.Int())
		}
		rs += fmt.Sprintf("%s%s", sep, dec)
	}
	return rs
}

func (p *Bps) Currency(currency money.Currency, lang money.Language) string {
	options := []BpsFormatOption{}
	c := money.GetCurrency(currency)
	if c == nil {
		p.Error = ErrorUndefinedCurrency
		return p.Format(&lang, options...)
	}
	l := money.GetLanguage(lang)
	if l == nil {
		p.Error = ErrorUndefinedLanguage
		return p.Format(&lang, options...)
	}
	switch c.Fraction {
	case 0:
		options = append(options, OptionInt64)
	case 2:
		options = append(options, OptionFloatPercent)
	case 3:
		options = append(options, OptionFloatPerMille)
	case 4:
		options = append(options, OptionFloatBasisPoint)
	default:
		p.Error = ErrorUndefinedCurrencyFormat
	}
	switch l.Template {
	case "%s %d":
		return fmt.Sprintf("%s %s", c.Grapheme, p.Format(&lang, options...))
	case "%d %s":
		return fmt.Sprintf("%s %s", p.Format(&lang, options...), c.Grapheme)
	case "%d%s":
		return fmt.Sprintf("%s%s", p.Format(&lang, options...), c.Grapheme)
	}
	return fmt.Sprintf("%s%s", c.Grapheme, p.Format(&lang, options...))
}

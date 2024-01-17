package bps_test

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/h-nosaka/catwalk/bps"
	"github.com/h-nosaka/catwalk/money"
)

func TestNewBps(t *testing.T) {
	if val := bps.New(1234); val.Raw.Int64() != 12340000 {
		t.Error("int64: 期待値と異なる", val.Raw.Int64())
	}
	if val := bps.New(1000.12345); val.Raw.Int64() != 10001235 {
		t.Error("float64: 期待値と異なる", val.Raw.Int64())
	}
	if val := bps.New(1000.1234); val.Raw.Int64() != 10001234 {
		t.Error("float64: 期待値と異なる", val.Raw.Int64())
	}
	if val := bps.New("1000.12345"); val.Raw.Int64() != 10001235 {
		t.Error("string: 期待値と異なる", val.Raw.Int64())
	}
	if val := bps.New("1000.12344"); val.Raw.Int64() != 10001234 {
		t.Error("string: 期待値と異なる", val.Raw.Int64())
	}
	if val := bps.New("1000.12"); val.Raw.Int64() != 10001200 {
		t.Error("string: 期待値と異なる", val.Raw.Int64())
	}
	if val := bps.New(""); val.Raw.Int64() != 0 {
		t.Error("string: 期待値と異なる", val.Raw.Int64())
	}
}

func TestBpsInt64(t *testing.T) {
	if val := bps.New(1234.5); val.Int64() != 1235 {
		t.Error("int64: 期待値と異なる", val.Int64())
	}
	if val := bps.New(1234.4); val.Int64() != 1234 {
		t.Error("int64: 期待値と異なる", val.Int64())
	}
}

func TestBpsString(t *testing.T) {
	if val := bps.New(1234.5678); val.String() != "1234.5678" {
		t.Error("string: 期待値と異なる", val.String())
	}
	if val := bps.New(1234.56785); val.String() != "1234.5679" {
		t.Error("string: 期待値と異なる", val.String())
	}
	if val := bps.New(1234.56784); val.String() != "1234.5678" {
		t.Error("string: 期待値と異なる", val.String())
	}
}

func TestBpsFloatString(t *testing.T) {
	if val := bps.New(1234.5678); val.FloatString(4) != "1234.5678" {
		t.Error("string: 期待値と異なる", val.FloatString(4))
	}
	if val := bps.New(1234.5678); val.FloatString(3) != "1234.568" {
		t.Error("string: 期待値と異なる", val.FloatString(3))
	}
	if val := bps.New(1234.564); val.FloatString(2) != "1234.56" {
		t.Error("string: 期待値と異なる", val.FloatString(2))
	}
	if val := bps.New(1234.56784); val.FloatString(10) != "1234.5678000000" {
		t.Error("string: 期待値と異なる", val.FloatString(10))
	}
	if val := bps.New(1234.56785); val.FloatString(10) != "1234.5679000000" {
		t.Error("string: 期待値と異なる", val.FloatString(10))
	}
}

func TestBpsFloat64(t *testing.T) {
	if val := bps.New(1234.5678); val.Float64() != 1234.5678 {
		t.Error("string: 期待値と異なる", val.Float64())
	}
	if val := bps.New(1234.56785); val.Float64() != 1234.5679 {
		t.Error("string: 期待値と異なる", val.Float64())
	}
	if val := bps.New(1234.56784); val.Float64() != 1234.5678 {
		t.Error("string: 期待値と異なる", val.Float64())
	}
}

func TestBpsBps(t *testing.T) {
	if val := bps.New(bps.New(1234.5678)); val.String() != "1234.5678" {
		t.Error("bps: 期待値と異なる", val.Float64())
	}
	if val := bps.New(*bps.New(1234.5678)); val.String() != "1234.5678" {
		t.Error("bps: 期待値と異なる", val.Float64())
	}
}

func TestBpsSetRound(t *testing.T) {
	if val := bps.New(1234.5670).SetRound(math.Ceil).FloatString(3); val != "1234.567" {
		t.Error("round: 期待値と異なる", val)
	}
	if val := bps.New(1234.5675).SetRound(math.Ceil).FloatString(3); val != "1234.568" {
		t.Error("round: 期待値と異なる", val)
	}
	if val := bps.New(1234.5671).SetRound(math.Ceil).FloatString(3); val != "1234.568" {
		t.Error("round: 期待値と異なる", val)
	}
	if val := bps.New(1234.5670).SetRound(math.Floor).FloatString(3); val != "1234.567" {
		t.Error("round: 期待値と異なる", val)
	}
	if val := bps.New(1234.5675).SetRound(math.Floor).FloatString(3); val != "1234.567" {
		t.Error("round: 期待値と異なる", val)
	}
	if val := bps.New(1234.5671).SetRound(math.Floor).FloatString(3); val != "1234.567" {
		t.Error("round: 期待値と異なる", val)
	}
}

func TestBpsChangeDigit(t *testing.T) {
	if val := bps.New(1234.5678).ChangeDigit(bps.DigitPpm).String(); val != "1234.567800" {
		t.Error("digit: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).ChangeDigit(bps.DigitPpb).String(); val != "1234.567800000" {
		t.Error("digit: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).ChangeDigit(bps.DigitPpt).String(); val != "1234.567800000000" {
		t.Error("digit: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).ChangeDigit(bps.DigitPpq).String(); val != "1234.567800000000000" {
		t.Error("digit: 期待値と異なる", val)
	}
	bps.Init(bps.DigitPpb, math.Round)
	if val := bps.New(1234.567890412).ChangeDigit(bps.DigitPpm).String(); val != "1234.567890" {
		t.Error("digit: 期待値と異なる", val)
	}
	if val := bps.New(1234.567890512).ChangeDigit(bps.DigitPpm).String(); val != "1234.567891" {
		t.Error("digit: 期待値と異なる", val)
	}
	bps.Reset()
}

func TestBpsCalc(t *testing.T) {
	if val := bps.New(1234.5678).Calc("+", 10).String(); val != "1244.5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("+", 0.010).String(); val != "1234.5778" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("-", 10).String(); val != "1224.5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("-", 0.010).String(); val != "1234.5578" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("*", 10).String(); val != "12345.6780" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("*", 10.4).String(); val != "12345.6780" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("*", 9.5).String(); val != "12345.6780" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("*", 0.1).Error; val != bps.ErrorInvalidSourceIntegerZero {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("/", 10).String(); val != "123.4567" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("/", 10.4).String(); val != "123.4567" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("/", 9.5).String(); val != "123.4567" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("/", 0.1).Error; val != bps.ErrorInvalidSourceIntegerZero {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("%", 100).String(); val != "34.5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("%", 10).String(); val != "4.5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Calc("%", 0.1).String(); val != "0.678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(10).Calc("^", 2).String(); val != "100.0" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(10).Calc("^", 2.4).String(); val != "100.0" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(10).Calc("^", 1.5).String(); val != "100.0" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(10).Calc("^", 0.1).Error; val != bps.ErrorInvalidSourceIntegerZero {
		t.Error("string: 期待値と異なる", val)
	}
}

func TestBpsFormat(t *testing.T) {
	if val := bps.New(1234.5678).Format(nil); val != "1234.5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Format(nil, bps.OptionInt64); val != "1235" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Format(money.JaJP.Ptr()); val != "1,234.5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Format(money.Nl.Ptr()); val != "1.234,5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Format(money.Sv.Ptr()); val != "1 234,5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Format(money.DeCH.Ptr()); val != "1’234.5678" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Format(nil, bps.OptionFloatPercent); val != "1234.57" {
		t.Error("string: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Format(nil, bps.OptionFloatPerMille); val != "1234.568" {
		t.Error("string: 期待値と異なる", val)
	}
}

func TestBpsCurrency(t *testing.T) {
	if val := bps.New(1234.5678).Currency(money.JPY, money.JaJP); val != "¥1,235" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.JPY, money.Nl); val != "1.235¥" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.JPY, money.Sv); val != "1 235 ¥" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.USD, money.EnUS); val != "$1,234.57" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.USD, money.Nl); val != "1.234,57$" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.USD, money.Sv); val != "1 234,57 $" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.VND, money.ViVN); val != "1.235 ₫" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.DKK, money.Ms); val != "kr 1,234.57" {
		t.Error("currency: 期待値と異なる", val)
	}
	if val := bps.New(1234.5678).Currency(money.VND, money.DeCH); val != "₫ 1’235" {
		t.Error("currency: 期待値と異なる", val)
	}
}

type JsonTest struct {
	Data *bps.Bps `json:"data"`
}

func TestBpsJson(t *testing.T) {
	encode := JsonTest{
		Data: bps.New(1234.1234),
	}
	rs, err := json.Marshal(encode)
	if err != nil {
		t.Error(err)
	}
	if string(rs) != `{"data":"1234.1234"}` {
		t.Error("json: 期待値と異なる", string(rs))
	}
	decode := JsonTest{}
	if err := json.Unmarshal(rs, &decode); err != nil {
		t.Error(err)
	}
	if decode.Data.Raw.Int64() != 12341234 {
		t.Error("json: 期待値と異なる", decode.Data.Raw.Int64())
	}
}

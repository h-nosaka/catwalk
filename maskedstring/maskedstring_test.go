package maskedstring_test

import (
	"testing"

	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/maskedstring"
)

type Foo struct {
	Bar maskedstring.MaskedString `json:"bar"`
}

func TestMaskedString(t *testing.T) {
	data := Foo{Bar: "foo"}
	if data.Bar != "foo" {
		t.Errorf("assert equal: value: (%T)%+v, want: %+v", data.Bar, data.Bar, "foo")
	}
	text := base.ToJson(data, "")
	if text != `{"bar":"****"}` {
		t.Errorf("assert equal: value: (%T)%+v, want: %+v", text, text, `{"bar":"****"}`)
	}
}

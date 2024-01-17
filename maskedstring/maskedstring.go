package maskedstring

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// json化時に文字列をフィルタリングする
type MaskedString string

func MaskedStrings(src string) MaskedString {
	return MaskedString(src)
}

func (p MaskedString) String() string {
	return string(p)
}

func (p MaskedString) MarshalJSON() ([]byte, error) {
	return json.Marshal("****")
}

func (p *MaskedString) UnmarshalJSON(data []byte) error {
	value := ""
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = MaskedString(value)
	return nil
}

func (p MaskedString) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p *MaskedString) Scan(data interface{}) error {
	switch src := data.(type) {
	case string:
		*p = MaskedStrings(src)
	default:
		return fmt.Errorf("invalid type: %T", src)
	}
	return nil
}

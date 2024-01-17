package models

import (
	"encoding/json"
	"time"

	"github.com/h-nosaka/catwalk/bps"
	maskedstring "github.com/h-nosaka/catwalk/maskedString"
)

// データバリエーション
type ItemEnumuint uint8

const (
	ItemEnumuintCreated = ItemEnumuint(1)
	ItemEnumuintActive  = ItemEnumuint(2)
	ItemEnumuintDeleted = ItemEnumuint(3)
)

func (p ItemEnumuint) String() string {
	switch p {
	case ItemEnumuintCreated:
		return "Created"
	case ItemEnumuintActive:
		return "Active"
	case ItemEnumuintDeleted:
		return "Deleted"
	}
	return ""
}

func ItemEnumuints(key string) ItemEnumuint {
	switch key {
	case "Created":
		return ItemEnumuintCreated
	case "Active":
		return ItemEnumuintActive
	case "Deleted":
		return ItemEnumuintDeleted
	}
	return 0
}

func (p ItemEnumuint) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *ItemEnumuint) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ItemEnumuints(value)
	return nil
}

type ItemEnumstring string

const (
	ItemEnumstringCreated = ItemEnumstring("created")
	ItemEnumstringActive  = ItemEnumstring("active")
	ItemEnumstringDeleted = ItemEnumstring("deleted")
)

func (p ItemEnumstring) String() string {
	return string(p)
}

type ItemEnumbitfield uint64

const (
	ItemEnumbitfieldRead ItemEnumbitfield = 1 << iota
	ItemEnumbitfieldWrite
	ItemEnumbitfieldManage
	ItemEnumbitfieldAdmin
)

func (p ItemEnumbitfield) Check(flag ItemEnumbitfield) bool {
	return (p & flag) == flag
}

type Item struct {
	// column
	Id           string                    `json:"id"`    // ID
	Price        string                    `json:"price"` // 価格
	Int8         int8                      `json:"int8"`
	Int16        int16                     `json:"int16"`
	Int32        int                       `json:"int32"`
	Int64        int64                     `json:"int64"`
	Uint8        uint8                     `json:"uint8"`
	Uint16       uint16                    `json:"uint16"`
	Uint32       uint                      `json:"uint32"`
	Uint64       uint64                    `json:"uint64"`
	Float64      float64                   `json:"float64"`
	String       string                    `json:"string"`
	Fixstring    string                    `json:"fixstring"`
	TextS        string                    `json:"text_s"`
	TextM        string                    `json:"text_m"`
	TextL        string                    `json:"text_l"`
	BlobS        []byte                    `json:"blob_s"`
	BlobM        []byte                    `json:"blob_m"`
	BlobL        []byte                    `json:"blob_l"`
	Bytes        []byte                    `json:"bytes"`
	Json         json.RawMessage           `json:"json"`
	Timestamp    time.Time                 `json:"timestamp"`
	Datetime     time.Time                 `json:"datetime"`
	Enumuint     ItemEnumuint              `json:"enumuint"`
	Enumstring   ItemEnumstring            `json:"enumstring"`
	Enumbitfield ItemEnumbitfield          `json:"enumbitfield"`
	Bps          *bps.Bps                  `json:"bps"`
	Masked       maskedstring.MaskedString `json:"masked"`
	CreatedAt    time.Time                 `json:"created_at"` // 作成日
	UpdatedAt    time.Time                 `json:"updated_at"` // 更新日

	// relation
}

func (c *Item) JsonMap() map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(c.Json), &data); err != nil {
		return nil
	}
	return data
}

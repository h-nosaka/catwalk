package models

import (
	"encoding/json"
	"time"
)

// ピンコードマスタ
type PinUsage int8

const (
	PinUsageOnetime = PinUsage(1)
	PinUsagePin     = PinUsage(2)
)

func (p PinUsage) String() string {
	switch p {
	case PinUsageOnetime:
		return "Onetime"
	case PinUsagePin:
		return "Pin"
	}
	return ""
}

func PinUsages(key string) PinUsage {
	switch key {
	case "Onetime":
		return PinUsageOnetime
	case "Pin":
		return PinUsagePin
	}
	return 0
}

func (p PinUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PinUsage) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PinUsages(value)
	return nil
}

type Pin struct {
	// column
	Id        string    `json:"id"`         // ID
	Pin       string    `json:"pin"`        // ピン
	Usage     PinUsage  `json:"usage"`      // 用途
	CreatedAt time.Time `json:"created_at"` // 作成日
	UpdatedAt time.Time `json:"updated_at"` // 更新日

	// relation
}

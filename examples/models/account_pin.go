package models

import (
	"time"
)

// アカウントとピンの紐付け
type AccountPin struct {
	// column
	Id        string    `json:"id"`         // ID
	AccountId string    `json:"account_id"` // accounts.id
	PinId     string    `json:"pin_id"`     // pins.id
	ExpiredAt time.Time `json:"expired_at"` // PIN有効期限日時
	DeletedAt time.Time `json:"deleted_at"` // 使用済み日時
	CreatedAt time.Time `json:"created_at"` // 作成日
	UpdatedAt time.Time `json:"updated_at"` // 更新日

	// relation
	Account *Account `gorm:"foreignKey:AccountId;references:Id"`
	Pin     *Pin     `gorm:"foreignKey:PinId;references:Id"`
}

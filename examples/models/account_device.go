package models

import (
	"time"
)

// デバイス管理マスタ
type AccountDevice struct {
	// column
	Id          string    `json:"id"`            // ID
	AccountId   string    `json:"account_id"`    // accounts.id
	Uuid        string    `json:"uuid"`          // デバイスID
	ActivatedAt time.Time `json:"activated_at"`  // アクティベート日時
	LastLoginAt time.Time `json:"last_login_at"` // 最終ログイン日時
	CreatedAt   time.Time `json:"created_at"`    // 作成日
	UpdatedAt   time.Time `json:"updated_at"`    // 更新日

	// relation
	Account *Account `gorm:"foreignKey:AccountId;references:Id"`
}

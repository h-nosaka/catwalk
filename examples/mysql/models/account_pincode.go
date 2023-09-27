package models

import (
	"gorm.io/gorm"
	"time"
)

// アカウントとピンコードの紐付け
type AccountPincode struct {
	// column
	Id        uint64     `json:"id"`         // primary key
	AccountId string     `json:"account_id"` // accounts.id
	PincodeId uint64     `json:"pincode_id"` // pincodes.id
	ExpiredAt *time.Time `json:"expired_at"` // PIN有効期限日時
	DeletedAt *time.Time `json:"deleted_at"` // 使用済み日時
	CreatedAt *time.Time `json:"created_at"` // 作成日
	UpdatedAt *time.Time `json:"updated_at"` // 更新日

	// relation
	Account *Account `gorm:"foreignKey:AccountId;references:Id"`
	Pincode *Pincode `gorm:"foreignKey:PincodeId;references:Id"`
}

func (p *AccountPincode) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.First(p).Error; err != nil {
		return err
	}
	return nil
}

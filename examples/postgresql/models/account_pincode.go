package model

import (
	"gorm.io/gorm"
	"time"
)

// アカウントとピンコードの紐付け
type AccountPincode struct {
	// column
	Id        int        `json:"id"`                           // primary key
	AccountId int64      `json:"account_id" gorm:"primarykey"` // accounts.id
	PincodeId int64      `json:"pincode_id" gorm:"primarykey"` // pincodes.id
	ExpiredAt *time.Time `json:"expired_at" gorm:"primarykey"` // PIN有効期限日時
	DeletedAt *time.Time `json:"deleted_at" gorm:"primarykey"` // 使用済み日時
	CreatedAt *time.Time `json:"created_at" gorm:"primarykey"` // 作成日
	UpdatedAt *time.Time `json:"updated_at" gorm:"primarykey"` // 更新日

	// relation
	Accounts []Account `gorm:"foreignKey:AccountId;references:Id"`
	Pincodes []Pincode `gorm:"foreignKey:PincodeId;references:Id"`
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

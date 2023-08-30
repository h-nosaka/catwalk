package models

import (
	"gorm.io/gorm"
	"time"
)

// アカウントとピンコードの紐付け
type AccountPincode struct {
	// column
	Id            string     `json:"id" gorm:"column:id;primarykey;size:255;default:uuid_generate_v4()"` // primary key
	TestAccountId string     `json:"test_account_id" gorm:"column:test_account_id"`                      // accounts.id
	TestPincodeId string     `json:"test_pincode_id" gorm:"column:test_pincode_id"`                      // pincodes.id
	ExpiredAt     *time.Time `json:"expired_at" gorm:"column:expired_at"`                                // PIN有効期限日時
	DeletedAt     *time.Time `json:"deleted_at" gorm:"column:deleted_at"`                                // 使用済み日時
	CreatedAt     *time.Time `json:"created_at" gorm:"column:created_at"`                                // 作成日
	UpdatedAt     *time.Time `json:"updated_at" gorm:"column:updated_at"`                                // 更新日

	// relation
	TestAccount *Account `gorm:"foreignKey:TestAccountId;references:Id"`
	TestPincode *Pincode `gorm:"foreignKey:TestPincodeId;references:Id"`
}

func (p *AccountPincode) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.Where(p).First(p).Error; err != nil {
		return err
	}
	return nil
}

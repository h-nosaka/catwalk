package model

import (
	"gorm.io/gorm"
	"time"
)

// UUID管理マスタ
type AccountActivate struct {
	// column
	Id          int        `json:"id"`                              // primary key
	AccountId   int64      `json:"account_id" gorm:"primarykey"`    // accounts.id
	Uuid        string     `json:"uuid" gorm:"primarykey"`          // UUID
	PincodeId   int64      `json:"pincode_id" gorm:"primarykey"`    // pincodes.id
	ExpiredAt   *time.Time `json:"expired_at" gorm:"primarykey"`    // PIN有効期限日時
	ActivatedAt *time.Time `json:"activated_at" gorm:"primarykey"`  // アクティベート日時
	LastLoginAt *time.Time `json:"last_login_at" gorm:"primarykey"` // 最終ログイン日時
	CreatedAt   *time.Time `json:"created_at" gorm:"primarykey"`    // 作成日
	UpdatedAt   *time.Time `json:"updated_at" gorm:"primarykey"`    // 更新日

	// relation
	Accounts []Account `gorm:"foreignKey:AccountId;references:Id"`
	Pincodes []Pincode `gorm:"foreignKey:PincodeId;references:Id"`
}

func (p *AccountActivate) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.First(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *AccountActivate) Active(db *gorm.DB, uuid string) error {
	return db.Joins("Account").Where("uuid = ? and activated_at is not NULL and Account.status = ?", uuid, AccountStatusActivated).First(&p).Error
}
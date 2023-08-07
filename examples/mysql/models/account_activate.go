package models

import (
	"gorm.io/gorm"
	"time"
)

// UUID管理マスタ
type AccountActivate struct {
	// column
	Id          uint64     `json:"id"`            // primary key
	AccountId   uint64     `json:"account_id"`    // accounts.id
	Uuid        string     `json:"uuid"`          // UUID
	PincodeId   uint64     `json:"pincode_id"`    // pincodes.id
	ExpiredAt   *time.Time `json:"expired_at"`    // PIN有効期限日時
	ActivatedAt *time.Time `json:"activated_at"`  // アクティベート日時
	LastLoginAt *time.Time `json:"last_login_at"` // 最終ログイン日時
	CreatedAt   *time.Time `json:"created_at"`    // 作成日
	UpdatedAt   *time.Time `json:"updated_at"`    // 更新日

	// relation
	Account *Account `gorm:"foreignKey:AccountId;references:Id"`
	Pincode *Pincode `gorm:"foreignKey:PincodeId;references:Id"`
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

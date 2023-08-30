package models

import (
	"gorm.io/gorm"
	"time"
)

// UUID管理マスタ
type AccountActivate struct {
	// column
	Id          string     `json:"id" gorm:"column:id;primarykey;size:255;default:uuid_generate_v4()"` // primary key
	AccountId   string     `json:"account_id" gorm:"column:account_id"`                                // accounts.id
	Uuid        string     `json:"uuid" gorm:"column:uuid"`                                            // UUID
	PincodeId   string     `json:"pincode_id" gorm:"column:pincode_id"`                                // pincodes.id
	ExpiredAt   *time.Time `json:"expired_at" gorm:"column:expired_at"`                                // PIN有効期限日時
	ActivatedAt *time.Time `json:"activated_at" gorm:"column:activated_at"`                            // アクティベート日時
	LastLoginAt *time.Time `json:"last_login_at" gorm:"column:last_login_at"`                          // 最終ログイン日時
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`                                // 作成日
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`                                // 更新日

	// relation
	Accounts []Account `gorm:"foreignKey:AccountId;references:Id"`
	Pincodes []Pincode `gorm:"foreignKey:PincodeId;references:Id"`
}

func (p *AccountActivate) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.Where(p).First(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *AccountActivate) Active(db *gorm.DB, uuid string) error {
	return db.Joins("Account").Where("uuid = ? and activated_at is not NULL and Account.status = ?", uuid, AccountStatusActivated).First(&p).Error
}

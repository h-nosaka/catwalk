package models

import (
	"gorm.io/gorm"
	"time"
)

// ピンコードマスタ
type Pincode struct {
	// column
	Id        string     `json:"id" gorm:"column:id;primarykey;size:255;default:uuid_generate_v4()"` // primary key
	Pin       string     `json:"pin" gorm:"column:pin"`                                              // ピン
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`                                // 作成日
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`                                // 更新日

	// relation
}

func (p *Pincode) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.Where(p).First(p).Error; err != nil {
		return err
	}
	return nil
}

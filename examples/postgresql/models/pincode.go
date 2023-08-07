package model

import (
	"gorm.io/gorm"
	"time"
)

// ピンコードマスタ
type Pincode struct {
	// column
	Id        int        `json:"id"`                           // primary key
	Pin       string     `json:"pin" gorm:"primarykey"`        // ピン
	CreatedAt *time.Time `json:"created_at" gorm:"primarykey"` // 作成日
	UpdatedAt *time.Time `json:"updated_at" gorm:"primarykey"` // 更新日

	// relation
}

func (p *Pincode) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.First(p).Error; err != nil {
		return err
	}
	return nil
}

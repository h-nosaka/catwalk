package models

import (
	"time"

	"github.com/h-nosaka/catwalk/bps"
	"gorm.io/gorm"
)

// アイテムマスタ
type Item struct {
	// column
	Id        string     `json:"id"`         // primary key
	Price     *bps.Bps   `json:"price"`      // メールアドレス
	CreatedAt *time.Time `json:"created_at"` // 作成日
	UpdatedAt *time.Time `json:"updated_at"` // 更新日

	// relation
}

func (p *Item) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.First(p).Error; err != nil {
		return err
	}
	return nil
}

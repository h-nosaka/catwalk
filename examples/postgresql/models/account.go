package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

// アカウントマスタ
type AccountRole uint64

const (
	AccountRoleViewer AccountRole = 1 << iota
	AccountRoleWriter
	AccountRoleManager
)

func (p AccountRole) Check(flag AccountRole) bool {
	return (p & flag) == flag
}

type AccountStatus uint

const (
	AccountStatusCreated   = AccountStatus(0)
	AccountStatusActivated = AccountStatus(1)
	AccountStatusFreezed   = AccountStatus(8)
	AccountStatusDeleted   = AccountStatus(9)
)

func (p AccountStatus) String() string {
	switch p {
	case AccountStatusCreated:
		return "Created"
	case AccountStatusActivated:
		return "Activated"
	case AccountStatusFreezed:
		return "Freezed"
	case AccountStatusDeleted:
		return "Deleted"
	}
	return ""
}

func AccountStatuses(key string) AccountStatus {
	switch key {
	case "Created":
		return AccountStatusCreated
	case "Activated":
		return AccountStatusActivated
	case "Freezed":
		return AccountStatusFreezed
	case "Deleted":
		return AccountStatusDeleted
	}
	return 0
}

func (p AccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *AccountStatus) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = AccountStatuses(value)
	return nil
}

type Account struct {
	// column
	Id             string        `json:"id" gorm:"column:id;primarykey;size:255;default:uuid_generate_v4()"` // primary key
	Email          string        `json:"email" gorm:"column:email"`                                          // メールアドレス
	HashedPassword string        `json:"hashed_password" gorm:"column:hashed_password"`                      // ハッシュ化済みパスワード
	Salt           string        `json:"salt" gorm:"column:salt"`                                            // ソルト
	Code           string        `json:"code" gorm:"column:code"`                                            // 表示ID
	NotificationId *int64        `json:"notification_id" gorm:"column:notification_id"`                      // notifications.id
	Role           AccountRole   `json:"role" gorm:"column:role"`                                            // ロール
	Status         AccountStatus `json:"status" gorm:"column:status"`                                        // ステータス
	Flags          *int          `json:"flags" gorm:"column:flags"`                                          // フラグ
	FreezedAt      *time.Time    `json:"freezed_at" gorm:"column:freezed_at"`                                // 凍結日
	DeletedAt      *time.Time    `json:"deleted_at" gorm:"column:deleted_at"`                                // 削除日
	CreatedAt      *time.Time    `json:"created_at" gorm:"column:created_at"`                                // 作成日
	UpdatedAt      *time.Time    `json:"updated_at" gorm:"column:updated_at"`                                // 更新日

	// relation
}

func (p *Account) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.Where(p).First(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *Account) Auth(db *gorm.DB, email string) bool {
	db.Where("Email = ? and status = ?", email, AccountStatusActivated).First(&p)
	return p.Email == email
}

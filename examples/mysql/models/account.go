package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

// アカウントマスタ
type AccountRole uint64

const (
	AccountRoleManager AccountRole = 1 << iota
	AccountRoleViewer
	AccountRoleWriter
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
	case AccountStatusActivated:
		return "Activated"
	case AccountStatusFreezed:
		return "Freezed"
	case AccountStatusDeleted:
		return "Deleted"
	case AccountStatusCreated:
		return "Created"
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
	Id             uint64        `json:"id"`              // primary key
	Email          string        `json:"email"`           // メールアドレス
	HashedPassword string        `json:"hashed_password"` // ハッシュ化済みパスワード
	Salt           string        `json:"salt"`            // ソルト
	Code           string        `json:"code"`            // 表示ID
	NotificationId *uint64       `json:"notification_id"` // notifications.id
	Role           AccountRole   `json:"role"`            // ロール
	Status         AccountStatus `json:"status"`          // ステータス
	Flags          *uint         `json:"flags"`           // フラグ
	FreezedAt      *time.Time    `json:"freezed_at"`      // 削除日
	DeletedAt      *time.Time    `json:"deleted_at"`      // 削除日
	CreatedAt      *time.Time    `json:"created_at"`      // 作成日
	UpdatedAt      *time.Time    `json:"updated_at"`      // 更新日

	// relation
	AccountActivate *AccountActivate `gorm:"foreignKey:AccountId;references:Id"`
}

func (p *Account) Find(db *gorm.DB, preloads ...string) error {
	tx := db
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	if err := tx.First(p).Error; err != nil {
		return err
	}
	return nil
}

func (p *Account) Auth(db *gorm.DB, email string) bool {
	db.Where("Email = ? and status = ?", email, AccountStatusActivated).First(&p)
	return p.Email == email
}

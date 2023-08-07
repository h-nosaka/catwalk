package model

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
	Id             int           `json:"id"`                                // primary key
	Email          string        `json:"email" gorm:"primarykey"`           // メールアドレス
	HashedPassword string        `json:"hashed_password" gorm:"primarykey"` // ハッシュ化済みパスワード
	Salt           string        `json:"salt" gorm:"primarykey"`            // ソルト
	Code           string        `json:"code" gorm:"primarykey"`            // 表示ID
	NotificationId *int64        `json:"notification_id" gorm:"primarykey"` // notifications.id
	Role           AccountRole   `json:"role" gorm:"primarykey"`            // ロール
	Status         AccountStatus `json:"status" gorm:"primarykey"`          // ステータス
	Flags          *int          `json:"flags" gorm:"primarykey"`           // フラグ
	FreezedAt      *time.Time    `json:"freezed_at" gorm:"primarykey"`      // 凍結日
	DeletedAt      *time.Time    `json:"deleted_at" gorm:"primarykey"`      // 削除日
	CreatedAt      *time.Time    `json:"created_at" gorm:"primarykey"`      // 作成日
	UpdatedAt      *time.Time    `json:"updated_at" gorm:"primarykey"`      // 更新日

	// relation
	AccountActivates []AccountActivate `gorm:"foreignKey:AccountId;references:Id"`
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

package models

import (
	"encoding/json"
	"time"
)

// アカウントマスタ
type AccountRole uint

const (
	AccountRoleViewer AccountRole = 1 << iota
	AccountRoleWriter
	AccountRoleManager
)

func (p AccountRole) Check(flag AccountRole) bool {
	return (p & flag) == flag
}

type AccountState uint8

const (
	AccountStateCreated   = AccountState(1)
	AccountStateActivated = AccountState(2)
	AccountStateFreezed   = AccountState(8)
	AccountStateDeleted   = AccountState(9)
)

func (p AccountState) String() string {
	switch p {
	case AccountStateCreated:
		return "Created"
	case AccountStateActivated:
		return "Activated"
	case AccountStateFreezed:
		return "Freezed"
	case AccountStateDeleted:
		return "Deleted"
	}
	return ""
}

func AccountStates(key string) AccountState {
	switch key {
	case "Created":
		return AccountStateCreated
	case "Activated":
		return AccountStateActivated
	case "Freezed":
		return AccountStateFreezed
	case "Deleted":
		return AccountStateDeleted
	}
	return 0
}

func (p AccountState) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *AccountState) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = AccountStates(value)
	return nil
}

type Account struct {
	// column
	Id             string       `json:"id"`              // ID
	Email          string       `json:"email"`           // メールアドレス
	HashedPassword string       `json:"hashed_password"` // ハッシュ化済みパスワード
	Salt           string       `json:"salt"`            // ソルト
	NotificationId string       `json:"notification_id"` // notifications.id
	Role           AccountRole  `json:"role"`            // ロール
	State          AccountState `json:"state"`           // ステータス
	Flags          uint64       `json:"flags"`           // フラグ
	FreezedAt      time.Time    `json:"freezed_at"`      // 凍結日
	DeletedAt      time.Time    `json:"deleted_at"`      // 削除日
	CreatedAt      time.Time    `json:"created_at"`      // 作成日
	UpdatedAt      time.Time    `json:"updated_at"`      // 更新日

	// relation
	AccountDevices []AccountDevice `gorm:"foreignKey:Id"`
}

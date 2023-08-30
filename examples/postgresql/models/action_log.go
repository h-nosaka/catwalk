package models

import (
	"encoding/json"
	"time"
)

// アクションログ ESIDX
type ActionLogActionType uint

const (
	ActionLogActionTypeINACTIVE   = ActionLogActionType(2)
	ActionLogActionTypeSAVEYOU    = ActionLogActionType(11)
	ActionLogActionTypeKOTADMIN   = ActionLogActionType(13)
	ActionLogActionTypeGAROON     = ActionLogActionType(14)
	ActionLogActionTypeRESUMED    = ActionLogActionType(1)
	ActionLogActionTypePAUSED     = ActionLogActionType(3)
	ActionLogActionTypeDETACHED   = ActionLogActionType(4)
	ActionLogActionTypeKINGOFTIME = ActionLogActionType(12)
	ActionLogActionTypeCLOUDMAIL  = ActionLogActionType(15)
	ActionLogActionTypeSLACK      = ActionLogActionType(16)
)

func (p ActionLogActionType) String() string {
	switch p {
	case ActionLogActionTypeINACTIVE:
		return "INACTIVE"
	case ActionLogActionTypeSAVEYOU:
		return "SAVEYOU"
	case ActionLogActionTypeKOTADMIN:
		return "KOTADMIN"
	case ActionLogActionTypeGAROON:
		return "GAROON"
	case ActionLogActionTypeRESUMED:
		return "RESUMED"
	case ActionLogActionTypePAUSED:
		return "PAUSED"
	case ActionLogActionTypeDETACHED:
		return "DETACHED"
	case ActionLogActionTypeKINGOFTIME:
		return "KINGOFTIME"
	case ActionLogActionTypeCLOUDMAIL:
		return "CLOUDMAIL"
	case ActionLogActionTypeSLACK:
		return "SLACK"
	}
	return ""
}

func ActionLogActionTypes(key string) ActionLogActionType {
	switch key {
	case "KINGOFTIME":
		return ActionLogActionTypeKINGOFTIME
	case "CLOUDMAIL":
		return ActionLogActionTypeCLOUDMAIL
	case "SLACK":
		return ActionLogActionTypeSLACK
	case "RESUMED":
		return ActionLogActionTypeRESUMED
	case "PAUSED":
		return ActionLogActionTypePAUSED
	case "DETACHED":
		return ActionLogActionTypeDETACHED
	case "GAROON":
		return ActionLogActionTypeGAROON
	case "INACTIVE":
		return ActionLogActionTypeINACTIVE
	case "SAVEYOU":
		return ActionLogActionTypeSAVEYOU
	case "KOTADMIN":
		return ActionLogActionTypeKOTADMIN
	}
	return 0
}

func (p ActionLogActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *ActionLogActionType) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = ActionLogActionTypes(value)
	return nil
}

type ActionLog struct {
	// column
	Id         string              `json:"_id" gorm:"column:_id"`                 // ID
	Uuid       string              `json:"uuid" gorm:"column:uuid"`               // UUID
	Email      *string             `json:"email" gorm:"column:email"`             // メールアドレス
	ActionType ActionLogActionType `json:"action_type" gorm:"column:action_type"` // タイプ
	Message    string              `json:"message" gorm:"column:message"`         // メッセージ
	RecordedAt time.Time           `json:"recorded_at" gorm:"column:recorded_at"` // 実行日時
	CreatedAt  *time.Time          `json:"created_at" gorm:"column:created_at"`   // 作成日
	UpdatedAt  *time.Time          `json:"updated_at" gorm:"column:updated_at"`   // 更新日

	// relation
}

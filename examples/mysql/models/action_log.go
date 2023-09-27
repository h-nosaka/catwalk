package models

import (
	"encoding/json"
	"time"
)

// アクションログ ESIDX
type ActionLogActionType uint

const (
	ActionLogActionTypeRESUMED    = ActionLogActionType(1)
	ActionLogActionTypeINACTIVE   = ActionLogActionType(2)
	ActionLogActionTypePAUSED     = ActionLogActionType(3)
	ActionLogActionTypeDETACHED   = ActionLogActionType(4)
	ActionLogActionTypeSAVEYOU    = ActionLogActionType(11)
	ActionLogActionTypeKINGOFTIME = ActionLogActionType(12)
	ActionLogActionTypeKOTADMIN   = ActionLogActionType(13)
	ActionLogActionTypeGAROON     = ActionLogActionType(14)
	ActionLogActionTypeCLOUDMAIL  = ActionLogActionType(15)
	ActionLogActionTypeSLACK      = ActionLogActionType(16)
)

func (p ActionLogActionType) String() string {
	switch p {
	case ActionLogActionTypeRESUMED:
		return "RESUMED"
	case ActionLogActionTypeINACTIVE:
		return "INACTIVE"
	case ActionLogActionTypePAUSED:
		return "PAUSED"
	case ActionLogActionTypeDETACHED:
		return "DETACHED"
	case ActionLogActionTypeSAVEYOU:
		return "SAVEYOU"
	case ActionLogActionTypeKINGOFTIME:
		return "KINGOFTIME"
	case ActionLogActionTypeKOTADMIN:
		return "KOTADMIN"
	case ActionLogActionTypeGAROON:
		return "GAROON"
	case ActionLogActionTypeCLOUDMAIL:
		return "CLOUDMAIL"
	case ActionLogActionTypeSLACK:
		return "SLACK"
	}
	return ""
}

func ActionLogActionTypes(key string) ActionLogActionType {
	switch key {
	case "RESUMED":
		return ActionLogActionTypeRESUMED
	case "INACTIVE":
		return ActionLogActionTypeINACTIVE
	case "PAUSED":
		return ActionLogActionTypePAUSED
	case "DETACHED":
		return ActionLogActionTypeDETACHED
	case "SAVEYOU":
		return ActionLogActionTypeSAVEYOU
	case "KINGOFTIME":
		return ActionLogActionTypeKINGOFTIME
	case "KOTADMIN":
		return ActionLogActionTypeKOTADMIN
	case "GAROON":
		return ActionLogActionTypeGAROON
	case "CLOUDMAIL":
		return ActionLogActionTypeCLOUDMAIL
	case "SLACK":
		return ActionLogActionTypeSLACK
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
	Id         string              `json:"_id"`         // ID
	Uuid       string              `json:"uuid"`        // UUID
	Email      *string             `json:"email"`       // メールアドレス
	ActionType ActionLogActionType `json:"action_type"` // タイプ
	Message    string              `json:"message"`     // メッセージ
	RecordedAt time.Time           `json:"recorded_at"` // 実行日時
	CreatedAt  *time.Time          `json:"created_at"`  // 作成日
	UpdatedAt  *time.Time          `json:"updated_at"`  // 更新日

	// relation
}

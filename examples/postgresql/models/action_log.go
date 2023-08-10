package model

import (
	"encoding/json"
	"time"
)

// アクションログ ESIDX
type ActionLogActionType uint

const (
	ActionLogActionTypePAUSED     = ActionLogActionType(3)
	ActionLogActionTypeSAVEYOU    = ActionLogActionType(11)
	ActionLogActionTypeRESUMED    = ActionLogActionType(1)
	ActionLogActionTypeDETACHED   = ActionLogActionType(4)
	ActionLogActionTypeKINGOFTIME = ActionLogActionType(12)
	ActionLogActionTypeKOTADMIN   = ActionLogActionType(13)
	ActionLogActionTypeGAROON     = ActionLogActionType(14)
	ActionLogActionTypeCLOUDMAIL  = ActionLogActionType(15)
	ActionLogActionTypeSLACK      = ActionLogActionType(16)
	ActionLogActionTypeINACTIVE   = ActionLogActionType(2)
)

func (p ActionLogActionType) String() string {
	switch p {
	case ActionLogActionTypeGAROON:
		return "GAROON"
	case ActionLogActionTypeCLOUDMAIL:
		return "CLOUDMAIL"
	case ActionLogActionTypeSLACK:
		return "SLACK"
	case ActionLogActionTypeINACTIVE:
		return "INACTIVE"
	case ActionLogActionTypeDETACHED:
		return "DETACHED"
	case ActionLogActionTypeKINGOFTIME:
		return "KINGOFTIME"
	case ActionLogActionTypeKOTADMIN:
		return "KOTADMIN"
	case ActionLogActionTypeRESUMED:
		return "RESUMED"
	case ActionLogActionTypePAUSED:
		return "PAUSED"
	case ActionLogActionTypeSAVEYOU:
		return "SAVEYOU"
	}
	return ""
}

func ActionLogActionTypes(key string) ActionLogActionType {
	switch key {
	case "CLOUDMAIL":
		return ActionLogActionTypeCLOUDMAIL
	case "SLACK":
		return ActionLogActionTypeSLACK
	case "INACTIVE":
		return ActionLogActionTypeINACTIVE
	case "DETACHED":
		return ActionLogActionTypeDETACHED
	case "KINGOFTIME":
		return ActionLogActionTypeKINGOFTIME
	case "KOTADMIN":
		return ActionLogActionTypeKOTADMIN
	case "GAROON":
		return ActionLogActionTypeGAROON
	case "RESUMED":
		return ActionLogActionTypeRESUMED
	case "PAUSED":
		return ActionLogActionTypePAUSED
	case "SAVEYOU":
		return ActionLogActionTypeSAVEYOU
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

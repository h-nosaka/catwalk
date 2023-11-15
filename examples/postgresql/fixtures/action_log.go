package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func ActionLog(setter func(model *models.ActionLog)) *models.ActionLog {
	model := &models.ActionLog{
		Id: uuid.NewString(),
	}
	setter(model)
	return model
}

func CreateActionLog(db *gorm.DB, setter func(model *models.ActionLog)) *models.ActionLog {
	model := ActionLog(setter)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

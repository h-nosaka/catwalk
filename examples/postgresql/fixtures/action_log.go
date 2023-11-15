package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func ActionLog(setters ...func(model *models.ActionLog)) *models.ActionLog {
	model := &models.ActionLog{
		Id: uuid.NewString(),
	}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateActionLog(db *gorm.DB, setters ...func(model *models.ActionLog)) *models.ActionLog {
	model := ActionLog(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

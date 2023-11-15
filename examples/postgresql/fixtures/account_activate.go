package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func AccountActivate(setters ...func(model *models.AccountActivate)) *models.AccountActivate {
	model := &models.AccountActivate{
		Id: uuid.NewString(),
	}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateAccountActivate(db *gorm.DB, setters ...func(model *models.AccountActivate)) *models.AccountActivate {
	model := AccountActivate(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

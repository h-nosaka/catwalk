package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/mysql/models"
	"gorm.io/gorm"
)

func Account(setters ...func(model *models.Account)) *models.Account {
	model := &models.Account{
		Id: uuid.NewString()}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateAccount(db *gorm.DB, setters ...func(model *models.Account)) *models.Account {
	model := Account(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

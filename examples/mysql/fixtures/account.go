package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/mysql/models"
	"gorm.io/gorm"
)

func Account(setter func(model *models.Account)) *models.Account {
	model := &models.Account{
		Id: uuid.NewString()}
	setter(model)
	return model
}

func CreateAccount(db *gorm.DB, setter func(model *models.Account)) *models.Account {
	model := Account(setter)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func AccountPincode(setters ...func(model *models.AccountPincode)) *models.AccountPincode {
	model := &models.AccountPincode{
		Id: uuid.NewString(),
	}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateAccountPincode(db *gorm.DB, setters ...func(model *models.AccountPincode)) *models.AccountPincode {
	model := AccountPincode(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

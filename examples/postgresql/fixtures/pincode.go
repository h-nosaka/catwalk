package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func Pincode(setter func(model *models.Pincode)) *models.Pincode {
	model := &models.Pincode{
		Id: uuid.NewString(),
	}
	setter(model)
	return model
}

func CreatePincode(db *gorm.DB, setter func(model *models.Pincode)) *models.Pincode {
	model := Pincode(setter)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

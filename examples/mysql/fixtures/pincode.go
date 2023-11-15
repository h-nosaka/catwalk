package fixtures

import (
	"github.com/h-nosaka/catwalk/examples/mysql/models"
	"gorm.io/gorm"
)

func Pincode(setter func(model *models.Pincode)) *models.Pincode {
	model := &models.Pincode{}
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

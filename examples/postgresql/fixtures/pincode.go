package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func Pincode(setters ...func(model *models.Pincode)) *models.Pincode {
	model := &models.Pincode{
		Id: uuid.NewString(),
	}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreatePincode(db *gorm.DB, setters ...func(model *models.Pincode)) *models.Pincode {
	model := Pincode(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}

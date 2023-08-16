package schema

import (
	"testing"

	"github.com/h-nosaka/catwalk/base"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	db "github.com/h-nosaka/catwalk/postgresql"
)

func TestAccountsCreate(t *testing.T) {
	SetEnv()
	base.Init()
	tbl := Accounts()
	sql := tbl.Create()
	t.Log(sql)
}

func TestAccountsRun(t *testing.T) {
	SetEnv()
	base.Init()
	schema := &db.ISchema{
		Name: "app",
		Tables: []db.ITable{
			Accounts(),
		},
	}
	schema.Run()
}

func TestAccountsModel(t *testing.T) {
	SetEnv()
	base.Init()
	data := models.Account{
		Email:          "test@test.com",
		HashedPassword: "password",
		Salt:           "12345678",
		Code:           "test",
		Role:           models.AccountRoleManager,
		Status:         models.AccountStatusActivated,
	}
	if err := base.DB.Where(data).First(&data).Error; err != nil {
		if e := base.DB.Save(&data).Error; e != nil {
			t.Error(e)
		}
	}
	t.Log(data)
	pin := models.Pincode{
		Pin: "123456",
	}
	if err := base.DB.Where(pin).First(&pin).Error; err != nil {
		if e := base.DB.Save(&pin).Error; e != nil {
			t.Error(e)
		}
	}
	t.Log(pin)
	ap := models.AccountPincode{
		TestAccountId: data.Id,
		TestPincodeId: pin.Id,
	}
	if err := base.DB.Debug().Where(ap).First(&ap).Error; err != nil {
		if e := base.DB.Debug().Save(&ap).Error; e != nil {
			t.Error(e)
		}
	}
	t.Logf("%+v", ap)
}

func TestAccountsModelRelation(t *testing.T) {
	SetEnv()
	base.Init()
	rs := models.AccountPincode{}
	// if err := base.DB.Debug().Where(rs).First(&rs).Error; err != nil {
	if err := base.DB.Debug().Preload("TestPincode").Preload("TestAccount").First(&rs).Error; err != nil {
		t.Error(err)
	}
	t.Logf("%+v", rs)
	t.Error("test")
}

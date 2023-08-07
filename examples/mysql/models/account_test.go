package models_test

import (
	"encoding/json"
	"testing"

	"github.com/h-nosaka/catwalk/examples/mysql/models"
)

func TestAccount(t *testing.T) {
	account := models.Account{
		Id:     1,
		Email:  "test@example.com",
		Role:   models.AccountRoleManager | models.AccountRoleWriter,
		Status: models.AccountStatusActivated,
	}
	rs, err := json.Marshal(account)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(rs))
	if err := json.Unmarshal(rs, &account); err != nil {
		t.Error(err)
	}
	t.Logf("%+v", account)
	if account.Status != models.AccountStatusActivated {
		t.Errorf("account status is got: %d, want: %d", account.Status, models.AccountStatusActivated)
	}
	t.Logf("status: %d", account.Status)
	t.Log(account.Role.Check(models.AccountRoleManager))
	t.Log(account.Role.Check(models.AccountRoleWriter))
	t.Log(account.Role.Check(models.AccountRoleViewer))
	if !account.Role.Check(models.AccountRoleManager) {
		t.Errorf("role error: %d, want: %d", account.Role, models.AccountRoleManager)
	}
	if !account.Role.Check(models.AccountRoleWriter) {
		t.Errorf("role error: %d, want: %d", account.Role, models.AccountRoleWriter)
	}
	if account.Role.Check(models.AccountRoleViewer) {
		t.Errorf("role error: %d, want: %d", account.Role, models.AccountRoleViewer)
	}
	// t.Error("debug")
}

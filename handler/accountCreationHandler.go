package handler

import (
	"time"

	"github.com/denisaguilar/dummy_automation/external/vault"
	"github.com/denisaguilar/dummy_automation/generated/models"
	"github.com/denisaguilar/dummy_automation/generated/restapi/operations/account"
	"github.com/go-openapi/runtime/middleware"
)

func CreateAccountHandler(params account.CreateAccountParams) middleware.Responder {
	ad, ed := processRequest(&params)
	if ed != nil {
		return account.NewCreateAccountDefault(500).WithPayload(ed)
	}
	return account.NewCreateAccountOK().WithPayload(ad)
}

func processRequest(params *account.CreateAccountParams) (*models.AccountDetails, *models.ErrorDetails) {
	_, err := vault.SysHealth()
	if err != nil {
		e := &models.ErrorDetails{
			Message: "Error when verifing SysHealth",
			Time:    time.Now().String(),
		}
		return nil, e
	}
	a, err := createAccount(params.AccountCreationRequest.Name)
	if err != nil {
		e := &models.ErrorDetails{
			Message: "Error when verifing SysHealth",
			Time:    time.Now().String(),
		}
		return nil, e
	}

	return a, nil
}

func createAccount(accountName string) (*models.AccountDetails, error) {
	accountDetails := models.AccountDetails{
		Name: accountName + "_account",
		ID:   "xxx",
	}
	return &accountDetails, nil
}

package handler

import (
	"github.com/denisaguilar/dummy_automation/generated/models"
	"github.com/denisaguilar/dummy_automation/generated/restapi/operations/account"
	"github.com/go-openapi/runtime/middleware"
)

func CreateAccountHandler(params account.CreateAccountParams) middleware.Responder {
	accountDetails := models.AccountDetails{
		Name: params.AccountCreationRequest.Name + "_account",
	}
	return account.NewCreateAccountOK().WithPayload(&accountDetails)
}

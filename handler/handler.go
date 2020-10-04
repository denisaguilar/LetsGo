package handler

import (
	"github.com/denisaguilar/dummy_automation/generated/restapi/operations"
	"github.com/denisaguilar/dummy_automation/generated/restapi/operations/account"
)

func SetUpHandler(api *operations.DummyAutomationAPI) {
	api.AccountCreateAccountHandler = account.CreateAccountHandlerFunc(CreateAccountHandler)
}

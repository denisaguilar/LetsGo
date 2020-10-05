package handler

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/denisaguilar/dummy_automation/restclient"
	"github.com/denisaguilar/dummy_automation/utils/mocks"
)

func init() {
	restclient.Client = &mocks.MockClient{}
}

func TestAccountCreation(t *testing.T) {
	//given
	accountName := "account_test"
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	//when
	resp := createAccount(accountName)

	//then
	log.Println(resp)
}

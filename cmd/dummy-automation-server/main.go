package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"

	"github.com/denisaguilar/dummy_automation/generated/restapi"
	"github.com/denisaguilar/dummy_automation/generated/restapi/operations"
	"github.com/denisaguilar/dummy_automation/handler"
)

var portFlag = flag.Int("port", 8080, "Port to run this service on")

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewDummyAutomationAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	handler.SetUpHandler(api)

	server.Port = *portFlag
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

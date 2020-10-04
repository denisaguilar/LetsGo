## Install Go Swagger

    git clone https://github.com/go-swagger/go-swagger
    go install ./cmd/swagger

## Generate Swagger files

    swagger generate server -t generated -A dummy_automation -f .\swagger\swagger.yaml --exclude-main

## Run application

    go run .\cmd\dummy-automation-server\main.go --port 8080

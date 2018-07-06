package main

import (
	"fmt"

	"github.com/go_microservices/goblog/accountservice/dbclient"
	service "github.com/go_microservices/goblog/accountservice/service"
	"github.com/go_microservices/goblog/services"
)

var appName = "accountservice"
var port = "8000"

func main() {
	fmt.Printf("Starting app %v\n", appName)
	initializeBoltClient()
	services.StartWebServer(port)
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}

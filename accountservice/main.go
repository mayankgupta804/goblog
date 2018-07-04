package main

import (
	"fmt"

	service "github.com/go_microservices/goblog/services"
)

var appName = "accountservice"
var port = "8000"

func main() {
	fmt.Printf("Starting app %v\n", appName)
	service.StartWebServer(port)
}

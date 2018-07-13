package services

import (
	"net/http"

	"github.com/go_microservices/goblog/accountservice/service"
)

// Route contains information about the name of the route, an HTTP method and a pattern the function
// will execute when this route is called
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is a slice of type Route
type Routes []Route

// Initialize routes
var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		service.GetAccount,
	},
}

package services

import (
	"github.com/gorilla/mux"
)

// NewRouter returns a pointer to mux.Router we can use as a handler
func NewRouter() *mux.Router {

	// Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	// Iterate orouteer routes that are declared in routes file and attach them to our router instance
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

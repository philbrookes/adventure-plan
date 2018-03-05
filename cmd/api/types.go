package main

import (
	"github.com/gorilla/mux"
)

//APIHandler defines an interface for a handler used by the API
type APIHandler interface {
	ConfigureRouter(router *mux.Router)
}

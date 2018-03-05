package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/philbrookes/adventure-plan/pkg/config"
	"github.com/philbrookes/adventure-plan/pkg/maps"
	"github.com/philbrookes/adventure-plan/pkg/users"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	headersOk := handlers.AllowedHeaders(config.GetConfig(os.Getenv("ENV")).GetAllowedHeaders())
	originsOk := handlers.AllowedOrigins(config.GetConfig(os.Getenv("ENV")).GetAllowedOrigins())
	methodsOk := handlers.AllowedMethods(config.GetConfig(os.Getenv("ENV")).GetAllowedMethods())

	users.NewUsersHandler().ConfigureRouter(router.PathPrefix("/users").Subrouter())
	maps.NewMapsHandler().ConfigureRouter(router.PathPrefix("/maps").Subrouter())

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("/usr/src/web/public/")))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(config.GetConfig(os.Getenv("ENV")).GetPortListenerStr(), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

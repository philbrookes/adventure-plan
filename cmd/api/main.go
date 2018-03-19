package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/philbrookes/adventure-plan/pkg/config"
	"github.com/philbrookes/adventure-plan/pkg/maps"
	"github.com/philbrookes/adventure-plan/pkg/mysql"
	"github.com/philbrookes/adventure-plan/pkg/users"
)

func main() {
	db, err := mysql.GetDB(config.GetConfig(os.Getenv("ENV")).GetMySQLConfig())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	err = mysql.InitDB(db, config.GetConfig(os.Getenv("ENV")).GetMySQLConfig())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	router := mux.NewRouter().StrictSlash(true)

	headersOk := handlers.AllowedHeaders(config.GetConfig(os.Getenv("ENV")).GetAllowedHeaders())
	originsOk := handlers.AllowedOrigins(config.GetConfig(os.Getenv("ENV")).GetAllowedOrigins())
	methodsOk := handlers.AllowedMethods(config.GetConfig(os.Getenv("ENV")).GetAllowedMethods())

	users.NewUsersHandler().ConfigureRouter(router.PathPrefix("/users").Subrouter())
	maps.NewMapsHandler(db).ConfigureRouter(router.PathPrefix("/maps").Subrouter())

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("/usr/src/web/public/")))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(config.GetConfig(os.Getenv("ENV")).GetPortListenerStr(), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

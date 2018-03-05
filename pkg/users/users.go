package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//EchoRes echo response object
type EchoRes struct {
	Message string `json:"message"`
}

//Users Hander
type Users struct {
}

//NewUsersHandler creates a new users handler
func NewUsersHandler() *Users {
	return &Users{}
}

//ConfigureRouter to handle user related routes
func (u *Users) ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/", checkAuth)
	router.HandleFunc("/echo", echo)
}

func checkAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/JSON")
	json.NewEncoder(w).Encode(nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/JSON")
	er := EchoRes{Message: "test"}
	json.NewEncoder(w).Encode(er)
}

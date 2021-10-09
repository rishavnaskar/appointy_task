package routers

import (
	"github.com/gorilla/mux"
	"github.com/Sanskrita2001/insta/controllers"
)

func Router()  *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/users",controllers.CreateUser).Methods("POST")

	return r
}
package routers

import (
	"github.com/gorilla/mux"
	"github.com/rishavnaskar/insta/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetAUser).Methods("GET")

	r.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	r.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", controllers.GetAPost).Methods("GET")

	return r
}

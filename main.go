package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sanskrita2001/insta/routers"
)

func main() {
	fmt.Println("MongoDB API")
	r := routers.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening at port 5000 ...")
}

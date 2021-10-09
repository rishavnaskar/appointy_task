package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rishavnaskar/insta/routers"
)

func main() {
	fmt.Println("MongoDB API")
	r := routers.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Println("Listening at port 5000 ...")
}

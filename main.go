package main

import (
	"log"
	"net/http"

	"github.com/Harits514/BahasaKitaBETes/App"
)

func main() {
	log.Println("server stated at localhost/8080")
	http.ListenAndServe(":8080", App.Router.Route)
}

package App

import (
	"github.com/gorilla/mux"
)

//RouterStat : struct isi variabel router
type RouterStat struct {
	Route *mux.Router
}

//Router : u/ inisialisasi ruter
var Router RouterStat

func init() {
	Router.Route = mux.NewRouter()
	Router.Route.HandleFunc("/", InsertData).Methods("POST")
	Router.Route.HandleFunc("/", LoadData).Methods("GET")
	Router.Route.HandleFunc("/", UpdateData).Methods("PATCH")
}

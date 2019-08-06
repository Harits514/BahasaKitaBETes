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
	Router.Route.HandleFunc("/insertData", InsertData)
	Router.Route.HandleFunc("/loadData", LoadData)
	Router.Route.HandleFunc("/patchData", UpdateData)
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	confObj "./config"
	daoObject "./models/dao"
	mdbRouter "./routers"
	classSocket "./websocket"
)

var dao = daoObject.RDB_DAO{}
var config = confObj.Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	fmt.Println("Monitor")
	RouteFunc()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
	fmt.Fprintf(w, "Hello World")
}

func Stats(w http.ResponseWriter, r *http.Request) {
	ws, err := classSocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go classSocket.Writer(ws)
}

func RouteFunc() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePage)
	r.HandleFunc("/stats", Stats)
	r.HandleFunc("/api/v1/monitor", mdbRouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/monitor/{id}", mdbRouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/monitor", mdbRouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/monitor/{id}", mdbRouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/monitor/{id}", mdbRouter.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}

package main

import (
	"./websocket"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Monitor")
	fazRotas()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}

func fazRotas() {
	http.HandleFunc("/", stats)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


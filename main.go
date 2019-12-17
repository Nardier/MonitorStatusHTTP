package websocket

import (
	ClassSocket "./websocket"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var messagChannel = make(chan *string)

func main() {
	fmt.Println("Monitor")
	routeFunc()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := ClassSocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	clients[ws] = true

	go ClassSocket.Writer(ws)
}

func routeFunc() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

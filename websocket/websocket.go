package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func fazRequisicao() Retorno {
	url := "https://www.google.com.br"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	start := time.Now()
	res, _ := http.DefaultClient.Do(req)
	var response Retorno
	response.Latencia = time.Since(start).Seconds()
	response.Resposta = res.Status
	return response
}


func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(500 * time.Millisecond)

		for t := range ticker.C {
			fmt.Printf("Updating Stats: %+v\n", t)
			jsonString, err := json.Marshal(fazRequisicao())
			if err != nil {
				fmt.Println(err)
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

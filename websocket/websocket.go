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

func doRequest() ResponseDataBody {
	var returnData ResponseDataBody
	url := "https://www.google.com.br"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	start := time.Now()
	res, _ := http.DefaultClient.Do(req)
	returnData.Response_Time = time.Since(start).Seconds()
	returnData.Response_Status = res.Status
	returnData.Response_StatusCode = res.StatusCode
	fmt.Println(res.StatusCode)
	returnData.Request_Url = url
	return returnData
}

func emitMessage() {
	for {
		val := <-messagChannel
		latlong := fmt.Sprintf("Mensagem")
		// send to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(latlong))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func senderMessage(w http.ResponseWriter, r *http.Request) {
	var message string
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}
	defer r.Body.Close()
	go writer(&message)
}

func Writer(conn *websocket.Conn) {
	start := time.Now()

	var returnData ResponseDataBody
	for {
		ticker := time.NewTicker(time.Second)
		for t := range ticker.C {
			fmt.Printf("Updating Stats: %+v\n", t)
			//TODO: Converter "t" para retornar
			returnData = doRequest()
			returnData.TEOM = fmt.Sprintf("%f", time.Since(start).Seconds())[0:4]
			jsonString, err := json.Marshal(returnData)
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

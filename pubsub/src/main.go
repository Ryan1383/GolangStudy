package main

import (
	"fmt"
	"log"
	"net/http"
	"pubsub/pubsub"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func autoId() string {

	uuid.Must(uuid.NewV4()).String()
}

var ps = &pubsub.PubSub{}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("New client is connected")

	client := pubsub.Client{
		Id:         autoId(),
		Connection: conn,
	}

	// add this client into the list

	ps.AddClient(client)

	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println("Something went wrong", err)
			return
		}

		aMessage := []byte("Hi Client I'm Server")

		if err := conn.WriteMessage(messageType, aMessage); err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("New message from Client : %s", string(p))
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "static")
	})

	http.HandleFunc("/ws", webSocketHandler)

	http.ListenAndServe(":3000", nil)

	fmt.Println("Server is running : http://localhost:3000")
}

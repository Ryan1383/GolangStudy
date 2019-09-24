package pubsub

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type PubSub struct {
	Client []Client
}

type Client struct {
	Id         string
	Connection *websocket.Conn
}

func (ps *PubSub) AddClient(client Client) *PubSub {

	ps.Client = append(ps.Client, client)

	fmt.Println("adding new client to the list ", client.Id)
	return ps
}

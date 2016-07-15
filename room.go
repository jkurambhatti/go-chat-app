package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Room struct {
	Id string

	Join chan *Client // Join is a channel of Clients who want to join the room

	Forward chan []byte

	LiveConns map[*Client]bool
}

func (r *Room) joinRoom(c *Client) {
	fmt.Println("client added to room")
	// r.Join <- c // add client to the Room
	// for cli := range r.Join {
	// 	fmt.Println(cli.UUId)
	// }

	fmt.Println("leaving joinRoom()")
}

var upgrader = &websocket.Upgrader{}

// func NewRoom(s string) *Room {
// 	return &Room{Id: "geekSkool", Join: make(chan *Client), LiveConns: make(map[*Client]bool)}
// }

var globalRoom = Room{Id: "geekSkool",
	Join:      make(chan *Client),
	LiveConns: make(map[*Client]bool),
	Forward:   make(chan []byte)}

func newConn(w http.ResponseWriter, r *http.Request) {
	var newsock, err = upgrader.Upgrade(w, r, nil)
	fmt.Println("err :", err)
	if err == nil {
		fmt.Println("socket succesfull")
	}

	c := NewClient() // create a new client
	c.sock = newsock // attach newly created socket to client

	globalRoom.Join <- c // adding newly created client to Room

/*	go func(cn *Client) {
		fmt.Println("launched a new go routine")
		for {
			cn.sock.WriteMessage(1, []byte("hey thererererer"))
			_, msg, _ := cn.sock.ReadMessage()
			fmt.Println("received a new message")
			fmt.Println(string(msg))
			cn.room.Forward <- msg
		}
	}(c)
*/

	go c.write()
	c.read()
	defer recover()

}

// broadcast message to all the clients in the room
func (r *Room) Broadcast(message []byte) {
	fmt.Println("entered broadcast")
	for cli := range r.LiveConns {
		fmt.Println("clientId: ",cli.UUId)
		cli.send <- message
	}
}

func (r *Room) run() {
	fmt.Println("started room.Run")
	for {
		select {
		case c := <-r.Join:
			r.LiveConns[c] = true
			fmt.Println("add entry of the new client in the LiveConns")
		case msg := <-r.Forward:
			fmt.Println("broadcast to all the clients")
			r.Broadcast(msg)
		}
	}
}

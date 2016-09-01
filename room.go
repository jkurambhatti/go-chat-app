package main

import (
	"fmt"
	"net/http"

	"github.com/auth0/auth0-golang/examples/regular-web-app/app"
	"github.com/gorilla/websocket"
)

type Room struct {
	Id string // Room Id

	Join chan *Client //  Clients who want to join the room

	Leave chan *Client // clients who want to leave the room

	Forward chan []byte // any client sending a message is collected in this Forward channel

	LiveConns map[*Client]bool // keeps a track of all the connections present
}

var upgrader = &websocket.Upgrader{}

// func NewRoom(s string) *Room {
// 	return &Room{Id: "geekSkool", Join: make(chan *Client), LiveConns: make(map[*Client]bool)}
// }

// currently only single room with Id "DefaultRoom"
var globalRoom = Room{Id: "DefaultRoom",
	Join:      make(chan *Client),
	Leave:     make(chan *Client),
	LiveConns: make(map[*Client]bool),
	Forward:   make(chan []byte)}

// http handler to create a socket connection
func newConn(w http.ResponseWriter, r *http.Request) {
	var newsock, err = upgrader.Upgrade(w, r, nil) // upgrade http connection to websocket
	var profile interface{}
	fmt.Println("err :", err)
	if err == nil {
		fmt.Println("socket succesful")
	} else {
		fmt.Println("socket unsuccesful")
		return
	}

	session, _ := app.GlobalSessions.SessionStart(w, r)
	// defer session.SessionRelease(w)
	//
	// if session == nil {
	// 	fmt.Println("session is nil")
	// 	// http.Redirect(w, r, "/", http.StatusMovedPermanently)
	// } else {
	// 	fmt.Println("session is not nil")
	// }

	c := NewClient() // create a new client
	c.sock = newsock // attach newly created socket to client
	//
	profile = session.Get("profile")
	if profile == nil {
		fmt.Println("profile is nil")
		// http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		fmt.Println("profile is not nil")

		a := profile.(map[string]interface{})
		c.UUId = a["nickname"]
		fmt.Println(a["nickname"].(string))
	}

	globalRoom.Join <- c // adding newly created client to Room

	go c.write()    // create a go routine (thread) to write
	c.read()        // keep waiting for any message to be received
	defer recover() // recover in case of panic

}

// broadcast message to all the clients in the room
func (r *Room) Broadcast(message []byte) {
	fmt.Println("entered broadcast")

	// range over all the clients of the Room and send message to each
	for cli := range r.LiveConns {
		fmt.Println("clientId: ", cli.UUId)
		cli.send <- message // write message on the send channel of the client
	}
}

// keep the room running in the background concurrently
// infinitely wait for an event to take place on the channel
func (r *Room) run() {
	fmt.Println("started room.Run")
	for {
		select {
		case c := <-r.Join: // new client joins
			r.LiveConns[c] = true // add it to the LiveConns
			fmt.Println("add entry of the new client in the LiveConns")
		case msg := <-r.Forward: // save the message sent by a client on the room and broadcast.
			fmt.Println("broadcast to all the clients")
			r.Broadcast(msg)
		case lc := <-r.Leave: // get the client who wants to leave
			fmt.Println("r.Leave : lc.UUid", lc.UUId)
			r.LiveConns[lc] = false
			delete(r.LiveConns, lc)                 // delete from the list of liveConns
			if err := lc.sock.Close(); err != nil { // close the socket
				fmt.Println("error deleting client and closing client socket :", err)
			}
		}
	}
}

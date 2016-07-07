package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	// a socket connection to connect the client to the server
	socket *websocket.Conn
	// a channel to send the data
	data chan []byte
	// room to know which room the client belongs to
	room *room
}

func (c *Client) read() {
	for {
		msgtype, msg, err := c.socket.ReadMessage()
		if err != nil {
			fmt.Errorf("client.read : %s", err)
			break
		}
		defer c.socket.Close() // close the socket connection after reading
		if msgtype == 2 {      // check if the msg is in text format or binary
			fmt.Println("message type is text")
		}
		c.room.forward <- msg // send the recieved message to the forward channel of the room
	}
}

func (c *Client) write() {
	for m := range c.data {
		if err := c.socket.WriteMessage(websocket.TextMessage, m); err != nil {
			break
		}
	}
	c.socket.Close()
}

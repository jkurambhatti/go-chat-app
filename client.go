package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

var uniqueId = 0

type Client struct {
	UUId int `json:"name"`

	sock *websocket.Conn

	send chan []byte

	room *Room
}

func (c *Client) read() {

	var msgstring string
	for{
		mtp, msg, err := c.sock.ReadMessage()
		msgstring = fmt.Sprintf("client %d : %s",c.UUId,string(msg))
		if err != nil {
			fmt.Errorf("c.sock.ReadMessage: %s\n", err)
		}
		c.room.Forward <- []byte(msgstring)

		fmt.Println(mtp, string(msg))
	//	fmt.Println(mtp, msg)
	}
}

func (c *Client) write() {
	for m := range c.send {
		if err := c.sock.WriteMessage(websocket.TextMessage,m); err != nil {
			break
		}
	}
}

func NewClient() *Client {
	uniqueId++
	fmt.Println("new Incoming Client", uniqueId)
	return &Client{UUId: uniqueId, send: make(chan []byte), room: &globalRoom}
}

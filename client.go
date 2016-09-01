package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	UUId interface{} //unique Id

	sock *websocket.Conn // socket for the client

	send chan []byte // channel for recieving the messages

	room *Room // room to which it belongs
}

// infinitely keep reading the socket for an incoming message
// when received forward it to the room to broadcast
func (c *Client) read() {

	var msgstring string

	for {
		fmt.Println("start sock.Readmessage : ")

		mtp, msg, err := c.sock.ReadMessage()
		fmt.Println("sock.Readmessage : ", "mtype : ", mtp, "err :", err)
		if mtp == -1 {
			fmt.Println("sending to room.Leave to be closed")

			// globalRoom.Leave <- c
			c.sock.Close()
		}

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Println("client ", c.UUId, " has closed")
			}
		}

		msgstring = fmt.Sprintf(" %v : %s", c.UUId, string(msg))
		if err != nil {
			fmt.Errorf("c.sock.ReadMessage: %s\n", err)
		}
		c.room.Forward <- []byte(msgstring)

		fmt.Println(mtp, string(msg))
	}
}

// waits endlessly on the channel to receive the message
// when received write it to the socket
func (c *Client) write() {
	for m := range c.send {
		if err := c.sock.WriteMessage(websocket.TextMessage, m); err != nil {
			break
		}
	}
}

func NewClient() *Client {
	fmt.Println("new Incoming Client")
	return &Client{send: make(chan []byte), room: &globalRoom}
}

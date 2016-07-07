package main

type room struct {
	// forward is just normal forward command we use to forward
	// the existing buffered messages to the other clients
	forward chan []byte
	// join is a channel for clients who want to join the room
	join chan *Client

	// leave is a channel for clients who want to leave the room
	leave chan *Client

	// clients keeps a track of who all clients are present in the room
	clients map[*Client]bool
}

// the work of the run function is to keep track of the different channels
// and perform appropriate tasks based on the input recieved
func (r *room) run() {
	for {
		select {
		case message := <-r.forward:
			// send to all clients who have a boolean true
			for cli, _ := range r.clients {
				select {
				// send the message to data channel of every client
				case cli.data <- message:
				default:
					delete(r.clients, cli)
				}
				close(cli.data)
			}
		case jc := <-r.join:
			// add the jc to the index of clients and set the boolean value to true
			r.clients[jc] = true
		case lc := <-r.leave:
			// remove the lc from the clients list
			delete(r.clients, lc)
			close(lc.data)
		}
	}
}

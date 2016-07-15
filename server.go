package main

import (
	// "github.com/gorilla/websocket"
	"fmt"
	"html/template"
	"net/http"
)

var clientNames []string

var homeTemplate = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "i.html")

		w.Header().Set("Content-Type", "text/html")
		homeTemplate.Execute(w, r.Host)
		fmt.Println(r.Host)

		// http.ServeFile(w, r, "register.html")
	})

	go globalRoom.run()

	http.HandleFunc("/ws", newConn)
	http.ListenAndServe(":8080", nil)
}

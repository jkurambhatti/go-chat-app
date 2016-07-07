package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("a simple chat application"))
}

func main() {

	http.HandleFunc("/", index)
	fmt.Println("Listening on Port : 3000 ")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("ListenAndServe : %s ", err)
	}
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func index(w http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Errorf("template.ParseFiles %s \n", err)
	}
	u := User{Name: "Jayesh"}
	if err = tmp.Execute(w, u); err != nil {
		log.Println("tmp.Execute: ", err)
		return
	}

	w.Write([]byte("a simple chat application"))
}

func main() {

	http.HandleFunc("/", index)
	fmt.Println("Listening on Port : 3000 ")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("ListenAndServe : %s ", err)
	}
}

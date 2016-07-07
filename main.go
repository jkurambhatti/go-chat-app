// in this project we will build a chat application which will allow multiple users to communicate in real-time

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	fileName string
	tmpl     *template.Template
}

type User struct {
	Name string `json:"name"`
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.fileName)))
	})

	u := User{Name: "Jayesh"}
	if err := t.tmpl.Execute(w, u); err != nil {
		log.Println("t.tmpl.Execute: ", err)
		return
	}

	w.Write([]byte("a simple chat application"))
}

func main() {
	var t templateHandler
	t.fileName = "index.html"
	http.Handle("/", &t)
	fmt.Println("Listening on Port : 3000 ")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("ListenAndServe : %s ", err)
	}
}

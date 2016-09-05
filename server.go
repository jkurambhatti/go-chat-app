package main

import (
	"fmt"
	"net/http"

	"github.com/auth0/auth0-golang/examples/regular-web-app/routes/callback"
	"github.com/auth0/auth0-golang/examples/regular-web-app/routes/home"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jkurambhatti/go-chat-app/routes/middlewares"
	"github.com/jkurambhatti/go-chat-app/routes/user"
)


func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.HandleFunc("/chat", ChatHandler)
	r.HandleFunc("/ws", newConn)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "chat.html")
	fmt.Println("server.go :", r.Host)
}

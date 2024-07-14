package app

import (
	"log"
	"net/http"
	"parking/cmd/api/app/dependences"

	"github.com/gorilla/mux"
)

func New() {
	handlerContainer := dependences.NewWire()
	routes(&handlerContainer)
}

func routes(container *dependences.HandlerContainer) {
	mux := mux.NewRouter()

	mux.HandleFunc("/users", container.HandlerUser.CreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", container.HandlerUser.GetUser).Methods("GET")
	mux.HandleFunc("/users/{id}", container.HandlerUser.UpdateUser).Methods("PUT")
	mux.HandleFunc("/users", container.HandlerUser.GetAllUsers).Methods("GET")

	log.Print("Run Server: localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

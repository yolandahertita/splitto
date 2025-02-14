package main

import (
	"log"
	"net/http"
	"splitto/app/handler"
	"splitto/app/model"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// initialize database model
	databaseURL := "postgres://postgres:1234@localhost:5432/postgres?application_name"
	model := &model.Model{}
	model.ConnectDB(databaseURL)

	// initialize route handler
	handler := &handler.Handler{Model: model}

	// register routes
	router := mux.NewRouter()
	RegisterRoutes(router, handler)

	// start server
	StartServer(router, "127.0.0.1:8000")
}

func RegisterRoutes(router *mux.Router, handler *handler.Handler) {
	router.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/test", handler.TestUser).Methods("POST")
}

func StartServer(router *mux.Router, addr string) {
	server := &http.Server{
		Handler: router,
		Addr:    addr,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

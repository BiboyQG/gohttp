package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	log.Println("Starting server on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

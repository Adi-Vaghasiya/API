package main

import (
	"log"
	"net/http"

	"jwt-api/handler"
	"jwt-api/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestLogger)
	//r.Post("/login", handler.LoginHandler)
	r.HandleFunc("/login", handler.LoginHandler)
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Get("/profile", handler.ProfileHandler)
	})
	Port := 8080
	serverURL := "http://localhost:8080"
	log.Printf("Server Started At URL %v and Port %v", serverURL, Port)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

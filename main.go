package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	PORT := os.Getenv("PORT")

	fmt.Println("Hello PORT is starting on ", PORT)

	if PORT == "" {
		log.Fatal("PORT is not provided")
	}

	router := chi.NewRouter()

	_ = cors.Handler

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/ready", handlerReadiness)
	router.Mount("/v1", v1Router)

	v1Router.Get("/err", handlerError)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT : ", PORT)
}

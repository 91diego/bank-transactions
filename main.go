package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/91diego/bank-transactions/handlers"
	"github.com/91diego/bank-transactions/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file: " + err.Error())
	}

	PORT := os.Getenv("PORT")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// create new server
	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func BindRoutes(s server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/transaction-summary", handlers.InsertTransactions(s)).Methods(http.MethodPost)
}

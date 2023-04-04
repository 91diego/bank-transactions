package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/91diego/bank-transactions/database"
	trxRepo "github.com/91diego/bank-transactions/repositories"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port        string // puerto a ejecutar
	DatabaseUrl string // conexion a base de datos
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("DB Url is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {

	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)

	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	trxRepo.SetRepository(repo)
	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

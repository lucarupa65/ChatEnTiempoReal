package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/luisk6510/ChatEnTiempoReal/handlers"
	"github.com/luisk6510/ChatEnTiempoReal/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("No se encontro el archivo .env")
	}

	Port := os.Getenv("Port")
	JWTSecret := os.Getenv("JWTSecret")
	DatabaseURL := os.Getenv("DatabaseURL")

	s, err := server.NewServe(context.Background(), &server.Config{
		Port:        Port,
		JWTSecret:   JWTSecret,
		DatabaseURL: DatabaseURL,
	})

	if err != nil {
		log.Fatal(err)
	}
	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandlers(s)).Methods(http.MethodGet)
}

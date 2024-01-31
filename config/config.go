package config

import (
	"log"
	"os"
)

type ServerConfig struct {
	Port      string
	JWTSecret string
}

type DataBaseConfig struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

type ConfiInterna struct {
	Serve ServerConfig
	DB    DataBaseConfig
}

func NewConfig() *ConfiInterna {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("El puerto para el servidor es requerido")
	}
	JWTSecret := os.Getenv("JWTSecret")
	if JWTSecret == "" {
		log.Fatal("La llave de incriptacion es requerido")
	}
	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		log.Fatal("El usaurio para la base de datos es requerido")
	}
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		log.Fatal("la contrasena de la base de datos es requerido")
	}
	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		log.Fatal("El host de la base de datos es requerido")
	}
	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		log.Fatal("El puerto para la base de datos es requerido")
	}
	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		log.Fatal("El nombre de la base de datos es requerido")
	}
	confi := &ConfiInterna{
		Serve: ServerConfig{
			Port:      PORT,
			JWTSecret: JWTSecret,
		},
		DB: DataBaseConfig{
			DB_USER:     DB_USER,
			DB_PASSWORD: DB_PASSWORD,
			DB_HOST:     DB_HOST,
			DB_PORT:     DB_PORT,
			DB_NAME:     DB_NAME,
		},
	}
	return confi
}

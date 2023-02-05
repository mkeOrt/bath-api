package config

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	Server struct {
		Port int
	}
	Database struct {
		Name     string
		User     string
		Password string
	}
	JWT struct {
		SecretKey string
	}
	BCrypt struct {
		Cost int
	}
}

var C config

func ReadConfig() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Panic("error reading env variables")
	}

	bcryptCost, err := strconv.Atoi(os.Getenv("B_CRYPT_COST"))
	if err != nil {
		log.Panic("error reading env variables")
	}

	C = config{
		Server: struct{ Port int }{
			Port: port,
		},
		Database: struct {
			Name     string
			User     string
			Password string
		}{
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		JWT: struct{ SecretKey string }{
			os.Getenv("JWT_SECRET_KEY"),
		},
		BCrypt: struct{ Cost int }{
			Cost: bcryptCost,
		},
	}
}

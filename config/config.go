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
		Name string
	}
}

var C config

func ReadConfig() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Panic("error reading env variables")
	}

	C = config{
		Server: struct{ Port int }{
			Port: port,
		},
		Database: struct{ Name string }{
			Name: os.Getenv("DATABASE_NAME") + ".sqlite",
		},
	}
}

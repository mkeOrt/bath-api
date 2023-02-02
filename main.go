package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/mkeort/bath-hexagonal/config"
	"github.com/mkeort/bath-hexagonal/infrastructure/datastore"
	"github.com/mkeort/bath-hexagonal/infrastructure/router"
	"github.com/mkeort/bath-hexagonal/infrastructure/server"
	"github.com/mkeort/bath-hexagonal/registry"
)

func init() {
	config.ReadConfig()
}

func main() {
	db := datastore.NewDB()

	r := registry.NewRegistry(db)

	server := server.NewServer()
	router.NewRouter(server.App.Group("/api"), r.NewAppController())
	server.Listen()
}

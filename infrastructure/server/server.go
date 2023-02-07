package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mkeort/bath-hexagonal/config"
)

type server struct {
	App *fiber.App
}

func NewServer() *server {
	s := server{
		App: fiber.New(),
	}
	s.App.Use(cors.New())

	return &s
}

func (s *server) Listen() {
	s.App.Listen(fmt.Sprintf(":%d", config.C.Server.Port))
}

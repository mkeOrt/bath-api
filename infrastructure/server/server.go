package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mkeort/bath-hexagonal/config"
)

type server struct {
	App *fiber.App
}

func NewServer() *server {
	return &server{
		App: fiber.New(),
	}
}

func (s *server) Listen() {
	s.App.Listen(fmt.Sprintf(":%d", config.C.Server.Port))
}

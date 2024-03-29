package registry

import (
	"github.com/mkeort/bath-hexagonal/interface/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Middleware: controller.NewAppMiddleware(r.db),
		User:       r.NewUserController(),
		Poop:       r.NewPoopController(),
	}
}

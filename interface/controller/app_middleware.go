package controller

import (
	"gorm.io/gorm"
)

type appMiddleware struct {
	Auth AuthMiddleware
}

type AppMiddleware interface {
	GetAuthMiddlewares() AuthMiddleware
}

func NewAppMiddleware(db *gorm.DB) AppMiddleware {
	return &appMiddleware{Auth: NewAuthMiddleware(db)}
}

func (am *appMiddleware) GetAuthMiddlewares() AuthMiddleware {
	return am.Auth
}

package controller

type AppController struct {
	Middleware interface{ AppMiddleware }
	User       interface{ UserController }
	Poop       interface{ PoopController }
}

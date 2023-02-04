package controller

type AppController struct {
	User interface{ UserController }
	Poop interface{ PoopController }
}

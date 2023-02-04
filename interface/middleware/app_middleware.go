package middleware

type AppMiddleware struct {
	Auth interface{ AuthMiddleware }
}

func NewAppMiddleware() AppMiddleware {
	return AppMiddleware{
		Auth: NewAuthMiddleware(),
	}
}

package httpserver

type Handler interface {
	GetRoutes() []*RouteHandler
}

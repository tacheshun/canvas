package server

import "github.com/tacheshun/canvas/handlers"

func (s Server) setupRoutes() {
	handlers.Health(s.mux)
}

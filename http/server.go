package http

import "github.com/gin-gonic/gin"

type Server struct {
	handlers Handlers
}

func NewServer(handlers Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) ListenAndServe(address string) error {
	router := gin.Default()
	router.GET("/", s.handlers.RecipieHandlers.GetRecipie)
	return router.Run(address)
}
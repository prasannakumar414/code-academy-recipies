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
	router.POST("/", s.handlers.RecipieHandlers.CreateRecipie)
	router.DELETE("/", s.handlers.RecipieHandlers.DeleteRecipie)
	router.PUT("/", s.handlers.RecipieHandlers.UpdateRecipie)
	return router.Run(address)
}

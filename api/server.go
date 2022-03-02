package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	AcccountRoutes(router)
	TransactionRoutes(router)
	UserRoutes(router)
	server.Router = router

	return server
}

func (s *Server) Serve() error {
	return s.Router.Run()
}

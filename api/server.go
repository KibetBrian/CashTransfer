package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func  (s *Server) Serve(port string) *Server{
	server := &Server{
		Router: gin.Default(),
	}
	AcccountRoutes(server.Router)
	TransactionRoutes(server.Router)
	UserRoutes(server.Router)
	s.Start(port)	
	return server
}

func (server *Server) Start(port string) error{
	return server.Router.Run(port)
}
package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	router := gin.Default()
	server := &Server{
		router: router,
	}

	server.generateRoutes()
	return server, nil
}

func (server *Server) Start() {
	err := server.router.Run("0.0.0.0:7600")
	if err != nil {
		fmt.Println("unable to start the server", err.Error())
		return
	}
	fmt.Println("Server started at port 7600.")
}

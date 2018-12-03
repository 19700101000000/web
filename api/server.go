package main

import (
	"github.com/19700101000000/system-sample/api/handler"
	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
}

func newServer() *server {
	return &server{
		router: gin.Default(),
	}
}
func (s *server) serverInit() {
	s.router.GET("/", handler.Index)

	/* OAuth2 from GOOGLE */
	s.router.GET("/auth/google/login", handler.AuthLogin)
	s.router.GET("/auth/google/callback", handler.AuthCallback)
}
func (s *server) serverRun() {
	s.router.Run(":8080")
}

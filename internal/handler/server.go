package handler

import "github.com/gin-gonic/gin"

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(addr string) error {
	router := gin.New()

	router.GET("/:id", redirect)
	router.POST("/", shortURL)

	return router.Run()
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muradrmagomedov/go-shortener/internal/service"
)

type Handler struct {
	Service service.Shortener
}

func NewHandler(service service.Shortener) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) Run(addr string) error {
	router := gin.New()

	router.GET("/:id", h.redirect)
	router.POST("/", h.saveURL)

	return router.Run()
}

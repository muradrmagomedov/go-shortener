package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"error"`
}

func SendError(c *gin.Context, code int, message string, funcName string) {
	error := errorResponse{Message: message}
	logrus.Errorf("%s:%s", funcName, message)
	c.JSON(code, error)
}

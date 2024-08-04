package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const addr = "http://localhost:8080/"
const id = "EwHXdJfB"
const location = "https://practicum.yandex.ru/"

// shortURL(c *gin.Context) принимает url-строку в формате text/plain и возвращает адрес, с которого будет осуществляться редирект
func shortURL(c *gin.Context) {
	const funcName = "handler.getURL"

	ct := c.Request.Header["Content-Type"]
	var url string
	//Проверяем content-type
	if len(ct) == 0 {
		SendError(c, http.StatusBadRequest, "Отсутствует Content-Type", funcName)
		return
	}
	if ct[0] != "text/plain" {
		SendError(c, http.StatusBadRequest, "Content-Type должен быть text/plain", funcName)
		return
	}

	data, err := c.GetRawData()
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error(), funcName)
		return
	}

	url = string(data)
	//Проверяем, что url не пустой
	if len(url) == 0 {
		SendError(c, http.StatusBadRequest, "Пустая ссылка", funcName)
		return
	}
	response := addr + id
	c.String(http.StatusCreated, response)
}

func redirect(c *gin.Context) {
	const funcName = "handler.redirect"
	paramId := c.Param("id")
	if paramId != id {
		SendError(c, http.StatusBadRequest, "Не удалось найти URL", funcName)
		return
	}
	c.Writer.Header().Add("Content-Type", "text/plain")
	c.Writer.Header().Add("Location", location)
	c.String(http.StatusTemporaryRedirect, "")
}

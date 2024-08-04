package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO вынести в конфиг
const addr = "http://localhost:8080/"

// shortURL(c *gin.Context) принимает url-строку в формате text/plain и возвращает адрес, с которого будет осуществляться редирект
func (h *Handler) saveURL(c *gin.Context) {
	const funcName = "handler.getURL"

	ct := c.Request.Header["Content-Type"]
	var url string
	//Проверяем content-type
	if len(ct) == 0 {
		SendError(c, http.StatusBadRequest, "Отсутствует Content-Type", funcName)
		return
	}
	// if ct[0] != "text/plain" {
	// 	SendError(c, http.StatusBadRequest, "Content-Type должен быть text/plain", funcName)
	// 	return
	// }

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
	alias, err := h.Service.SaveURL(url)
	if err != nil {
		SendError(c, http.StatusInternalServerError, err.Error(), funcName)
		return
	}
	response := addr + alias
	c.String(http.StatusCreated, response)
}

func (h *Handler) redirect(c *gin.Context) {
	const funcName = "handler.redirect"
	paramId := c.Param("id")
	url, err := h.Service.GetURL(paramId)
	if err != nil {
		SendError(c, http.StatusBadRequest, "Не удалось найти URL", funcName)
		return
	}
	c.Writer.Header().Add("Content-Type", "text/plain")
	c.Writer.Header().Add("Location", url)
	c.String(http.StatusTemporaryRedirect, "")
}

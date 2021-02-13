package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *HTTPHandler) Private(context echo.Context) error {
	//log.Println(goth.Session.Marshal())
	log.Println("Private")
	//h.repository.ExistsUser()
	return context.JSON(http.StatusOK, "You are authenticated!")
}


func (h *HTTPHandler) Public(context echo.Context) error {
	//log.Println(goth.Session.Marshal())
	log.Println("Private")
	//h.repository.ExistsUser()
	return context.JSON(http.StatusOK, "request public data")
}
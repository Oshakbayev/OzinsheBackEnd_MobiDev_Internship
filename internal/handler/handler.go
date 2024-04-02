package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"ozinshe/pkg/service"
)

type Handler struct {
	log *log.Logger
	svc service.SvcInterface
}

func CreateHandler(service service.SvcInterface, log *log.Logger) Handler {
	return Handler{svc: service, log: log}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	r.GET("/home", h.HomePageHandler)
}

package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"ozinshe/pkg/entity"
	"ozinshe/pkg/service"
)

type Handler struct {
	log *log.Logger
	svc service.SvcInterface
}

func CreateHandler(service service.SvcInterface, log *log.Logger) Handler {
	return Handler{svc: service, log: log}
}

func (h *Handler) InitRoutes() *gin.Engine {
	ginServer := gin.Default()
	ginServer.GET("/home", h.HomePageHandler)

	auth := ginServer.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.GET("/verifyAccount", h.VerifyAccount)
		auth.POST("/sign-in", h.SignIn)
	}
	ginServer.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ginServer.Use(h.AuthMiddleware())
	return ginServer
}

func (h *Handler) WriteHTTPResponse(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, entity.ErrorJSONResponse{Message: msg})
}

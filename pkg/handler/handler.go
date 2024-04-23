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
	core := ginServer.Group("/core", h.AuthMiddleware())
	{
		core.POST("/movie", h.AdminRoleMiddleware(), h.CreateMovie)
		core.POST("favorites", h.CreateFavoriteMovie)
		core.GET("/home", h.HomePageHandler)
		core.GET("/movies/page", h.GetAllMovies)
		core.GET("/movie/genres", h.GetAllGenres)
		core.GET("/movie/:id", h.GetMovieById)
		core.GET("/movie/:id/season/:seasonId", h.GetMovieSeasonById)
		core.GET("/movie/:id/season/:seasonId/series/:seriesId")
		core.GET("/categories", h.GetCategories)
		core.GET("/categories/:categoryId", h.GetMovieMainsByCategory)
		core.GET("/user/profile", h.GetUserProfile)
		core.GET("/movieMain/search", h.GetMovieMainsByTitle)
		core.GET("/movieMain/search/genre", h.GetMovieMainsByGenre)
		core.GET("/favorites/", h.GetFavoriteMovies)
		core.PUT("/user/profile", h.UpdateUserProfile)
		core.PUT("/user/profile/password", h.ChangePassword)
		core.PUT("/movie/:id", h.AdminRoleMiddleware(), h.UpdateMovieById)
		core.DELETE("/movie/:id", h.AdminRoleMiddleware(), h.DeleteMovieById)
		core.DELETE("/favorites/", h.DeleteFavoriteMovies)
	}
	auth := ginServer.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.GET("/verifyAccount", h.VerifyAccount)
		auth.POST("/sign-in", h.SignIn)
	}
	ginServer.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return ginServer
}

func (h *Handler) WriteHTTPResponse(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, entity.ErrorJSONResponse{Message: msg})
}

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ozinshe/cmd/configs"
	handlers "ozinshe/internal/handler"
	"ozinshe/pkg/repository"
	"ozinshe/pkg/repository/database"
	services "ozinshe/pkg/service"
)

type Server struct {
	Address string
	log     *log.Logger
	server  *gin.Engine
}

func (s *Server) Run() error {
	s.log.Println("starting api server at http://localhost" + s.Address)
	fmt.Println("starting api server at http://localhost" + s.Address)
	return s.server.Run(s.Address)
}

func InitServer(config *configs.Config, errLogger *log.Logger) *Server {
	db := database.CreateDB(config.DBDriver)
	repo := repository.CreateRepository(db, errLogger)
	service := services.CreateService(repo, errLogger)
	handler := handlers.CreateHandler(service, errLogger)
	ginServer := gin.Default()
	handler.InitRoutes(ginServer)
	return &Server{
		Address: config.HTTPPort,
		log:     errLogger,
		server:  ginServer,
	}
}

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"io"
	"log"
	"os"
	"os/signal"
	"ozinshe/cmd/configs"
	handlers "ozinshe/pkg/handler"
	"ozinshe/pkg/logs"
	"ozinshe/pkg/repository"
	"ozinshe/pkg/repository/database"
	"ozinshe/pkg/server"
	services "ozinshe/pkg/service"
	"syscall"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func main() {
	config := configs.CreateConfig()
	if err := configs.ReadConfig("cmd/configs/config.json", &config); err != nil {
		log.Fatal(err)
	}
	ginLogFile := logs.CreateLogFile("ginLogs")
	errLogFile := logs.CreateLogFile("errLogs")
	defer logs.CloseLogFile(errLogFile)
	defer logs.CloseLogFile(ginLogFile)
	logger := logs.NewLogger(errLogFile)
	gin.DefaultWriter = io.MultiWriter(ginLogFile, os.Stdout)
	//server := servers.InitServer(&config, logger)
	db, err := database.ConnectToDB(config.DSN, logger)
	if err != nil {
		logger.Println(err)
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.CreateRepository(db, logger)
	service := services.CreateService(repo, logger)
	handler := handlers.CreateHandler(service, logger)
	srv := &server.Server{}
	go func() {
		if err := srv.InitServerAndRun(config.HTTPPort, handler.InitRoutes()); err != nil {
			logger.Printf("Emergency:Server crushed!!!")
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Printf("error occured on server shutting down: %s", err.Error())
		log.Fatal("error occured on server shutting down:", err)
	}
}

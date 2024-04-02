package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"ozinshe/cmd/configs"
	"ozinshe/pkg/logs"
	servers "ozinshe/pkg/server"
)

func main() {
	config := configs.CreateConfig()
	if err := configs.ReadConfig("cmd/configs/config.json", &config); err != nil {
		log.Fatal(err)
	}
	ginLogFile := logs.CreateLogFile("ginLogs")
	errLogFile := logs.CreateLogFile("errLogs")
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(ginLogFile)
	logger := logs.NewLogger(errLogFile)
	gin.DefaultWriter = io.MultiWriter(ginLogFile, os.Stdout)
	server := servers.InitServer(&config, logger)
	if err := server.Run(); err != nil {
		logger.Printf("Emergency:Server crushed!!!")
		log.Fatal(err)
	}
	fmt.Println("check of new .gitignore file")
}

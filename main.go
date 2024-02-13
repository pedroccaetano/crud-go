package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedroccaetano/crud-go/src/configuration/logger"
	"github.com/pedroccaetano/crud-go/src/controller"
	"github.com/pedroccaetano/crud-go/src/controller/routes"
	"github.com/pedroccaetano/crud-go/src/model/service"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init depencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	// Os roteadores podem ser inicializados de duas maneiras
	// uma usando gin.New() e a outra usando div.Default()
	//
	// gin.New() inicializa um roteador sem nenhum middleware
	//  enquanto gin.Default() inicializa o roteador com logger
	// e middlewares de recovery
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

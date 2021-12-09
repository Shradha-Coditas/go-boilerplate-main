package server

import (
	"boilerplate/config"
	"boilerplate/docs"
	"fmt"

	"boilerplate/util/logwrapper"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var router *gin.Engine

// @title Boiler Plate
// @version 1.0
// @description Golang BoilerPlate Api.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
// StartServer - start Server server.
func StartServer(serverConfig config.ServerConfig) {

	// Update Data In Swagger
	docs.SwaggerInfo.Host = fmt.Sprintf("127.0.0.1:%d", serverConfig.Port)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.BasePath = "/v1/"

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.New()

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// set logger in gin
	router.Use(logwrapper.GinLogger(), gin.Recovery())

	// Initialize the routes
	initializeRoutes()

	middlewareFunc := cors.New(cors.Options{
		AllowedOrigins:   serverConfig.AllowOrigins,
		AllowedMethods:   []string{"POST", "GET", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(time.Duration(12 * time.Hour).Seconds()),
	})
	router.Use(middlewareFunc)

	// Start serving the application
	logwrapper.Logger.Info("Running Server on port : ", serverConfig.Port)
	logwrapper.Logger.Fatal(router.Run(fmt.Sprintf(":%d", serverConfig.Port)))
}

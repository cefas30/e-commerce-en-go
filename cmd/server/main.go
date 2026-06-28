package main

import (
	"fmt"

	"ecommerce/config"
	"ecommerce/middleware"
	"ecommerce/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	router := gin.Default()

	// Middleware
	router.Use(cors.Default())
	router.Use(middleware.Logger())
	router.Use(middleware.RecoveryMiddleware())

	// Rutas
	routes.ConfigurarRutas(router)

	fmt.Println("===================================")
	fmt.Println(cfg.AppName)
	fmt.Println("Versión:", cfg.Version)
	fmt.Println("Puerto:", cfg.Port)
	fmt.Println("===================================")

	router.Run(":" + cfg.Port)

}

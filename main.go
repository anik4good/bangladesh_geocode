package main

import (
	"log"
	"os"

	Configuration "github.com/anik4good/go-echo-apiboilerplate/config"
	"github.com/anik4good/go-echo-apiboilerplate/middleware"
	"github.com/anik4good/go-echo-apiboilerplate/models"
	"github.com/anik4good/go-echo-apiboilerplate/routes"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Conf models.Config

func main() {

	Configuration.Init()

	e := echo.New()

	// Add metrics middleware
	e.Use(middleware.MetricsMiddleware)

	// Add metrics endpoint
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	routes.ConfigureRoutes(e)
	//log.Println(Configuration.Conf.Mysqlconnstring)
	println("Server is running on port:", os.Getenv("SERVER_PORT"))
	if err := e.Start(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalln(err)
	}

}

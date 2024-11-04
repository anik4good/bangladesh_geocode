package main

import (
	"log"
	"os"

	Configuration "github.com/anik4good/go-echo-apiboilerplate/config"
	"github.com/anik4good/go-echo-apiboilerplate/models"
	"github.com/anik4good/go-echo-apiboilerplate/routes"
	"github.com/labstack/echo/v4"
)

var Conf models.Config

func main() {

	Configuration.Init()

	e := echo.New()
	routes.ConfigureRoutes(e)
	//log.Println(Configuration.Conf.Mysqlconnstring)

	if err := e.Start(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalln(err)
	}

}

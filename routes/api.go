package routes

import (
	"github.com/anik4good/go-echo-apiboilerplate/app/http/controllers"
	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo) {
	e.GET("/", controllers.Hello)

	api := e.Group("/api")

	//division := api.Group("/division")

	api.GET("/divisions", controllers.Index)
	api.GET("/division/:id", controllers.Show)
	api.GET("/districts", controllers.GetAllDistricts)
	api.GET("/division/:division_name", controllers.GetAllDistrict)
	api.GET("/division/:division_name/:district_name", controllers.GetAllUpozilla)
	api.GET("/division/:division_name/:district_name/:upazila_name", controllers.GetAllUnions)
}

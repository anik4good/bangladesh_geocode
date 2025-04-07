package routes

import (
	"github.com/anik4good/go-echo-apiboilerplate/app/http/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	// Custom Rate Limiter Configuration
	rateLimiterConfig := middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStore(20), // Allow 20 requests per second per IP
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			// Use the client's IP address as the identifier
			return ctx.RealIP(), nil
		},
		ErrorHandler: func(ctx echo.Context, err error) error {
			// Custom error response when something goes wrong with the rate limiter
			return ctx.JSON(500, map[string]string{
				"message": "Internal server error",
			})
		},
		DenyHandler: func(ctx echo.Context, identifier string, err error) error {
			// Custom response when the rate limit is exceeded
			return ctx.JSON(429, map[string]string{
				"message": "Too many requests. Please try again later.",
			})
		},
	}

	// Apply the rate limiter middleware to all routes
	e.Use(middleware.RateLimiterWithConfig(rateLimiterConfig))

	// Default route
	e.GET("/", controllers.Hello)

	// Route for serving documentation
	e.GET("/docs", controllers.ServeDocs)

	// API group
	api := e.Group("/api")

	// Division routes
	api.GET("/divisions", controllers.Index)
	api.GET("/division/:id", controllers.Show)
	api.GET("/districts", controllers.GetAllDistricts)
	api.GET("/division/:division_name", controllers.GetAllDistrict)
	api.GET("/division/:division_name/:district_name", controllers.GetAllUpozilla)
	api.GET("/division/:division_name/:district_name/:upazila_name", controllers.GetAllUnions)
}

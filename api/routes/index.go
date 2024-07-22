package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Index
func Index() *echo.Echo {
	// Echo Instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blank Page")
	})

	// Routes
	UserRoute(e.Group("/users"))

	return e
}
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Define Echo Frameworks
	e:=echo.New();

	// Define Route
	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello, World!");
	});

	// Handling Route Parameter
	
	// Path Parameter
	e.GET("/hello/:username", func (c echo.Context) error  {
		username := c.Param("username");
		return c.String(http.StatusOK,"Hello," +username+"!")
		
	})
	// Query Parameter
	e.GET("/greet", func(c echo.Context) error {
		name:=c.QueryParam("name")
		if name == "" {
			name = "Guest"
		}
		return c.String(http.StatusOK, "Hello," +name+ "!")
	})



	// Start Server
	e.Start(":8080")
}
package main

import(
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

	// Start Server
	e.Start(":8080")
}
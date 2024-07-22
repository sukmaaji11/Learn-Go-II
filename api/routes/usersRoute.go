package routes

import (
	"api/controllers"

	"github.com/labstack/echo/v4"
)


func UserRoute(g *echo.Group) {
	// Create Users Route
	g.POST("/", controllers.AddUser)

	// Get Users
	g.GET("/", controllers.GetUsers)

	// Delete User
	g.DELETE("/:id", controllers.DeleteUser)
}
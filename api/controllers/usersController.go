package controllers

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var users = models.Users{
	models.User{
		ID: 1,
		Nama: "Sukma",
	},
}

// Func Add User
func AddUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	newUser := models.User {
		ID : id,
		Nama : c.FormValue("nama"),
	}
	users = append(users, newUser)

	return c.JSON(http.StatusCreated, newUser)
} 

// Func Get Users
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// Func Delete Users
func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param(("id")))
	index := 0

	// Find index in Slice by ID
	for idx := range users {
		if users[idx].ID == id {
			index = idx

			break
		}
	}
	deletedUser := users[index]
	users = append(users[:index], users[index+1:]...)

	return c.JSON(http.StatusOK, deletedUser)
}
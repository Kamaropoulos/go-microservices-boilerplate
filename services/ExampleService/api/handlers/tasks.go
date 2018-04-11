package APIHandlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// Dummy Handlers

// GetTasks endpoint
func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{"tasks": "all tasks"})
	}
}

// PutTask endpoint
func PutTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{"tasks": "put"})
	}
}

// DeleteTask endpoint
func DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{"tasks": "deleted"})
	}
}

package controllers

import (
	"learn-go-echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	result, err := models.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

package app

import (
	"github.com/labstack/echo/v4"
	"store/database/config"
	"store/database/models"
	"net/http"
)

func GetOrder(c echo.Context) error {
	id := c.Param("id")
	var order models.Order

	if result := database.DB.First(&order, id); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, order)
}

func SaveOrder(c echo.Context) error {
	order := new(models.Order)

	if err := c.Bind(order); err != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	database.DB.Create(order)

	return c.JSON(http.StatusOK, order)
}
package app

import (
	"github.com/labstack/echo/v4"
	"store/database/config"
	"store/database/models"
	"net/http"
)

func GetCart(c echo.Context) error {
	id := c.Param("id")
	var cart models.Cart

	if result := database.DB.First(&cart, id); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, cart)
}

func SaveCart(c echo.Context) error {
	cart := new(models.Cart)

	if err := c.Bind(cart); err != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	database.DB.Create(cart)

	return c.JSON(http.StatusOK, cart)
}
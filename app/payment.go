package app

import (
	"github.com/labstack/echo/v4"
	"store/database/config"
	"store/database/models"
	"net/http"
)

func GetPayment(c echo.Context) error {
	id := c.Param("id")
	var payment models.Payment

	if result := database.DB.First(&payment, id); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, payment)
}

func SavePayment(c echo.Context) error {
	payment := new(models.Payment)

	if err := c.Bind(payment); err != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	database.DB.Create(payment)

	return c.JSON(http.StatusOK, payment)
}
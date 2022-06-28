package app

import (
	"github.com/labstack/echo/v4"
	"store/database/config"
	"store/database/models"
	"net/http"
)

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product

	if result := database.DB.First(&product, id); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, product)
}

func AllProducts(c echo.Context) error {
	var productList []models.Product

	if result := database.DB.Find(&productList); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, productList)
}

func SaveProduct(c echo.Context) error {
	product := new(models.Product)

	if err := c.Bind(product); err != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	database.DB.Create(product)

	return c.JSON(http.StatusOK, product)
}
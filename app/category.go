package app

import (
	"github.com/labstack/echo/v4"
	"store/database/config"
	"store/database/models"
	"net/http"
)

func GetCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category

	if result := database.DB.First(&category, id); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, category)
}

func AllCategories(c echo.Context) error {
	var categoryList []models.Category

	if result := database.DB.Find(&categoryList); result.Error != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	return c.JSON(http.StatusOK, categoryList)
}

func SaveCategory(c echo.Context) error {
	category := new(models.Category)

	if err := c.Bind(category); err != nil {
		return c.String(http.StatusNotFound, "Database Error")
	}

	database.DB.Create(category)

	return c.JSON(http.StatusOK, category)
}
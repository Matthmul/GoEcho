package main

import (
	"net/http"
	
	"store/app"
	"github.com/labstack/echo/v4"
	"store/database/config"
)

func main() {
	database.Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/cart", 		app.SaveCart)
	e.GET("/cart/:id", 		app.GetCart)

	e.POST("/category", 	app.SaveCategory)
	e.GET("/category/:id", 	app.GetCategory)
	e.GET("/category", 		app.AllCategories)

	e.POST("/order", 		app.SaveOrder)
	e.GET("/order/:id", 	app.GetOrder)

	e.POST("/payment", 		app.SavePayment)
	e.GET("/payment/:id", 	app.GetPayment)

	e.POST("/product", 		app.SaveProduct)
	e.GET("/product/:id", 	app.GetProduct)
	e.GET("/product", 		app.AllProducts)

	e.Logger.Fatal(e.Start(":1323"))
}

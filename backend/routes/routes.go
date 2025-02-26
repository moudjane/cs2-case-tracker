package routes

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(e *echo.Echo, collection *mongo.Collection) {
	e.GET("/api/history", func(c echo.Context) error {
		return GetHistory(c, collection)
	})

	e.POST("/api/purchase", func(c echo.Context) error {
		return AddPurchasePrice(c, collection)
	})

	e.GET("/api/history/:caseName", func(c echo.Context) error {
		return GetSingleCaseHistory(c, collection)
	})
}

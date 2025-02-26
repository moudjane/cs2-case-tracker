package routes

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSingleCaseHistory(c echo.Context, collection *mongo.Collection) error {
	caseName := c.Param("caseName")

	cursor, err := collection.Find(context.TODO(), bson.M{"name": caseName})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erreur MongoDB"})
	}
	defer cursor.Close(context.TODO())

	var prices []bson.M
	if err := cursor.All(context.TODO(), &prices); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erreur de décodage"})
	}

	return c.JSON(http.StatusOK, prices)
}

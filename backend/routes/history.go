package routes

import (
	"context"
	"fmt"
	"net/http"

	"cs2-case-tracker-go/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetHistory(c echo.Context, collection *mongo.Collection) error {
	fmt.Println("🔍 Requête API reçue : /api/history")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("❌ Erreur MongoDB:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erreur MongoDB"})
	}
	defer cursor.Close(context.TODO())

	var prices []models.CasePrice
	if err := cursor.All(context.TODO(), &prices); err != nil {
		fmt.Println("❌ Erreur décodage MongoDB:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erreur de décodage"})
	}

	fmt.Println("✅ Données récupérées :", prices)
	return c.JSON(http.StatusOK, prices)
}

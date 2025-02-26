package routes

import (
	"context"
	"net/http"
	"time"

	"cs2-case-tracker-go/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddPurchasePrice(c echo.Context, collection *mongo.Collection) error {
	var casePrice models.CasePrice
	if err := c.Bind(&casePrice); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Format JSON invalide"})
	}

	if casePrice.Name == "" || casePrice.PurchasePrice == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Nom de la caisse et prix d'achat requis"})
	}

	casePrice.Date = time.Now()
	_, err := collection.InsertOne(context.TODO(), casePrice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Erreur MongoDB"})
	}

	return c.JSON(http.StatusCreated, casePrice)
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cs2-case-tracker-go/config"
	"cs2-case-tracker-go/routes"
	"cs2-case-tracker-go/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Aucun fichier .env trouvé, utilisation des variables système")
	}

	client, collection := config.ConnectDB()
	defer client.Disconnect(context.TODO())

	e := echo.New()
	routes.SetupRoutes(e, collection)

	services.StartCron(collection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	fmt.Println("🚀 Serveur Go lancé sur http://localhost:" + port)
	e.Logger.Fatal(e.Start(":" + port))
}

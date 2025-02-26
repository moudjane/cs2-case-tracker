package services

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartCron(collection *mongo.Collection) {
	c := cron.New()
	c.AddFunc("*/30 * * * *", func() { UpdatePrices(collection) })
	c.Start()
	fmt.Println("🕒 Cron démarré : Mise à jour des prix toutes les 10 minutes")
}

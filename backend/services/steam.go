package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"cs2-case-tracker-go/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func UpdatePrices(collection *mongo.Collection) {
	fmt.Println("⏳ Mise à jour automatique des prix...")

	currentTime := time.Now().Truncate(time.Minute)

	for _, caseName := range GetCaseNames() {
		caseNameEncoded := url.QueryEscape(caseName)
		steamURL := fmt.Sprintf("https://steamcommunity.com/market/priceoverview/?appid=730&currency=3&market_hash_name=%s", caseNameEncoded)

		client := &http.Client{}
		req, err := http.NewRequest("GET", steamURL, nil)
		if err != nil {
			fmt.Printf("❌ Erreur création requête Steam pour %s: %v\n", caseName, err)
			continue
		}

		req.Header.Set("User-Agent", "Mozilla/5.0")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("❌ Erreur requête Steam pour %s: %v\n", caseName, err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("❌ Erreur lecture réponse pour %s: %v\n", caseName, err)
			continue
		}

		fmt.Printf("🔍 Réponse brute Steam pour %s :\n%s\n", caseName, string(body))

		var data map[string]interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Printf("❌ Erreur parsing JSON pour %s: %v\n", caseName, err)
			continue
		}

		if success, ok := data["success"].(bool); !ok || !success {
			fmt.Printf("⚠️ Données Steam invalides pour %s\n", caseName)
			continue
		}

		medianPriceStr, ok := data["median_price"].(string)
		if !ok {
			fmt.Printf("⚠️ Aucune valeur `median_price` pour %s\n", caseName)
			continue
		}

		medianPrice := ParsePrice(medianPriceStr)

		_, err = collection.InsertOne(context.TODO(), models.CasePrice{
			Name:  caseName,
			Price: medianPrice,
			Date:  currentTime,
		})
		if err != nil {
			fmt.Printf("❌ Erreur MongoDB pour %s: %v\n", caseName, err)
		} else {
			fmt.Printf("✅ Prix enregistré pour %s : %.2f€\n", caseName, medianPrice)
		}
	}
}

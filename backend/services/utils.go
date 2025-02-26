package services

import (
	"fmt"
	"strconv"
	"strings"
)

func ParsePrice(priceStr string) float64 {
	priceStr = strings.Replace(strings.Replace(priceStr, "€", "", -1), ",", ".", -1)

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("❌ Erreur conversion prix :", err)
		return 0
	}
	return price
}

package models

import "time"

type CasePrice struct {
	Name          string    `json:"name" bson:"name"`
	Price         float64   `json:"price" bson:"price"`
	PurchasePrice *float64  `json:"purchasePrice,omitempty" bson:"purchasePrice,omitempty"`
	Date          time.Time `json:"date" bson:"date"`
}

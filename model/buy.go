package model

import "time"

type BuyWagerBody struct {
	BuyingPrice float64 `json:"buying_price" validate:"required,min=0"`
}

type Buy struct {
	ID          int64     `json:"id" db:"id"`
	WagerID     int64     `json:"wager_id" db:"wager_id"`
	BuyingPrice float64   `json:"buying_price" db:"buying_price"`
	BoughtAt    time.Time `json:"bought_at" db:"bought_at"`
}

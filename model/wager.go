package model

import "time"

type CreateWagerBody struct {
	TotalWagerValue   int     `json:"total_wager_value" validate:"required,min=0"`
	Odds              int     `json:"odds" validate:"required,min=0"`
	SellingPercentage int     `json:"selling_percentage" validate:"required,gte=0,lte=100"`
	SellingPrice      float64 `json:"selling_price" validate:"required,min=0"`
}

type Wager struct {
	ID                  int64     `json:"id" db:"id"`
	TotalWagerValue     int       `json:"total_wager_value" db:"total_wager_value"`
	Odds                int       `json:"odds" db:"odds"`
	SellingPercentage   int       `json:"selling_percentage" db:"selling_percentage"`
	SellingPrice        float64   `json:"selling_price" db:"selling_price"`
	CurrentSellingPrice float64   `json:"current_selling_price" db:"current_selling_price"`
	PercentageSold      *int      `json:"percentage_sold" db:"percentage_sold"`
	AmountSold          *float64  `json:"amount_sold" db:"amount_sold"`
	PlacedAt            time.Time `json:"placed_at" db:"placed_at"`
}

type ListWagerQueryString struct {
	PagingQuery
}

package repository

import (
	"database/sql"

	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
)

type IBuyRepository interface {
	InsertBuy(buy model.Buy) (*model.Buy, error)
}

type BuyRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewBuyRepository(db *sql.DB, logger *log.Logger) IBuyRepository {
	return &BuyRepository{
		db:     db,
		logger: logger,
	}
}

func (r *BuyRepository) InsertBuy(buy model.Buy) (*model.Buy, error) {
	res, err := r.db.Query("CALL buy_wager(?, ?, ?)", buy.WagerID, buy.BuyingPrice, buy.BoughtAt)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		err = res.Scan(&buy.ID, &buy.WagerID, &buy.BuyingPrice, &buy.BoughtAt)
		if err != nil {
			return nil, err
		}
	}

	return &buy, nil
}

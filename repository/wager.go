package repository

import (
	"database/sql"

	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
)

type IWagerRepository interface {
	InsertWager(wager model.Wager) (int64, error)
	ListWager(limit int, offset int) (*[]model.Wager, error)
}

type WagerRepository struct {
	logger *log.Logger
	db     *sql.DB
}

func NewWagerRepository(db *sql.DB, logger *log.Logger) IWagerRepository {
	return &WagerRepository{
		db:     db,
		logger: logger,
	}
}

func (r *WagerRepository) InsertWager(wager model.Wager) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO wagers (total_wager_value,odds,selling_percentage,selling_price,current_selling_price,percentage_sold,amount_sold,placed_at) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(wager.TotalWagerValue, wager.Odds, wager.SellingPercentage, wager.SellingPrice, wager.CurrentSellingPrice, wager.PercentageSold, wager.AmountSold, wager.PlacedAt)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (r *WagerRepository) ListWager(limit int, offset int) (*[]model.Wager, error) {
	var wagers []model.Wager
	res, err := r.db.Query("SELECT * FROM wagers LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		var wager model.Wager
		err := res.Scan(&wager.ID, &wager.TotalWagerValue, &wager.Odds, &wager.SellingPercentage, &wager.SellingPrice, &wager.CurrentSellingPrice, &wager.PercentageSold, &wager.AmountSold, &wager.PlacedAt)
		if err != nil {
			return nil, err
		}
		wagers = append(wagers, wager)
	}

	return &wagers, nil
}

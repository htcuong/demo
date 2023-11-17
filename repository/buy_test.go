package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestInsertBuyRepository(t *testing.T) {
	logger := log.NewLogger()

	buyRepo := BuyRepository{
		logger: logger,
	}

	t.Run("InsertBuy is failed with return error", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("err not expected while open a mock db, %v", err)
		}
		defer db.Close()
		buyRepo.db = db

		buy := model.Buy{
			WagerID:     1,
			BuyingPrice: 100,
		}
		mockDB.ExpectQuery("CALL buy_wager(?, ?, ?)").WithArgs(int64(1), float64(100), time.Now()).WillReturnError(pgx.ErrNoRows)

		result, err := buyRepo.InsertBuy(buy)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("InsertBuy is succeed", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("err not expected while open a mock db, %v", err)
		}
		defer db.Close()
		buyRepo.db = db

		boughtAt := time.Now()
		buy := model.Buy{
			WagerID:     1,
			BuyingPrice: 100,
			BoughtAt:    boughtAt,
		}

		rows := sqlmock.NewRows([]string{"id", "wager_id", "buying_price", "bought_at"}).AddRow(1, 1, 100, boughtAt)
		mockDB.ExpectQuery("CALL buy_wager(?, ?, ?)").WithArgs(int64(1), float64(100), boughtAt).WillReturnRows(rows)

		result, err := buyRepo.InsertBuy(buy)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

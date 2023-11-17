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

func TestInsertWagerRepository(t *testing.T) {
	logger := log.NewLogger()

	wagerRepo := WagerRepository{
		logger: logger,
	}

	t.Run("InsertWager is failed with return error", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("err not expected while open a mock db, %v", err)
		}
		defer db.Close()
		wagerRepo.db = db

		wager := model.Wager{
			TotalWagerValue: 1,
			SellingPrice:    100,
		}
		mockDB.ExpectQuery("INSERT INTO wagers (total_wager_value,odds,selling_percentage,selling_price,current_selling_price,percentage_sold,amount_sold,placed_at) VALUES (?,?,?,?,?,?,?,?)").WithArgs(int64(1), float64(100), time.Now()).WillReturnError(pgx.ErrNoRows)

		id, err := wagerRepo.InsertWager(wager)

		assert.NotNil(t, err)
		assert.Equal(t, int64(-1), id)
	})

	t.Run("InsertWager is succeed", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("err not expected while open a mock db, %v", err)
		}
		defer db.Close()
		wagerRepo.db = db

		placedAt := time.Now()
		wager := model.Wager{
			TotalWagerValue:     1,
			Odds:                1,
			SellingPrice:        100,
			SellingPercentage:   100,
			CurrentSellingPrice: 100,
			PlacedAt:            placedAt,
		}
		prep := mockDB.ExpectPrepare("INSERT INTO wagers (total_wager_value,odds,selling_percentage,selling_price,current_selling_price,percentage_sold,amount_sold,placed_at) VALUES (?,?,?,?,?,?,?,?)")
		prep.ExpectExec().WithArgs(int64(1), int64(1), 100, float64(100), float64(100), nil, nil, placedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := wagerRepo.InsertWager(wager)

		assert.Nil(t, err)
		assert.Equal(t, int64(1), id)
	})
}

func TestListWagerRepository(t *testing.T) {
	logger := log.NewLogger()

	wagerRepo := WagerRepository{
		logger: logger,
	}

	t.Run("ListWager is failed with return error", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("err not expected while open a mock db, %v", err)
		}
		defer db.Close()
		wagerRepo.db = db

		mockDB.ExpectQuery("SELECT * FROM wagers LIMIT ? OFFSET ?").WithArgs(1, 2).WillReturnError(pgx.ErrNoRows)

		wagers, err := wagerRepo.ListWager(1, 2)

		assert.NotNil(t, err)
		assert.Nil(t, wagers)
	})

	t.Run("ListWager is succeed", func(t *testing.T) {
		db, mockDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("err not expected while open a mock db, %v", err)
		}
		defer db.Close()
		wagerRepo.db = db
		rows := sqlmock.NewRows([]string{"id", "total_wager_value", "odds", "selling_percentage", "selling_price", "current_selling_price", "percentage_sold", "amount_sold", "placed_at"}).AddRow(1, 1, 100, 100, 100, 100, 100, 100, time.Now())

		mockDB.ExpectQuery("SELECT * FROM wagers LIMIT ? OFFSET ?").WithArgs(1, 2).WillReturnRows(rows)

		wagers, err := wagerRepo.ListWager(1, 2)

		assert.Nil(t, err)
		assert.NotNil(t, wagers)
	})
}

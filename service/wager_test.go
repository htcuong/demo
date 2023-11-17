package service

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	mocks "github.com/htcuong/demo/util/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWagerServide(t *testing.T) {

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("err not expected while open a mock db, %v", err)
	}
	defer db.Close()

	logger := log.NewLogger()

	wagerService := WagerService{
		db:     db,
		logger: logger,
	}

	gin.SetMode(gin.TestMode)

	t.Run("CreateWager is failed with return error", func(t *testing.T) {
		wagerRepo := &mocks.MockWagerRepository{}
		wagerService.wagerRepo = wagerRepo

		input := model.CreateWagerBody{
			TotalWagerValue: 100,
		}
		wagerRepo.On("InsertWager", mock.Anything).Return(int64(-1), errors.New("error"))
		result, err := wagerService.CreateWager(input)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("BuyWager is succeed", func(t *testing.T) {
		wagerRepo := &mocks.MockWagerRepository{}
		wagerService.wagerRepo = wagerRepo

		input := model.CreateWagerBody{
			TotalWagerValue: 100,
		}

		wagerRepo.On("InsertWager", mock.Anything).Return(int64(1), nil)
		result, err := wagerService.CreateWager(input)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func TestListWagerServide(t *testing.T) {

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("err not expected while open a mock db, %v", err)
	}
	defer db.Close()

	logger := log.NewLogger()

	wagerService := WagerService{
		db:     db,
		logger: logger,
	}

	gin.SetMode(gin.TestMode)

	t.Run("ListWager is failed with return error", func(t *testing.T) {
		wagerRepo := &mocks.MockWagerRepository{}
		wagerService.wagerRepo = wagerRepo

		wagerRepo.On("ListWager", mock.Anything).Return(nil, errors.New("error"))
		result, err := wagerService.ListWager(1, 2)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("ListWager is succeed", func(t *testing.T) {
		wagerRepo := &mocks.MockWagerRepository{}
		wagerService.wagerRepo = wagerRepo

		wagers := []model.Wager{model.Wager{ID: 1}}
		wagerRepo.On("ListWager", mock.Anything).Return(&wagers, nil)
		result, err := wagerService.ListWager(1, 2)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

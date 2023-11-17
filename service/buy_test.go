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

func TestBuyWagerServide(t *testing.T) {

	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("err not expected while open a mock db, %v", err)
	}
	defer db.Close()

	logger := log.NewLogger()

	buyService := BuyService{
		db:     db,
		logger: logger,
	}

	gin.SetMode(gin.TestMode)

	t.Run("BuyWager is failed with return error", func(t *testing.T) {
		buyRepo := &mocks.MockBuyRepository{}
		buyService.buyRepo = buyRepo

		wagerID := int64(1)
		input := model.BuyWagerBody{
			BuyingPrice: 100,
		}
		buyRepo.On("InsertBuy", mock.Anything).Return(nil, errors.New("error"))
		buy, err := buyService.CreateBuy(wagerID, input)
		assert.NotNil(t, err)
		assert.Nil(t, buy)
	})

	t.Run("BuyWager is succeed", func(t *testing.T) {
		buyRepo := &mocks.MockBuyRepository{}
		buyService.buyRepo = buyRepo

		wagerID := int64(1)
		input := model.BuyWagerBody{
			BuyingPrice: 100,
		}

		buy := model.Buy{
			ID: 1,
		}
		buyRepo.On("InsertBuy", mock.Anything).Return(&buy, nil)
		result, err := buyService.CreateBuy(wagerID, input)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

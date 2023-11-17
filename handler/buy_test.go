package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	mocks "github.com/htcuong/demo/util/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBuyWagerHandler(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("err not expected while open a mock db, %v", err)
	}
	defer db.Close()

	logger := log.NewLogger()

	buyHandler := BuyHandler{
		db:       db,
		logger:   logger,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	gin.SetMode(gin.TestMode)

	t.Run("BuyWager is failed with return error on header when wager_id is not int", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		e.POST("/buy/:wager_id", buyHandler.BuyWager)
		c.Request, err = http.NewRequest("POST", "/buy/test", nil)
		assert.NoError(t, err)
		e.ServeHTTP(w, c.Request)
		err := w.Header().Get("error")
		assert.NotNil(t, err)
	})

	t.Run("BuyWager is failed with return error on header when request body is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		e.POST("/buy/:wager_id", buyHandler.BuyWager)
		buyBody := model.BuyWagerBody{}
		buyBodyJSON, _ := json.Marshal(buyBody)

		c.Request, err = http.NewRequest("POST", "/buy/1", bytes.NewBuffer(buyBodyJSON))
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		err := w.Header().Get("error")
		assert.NotNil(t, err)
	})

	t.Run("BuyWager is succeed", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		buySvc := &mocks.MockBuyService{}
		buyHandler.buyService = buySvc
		e.POST("/buy/:wager_id", buyHandler.BuyWager)

		jsonParam := `{"buying_price":100}`

		buy := model.Buy{
			ID:          1,
			WagerID:     1,
			BuyingPrice: 100,
			BoughtAt:    time.Now(),
		}
		buySvc.On("CreateBuy", mock.Anything).Return(&buy, nil)

		c.Request, err = http.NewRequest("POST", "/buy/1", strings.NewReader(string(jsonParam)))
		c.Request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		errHeader := w.Header().Get("error")
		assert.Equal(t, "", errHeader)

		var response model.Buy
		err = json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

	t.Run("BuyWager is failed when service.CreateBuy return err", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)

		buySvc := &mocks.MockBuyService{}
		buyHandler.buyService = buySvc

		e.POST("/buy/:wager_id", buyHandler.BuyWager)

		jsonParam := `{"buying_price":100}`

		buySvc.On("CreateBuy", mock.Anything).Return(nil, errors.New("error"))

		c.Request, err = http.NewRequest("POST", "/buy/1", strings.NewReader(string(jsonParam)))
		c.Request.Header.Set("Content-Type", "application/json")
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		errHeader := w.Header().Get("error")
		assert.NotEqual(t, "", errHeader)

		var response model.Buy
		err = json.Unmarshal([]byte(w.Body.String()), &response)
		assert.NotNil(t, err)
	})
}

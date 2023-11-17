package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	mocks "github.com/htcuong/demo/util/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateWagerHandler(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("err not expected while open a mock db, %v", err)
	}
	defer db.Close()

	logger := log.NewLogger()

	wagerHandler := WagerHandler{
		db:       db,
		logger:   logger,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	gin.SetMode(gin.TestMode)

	t.Run("CreateWager is failed with return error on header when request body is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		e.POST("/wagers", wagerHandler.CreateWager)
		wagerBody := model.CreateWagerBody{}
		wagerBodyJSON, _ := json.Marshal(wagerBody)

		c.Request, err = http.NewRequest("POST", "/wagers", bytes.NewBuffer(wagerBodyJSON))
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		err := w.Header().Get("error")
		assert.NotNil(t, err)
	})

	t.Run("BuyWager is succeed", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		wagerSvc := &mocks.MockWagerService{}
		wagerHandler.wagerService = wagerSvc
		e.POST("/wagers", wagerHandler.CreateWager)

		jsonParam := `{
			"total_wager_value": 100,
			"odds": 1,
			"selling_percentage": 100,
			"selling_price": 102.101
		}`

		wager := model.Wager{
			ID: 1,
		}
		wagerSvc.On("CreateWager", mock.Anything).Return(&wager, nil)

		c.Request, err = http.NewRequest("POST", "/wagers", strings.NewReader(string(jsonParam)))
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

	t.Run("CreateWager is failed when service.Create return err", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		wagerSvc := &mocks.MockWagerService{}
		wagerHandler.wagerService = wagerSvc
		e.POST("/wagers", wagerHandler.CreateWager)

		jsonParam := `{
			"total_wager_value": 100,
			"odds": 1,
			"selling_percentage": 100,
			"selling_price": 102.101
		}`

		wagerSvc.On("CreateWager", mock.Anything).Return(nil, errors.New("error"))

		c.Request, err = http.NewRequest("POST", "/wagers", strings.NewReader(string(jsonParam)))
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

func TestListWagerHandler(t *testing.T) {
	db, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("err not expected while open a mock db, %v", err)
	}
	defer db.Close()

	logger := log.NewLogger()

	wagerHandler := WagerHandler{
		db:       db,
		logger:   logger,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}

	gin.SetMode(gin.TestMode)

	t.Run("ListWager is failed with return error on header when request query is empty", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		e.GET("/wagers", wagerHandler.GetListWager)

		c.Request, err = http.NewRequest("GET", "/wagers", nil)
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		err := w.Header().Get("error")
		assert.NotEqual(t, "", err)
	})

	t.Run("ListWager is failed with return error on header when request query is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)
		e.GET("/wagers", wagerHandler.GetListWager)
		c.Request, err = http.NewRequest("GET", "/wagers?page=1&limit=ab", nil)
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		errHeader := w.Header().Get("error")
		assert.NotEqual(t, "", errHeader)
	})

	t.Run("ListWager is succeed", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)

		wagerSvc := &mocks.MockWagerService{}
		wagerHandler.wagerService = wagerSvc

		wagers := []model.Wager{model.Wager{ID: 1}}
		wagerSvc.On("ListWager", mock.Anything).Return(&wagers, nil)

		e.GET("/wagers", wagerHandler.GetListWager)
		c.Request, err = http.NewRequest("GET", "/wagers?page=1&limit=2", nil)
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		errHeader := w.Header().Get("error")
		assert.Equal(t, "", errHeader)

		var response []model.Wager
		err = json.Unmarshal([]byte(w.Body.String()), &response)
		assert.Nil(t, err)
		assert.NotNil(t, response)
	})

	t.Run("ListWager is failed when service.ListWager return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, e := gin.CreateTestContext(w)

		wagerSvc := &mocks.MockWagerService{}
		wagerHandler.wagerService = wagerSvc

		wagerSvc.On("ListWager", mock.Anything).Return(nil, errors.New("some error"))

		e.GET("/wagers", wagerHandler.GetListWager)
		c.Request, err = http.NewRequest("GET", "/wagers?page=1&limit=2", nil)
		assert.NoError(t, err)

		e.ServeHTTP(w, c.Request)
		errHeader := w.Header().Get("error")
		assert.NotEqual(t, "", errHeader)
	})
}

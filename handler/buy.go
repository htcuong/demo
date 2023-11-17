package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	"github.com/htcuong/demo/service"
)

type IBuyHandler interface {
	BuyWager(c *gin.Context)
}

type BuyHandler struct {
	logger       *log.Logger
	db           *sql.DB
	wagerService service.IWagerService
	buyService   service.IBuyService
	validate     *validator.Validate
}

func NewBuyHandler(db *sql.DB, logger *log.Logger) IBuyHandler {
	return &BuyHandler{
		logger:       logger,
		db:           db,
		wagerService: service.NewWagerService(db, logger),
		buyService:   service.NewBuyService(db, logger),
		validate:     validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (h *BuyHandler) BuyWager(c *gin.Context) {
	wagerIDParam := c.Param("wager_id")
	wagerID, err := strconv.ParseInt(wagerIDParam, 10, 64)
	if err != nil {
		c.Header("error", err.Error())
		return
	}

	var body model.BuyWagerBody
	c.ShouldBind(&body)
	err = h.validate.Struct(body)
	if err != nil {
		c.Header("error", err.Error())
		return
	}
	buy, err := h.buyService.CreateBuy(wagerID, body)
	if err != nil {
		c.Header("error", err.Error())
		return
	}
	c.JSON(http.StatusCreated, buy)
}

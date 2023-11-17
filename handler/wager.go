package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	"github.com/htcuong/demo/service"
	"github.com/htcuong/demo/util"
)

type IWagerHandler interface {
	CreateWager(c *gin.Context)
	GetListWager(c *gin.Context)
}

type WagerHandler struct {
	logger       *log.Logger
	db           *sql.DB
	wagerService service.IWagerService
	validate     *validator.Validate
}

func NewWagerHandler(db *sql.DB, logger *log.Logger) IWagerHandler {
	wagerService := service.NewWagerService(db, logger)
	return &WagerHandler{
		logger:       logger,
		db:           db,
		wagerService: wagerService,
		validate:     validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (h *WagerHandler) CreateWager(c *gin.Context) {
	var body model.CreateWagerBody
	c.ShouldBind(&body)
	err := h.validate.Struct(body)
	if err != nil {
		c.Header("error", err.Error())
		return
	}

	sellingPrice := body.TotalWagerValue * (body.SellingPercentage / 100)
	if body.SellingPrice < float64(sellingPrice) {
		c.Header("error", "selling_price is not correct")
		return
	}

	wager, err := h.wagerService.CreateWager(body)
	if err != nil {
		c.Header("error", err.Error())
		return
	}
	c.JSON(http.StatusCreated, wager)
}

func (h *WagerHandler) GetListWager(c *gin.Context) {
	var query model.ListWagerQueryString
	c.ShouldBind(&query)

	err := h.validate.Struct(query)
	if err != nil {
		c.Header("error", err.Error())
		return
	}

	offset := util.OffsetFromPage(query.Page, query.Limit)
	wagers, err := h.wagerService.ListWager(offset, query.Limit)
	if err != nil {
		c.Header("error", err.Error())
		return
	}
	c.JSON(http.StatusCreated, wagers)
}

package service

import (
	"database/sql"
	"time"

	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	"github.com/htcuong/demo/repository"
)

type IBuyService interface {
	CreateBuy(wagerID int64, input model.BuyWagerBody) (*model.Buy, error)
}

type BuyService struct {
	logger  *log.Logger
	db      *sql.DB
	buyRepo repository.IBuyRepository
}

func NewBuyService(db *sql.DB, logger *log.Logger) IBuyService {
	BuyRepo := repository.NewBuyRepository(db, logger)
	return &BuyService{
		logger:  logger,
		db:      db,
		buyRepo: BuyRepo,
	}
}

func (s *BuyService) CreateBuy(wagerID int64, input model.BuyWagerBody) (*model.Buy, error) {
	buy := model.Buy{
		WagerID:     wagerID,
		BuyingPrice: input.BuyingPrice,
		BoughtAt:    time.Now(),
	}

	result, err := s.buyRepo.InsertBuy(buy)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return result, nil
}

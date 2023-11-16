package service

import (
	"database/sql"
	"math"
	"time"

	"github.com/htcuong/demo/model"
	"github.com/htcuong/demo/pkg/log"
	"github.com/htcuong/demo/repository"
)

type IWagerService interface {
	CreateWager(input model.CreateWagerBody) (*model.Wager, error)
	ListWager(offset int, limit int) (*[]model.Wager, error)
}

type WagerService struct {
	logger    *log.Logger
	db        *sql.DB
	wagerRepo repository.IWagerRepository
}

func NewWagerService(db *sql.DB, logger *log.Logger) IWagerService {
	wagerRepo := repository.NewWagerRepository(db, logger)
	return &WagerService{
		logger:    logger,
		db:        db,
		wagerRepo: wagerRepo,
	}
}

func (s *WagerService) CreateWager(input model.CreateWagerBody) (*model.Wager, error) {
	s.logger.Infof("%+v", input)
	sellingPrice := math.Round(input.SellingPrice*100) / 100
	wager := model.Wager{
		TotalWagerValue:     input.TotalWagerValue,
		Odds:                input.Odds,
		SellingPercentage:   input.SellingPercentage,
		SellingPrice:        sellingPrice,
		CurrentSellingPrice: sellingPrice,
		PlacedAt:            time.Now(),
	}
	id, err := s.wagerRepo.InsertWager(wager)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	s.logger.Infof("%+v | %+v", id, err)
	wager.ID = id
	return &wager, nil
}

func (s *WagerService) ListWager(offset int, limit int) (*[]model.Wager, error) {
	wagers, err := s.wagerRepo.ListWager(limit, offset)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return wagers, nil
}

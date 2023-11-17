package mock

import (
	"github.com/htcuong/demo/model"
	"github.com/stretchr/testify/mock"
)

type MockBuyService struct {
	mock.Mock
}

func (p *MockBuyService) CreateBuy(wagerID int64, input model.BuyWagerBody) (*model.Buy, error) {
	args := p.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Buy), args.Error(1)
}

type MockBuyRepository struct {
	mock.Mock
}

func (p *MockBuyRepository) InsertBuy(buy model.Buy) (*model.Buy, error) {
	args := p.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Buy), args.Error(1)
}

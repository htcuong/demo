package mock

import (
	"github.com/htcuong/demo/model"
	"github.com/stretchr/testify/mock"
)

type MockWagerService struct {
	mock.Mock
}

func (p *MockWagerService) CreateWager(input model.CreateWagerBody) (*model.Wager, error) {
	args := p.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Wager), args.Error(1)
}

func (p *MockWagerService) ListWager(offset int, limit int) (*[]model.Wager, error) {
	args := p.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]model.Wager), args.Error(1)
}

type MockWagerRepository struct {
	mock.Mock
}

func (p *MockWagerRepository) InsertWager(wager model.Wager) (int64, error) {
	args := p.Called()

	if args.Get(0) == nil {
		return -1, args.Error(1)
	}
	return args.Get(0).(int64), args.Error(1)
}

func (p *MockWagerRepository) ListWager(limit int, offset int) (*[]model.Wager, error) {
	args := p.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]model.Wager), args.Error(1)
}

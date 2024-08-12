package transactionsvc

import "github.com/amnestia/xyz-multifinance/mocks"

var (
	mockRepo     *mocks.TransactionRepository
	mockAuthRepo *mocks.AuthRepository
)

func initMock() {
	mockRepo = new(mocks.TransactionRepository)
	mockAuthRepo = new(mocks.AuthRepository)
}

func getMock() *Service {
	initMock()
	return &Service{
		Repo:     mockRepo,
		AuthRepo: mockAuthRepo,
	}
}

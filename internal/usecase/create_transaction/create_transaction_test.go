package CreateTransaction

import (
	"testing"

	"github.com.br/vyctor/fc-microsservicos/internal/entity"
	"github.com.br/vyctor/fc-microsservicos/internal/event"
	"github.com.br/vyctor/fc-microsservicos/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUsecase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "j@j")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000.0)

	client2, _ := entity.NewClient("John Doe", "j@j")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000.0)

	mockAccount := &AccountGatewayMock{}

	mockAccount.On("FindById", account1.ID).Return(account1, nil)
	mockAccount.On("FindById", account2.ID).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDto{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100.0,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()

	usecase := NewCreateTransactionUsecase(mockTransaction, mockAccount, dispatcher, event)
	output, err := usecase.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
}

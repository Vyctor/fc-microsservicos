package CreateTransaction

import (
	"github.com.br/vyctor/fc-microsservicos/internal/entity"
	"github.com.br/vyctor/fc-microsservicos/internal/gateway"
	"github.com.br/vyctor/fc-microsservicos/pkg/events"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDto struct {
	ID string
}

type CreateTransactionUsecase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUsecase(
	transactionGateway gateway.TransactionGateway,
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUsecase {
	return &CreateTransactionUsecase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (usecase *CreateTransactionUsecase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	accountFrom, err := usecase.AccountGateway.FindByID(input.AccountIDFrom)

	if err != nil {
		return nil, err
	}

	accountTo, err := usecase.AccountGateway.FindByID(input.AccountIDTo)

	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

	if err != nil {
		return nil, err
	}

	err = usecase.TransactionGateway.Create(transaction)

	if err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDto{
		ID: transaction.ID,
	}

	usecase.TransactionCreated.SetPayload(output)
	usecase.EventDispatcher.Dispatch(usecase.TransactionCreated)

	return &CreateTransactionOutputDto{ID: transaction.ID}, nil
}

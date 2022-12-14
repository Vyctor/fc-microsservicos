package CreateAccount

import (
	"github.com.br/vyctor/fc-microsservicos/internal/entity"
	"github.com.br/vyctor/fc-microsservicos/internal/gateway"
)

type CreateAccountInputDto struct {
	ClientID string
}

type CreateAccountOutput struct {
	ID string
}

type CreateAccountUsecase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUsecase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUsecase {
	return &CreateAccountUsecase{AccountGateway: accountGateway, ClientGateway: clientGateway}
}

func (u *CreateAccountUsecase) Execute(input CreateAccountInputDto) (*CreateAccountOutput, error) {
	client, err := u.ClientGateway.Get(input.ClientID)

	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)

	err = u.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutput{ID: account.ID}, nil
}

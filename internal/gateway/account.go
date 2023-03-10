package gateway

import "github.com.br/vyctor/fc-microsservicos/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}

package gateway

import "github.com.br/vyctor/fc-microsservicos/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

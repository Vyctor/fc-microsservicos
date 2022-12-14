package gateway

import "github.com.br/vyctor/fc-microsservicos/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

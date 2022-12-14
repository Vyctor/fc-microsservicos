package CreateClient

import (
	"testing"

	"github.com.br/vyctor/fc-microsservicos/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreateClientUsecase_Execute(t *testing.T) {
	mockClientGateway := &ClientGatewayMock{}

	mockClientGateway.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateClientUsecase(mockClientGateway)

	output, err := usecase.Execute(CreateClientInputDTO{
		Name:  "John Doe",
		Email: "j@j.com"})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j.com", output.Email)
	mockClientGateway.AssertExpectations(t)
}

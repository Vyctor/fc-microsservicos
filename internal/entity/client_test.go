package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")

	if err != nil {
		t.Error("Error creating new client")
	}

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")

	err := client.Update("John Doe Jr", "jjr@j.com")

	assert.Nil(t, err)
	assert.Equal(t, "John Doe Jr", client.Name)
	assert.Equal(t, "jjr@j.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")

	err := client.Update("", "jjr@j.com")

	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@#j.com")

	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestAddAnotherClientAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@#j.com")
	anotherClient, _ := NewClient("John Doe Jr", "jjr@#j.com")

	account := NewAccount(client)

	err := anotherClient.AddAccount(account)

	assert.Error(t, err, "account already belongs to another client")
}

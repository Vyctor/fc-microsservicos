package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@#j.com")
	client2, _ := NewClient("John Doe Jr", "jjr@#j.com")

	account1 := NewAccount(client1)
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account1.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@#j.com")
	client2, _ := NewClient("John Doe Jr", "jjr@#j.com")

	account1 := NewAccount(client1)
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 5000)

	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionInvalidValue(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@#j.com")
	client2, _ := NewClient("John Doe Jr", "jjr@#j.com")

	account1 := NewAccount(client1)
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, -5000)

	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionInvalidAccount(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@#j.com")

	account1 := NewAccount(client1)

	account1.Credit(1000)

	transaction, err := NewTransaction(account1, nil, -5000)

	assert.NotNil(t, err)
	assert.Error(t, err, "accounts are required")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account1.Balance)
}

package database

import (
	"database/sql"
	"testing"

	"github.com.br/vyctor/fc-microsservicos/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTEstSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTEstSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")

	client, _ := entity.NewClient("Jon Snow", "j@j.com")
	s.client = client
	client2, _ := entity.NewClient("Arya Stark", "arya@arya.com")
	s.client2 = client2

	s.accountFrom = entity.NewAccount(s.client)
	s.accountFrom.Balance = 1000
	s.accountTo = entity.NewAccount(s.client2)
	s.accountTo.Balance = 1000

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTEstSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("Drop table clients")
	s.db.Exec("Drop table accounts")
	s.db.Exec("Drop table transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTEstSuite))
}

func (s *TransactionDBTEstSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100.0)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}

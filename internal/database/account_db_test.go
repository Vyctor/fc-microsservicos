package database

import (
	"database/sql"
	"testing"

	"github.com.br/vyctor/fc-microsservicos/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")

	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("Jon Snow", "j@j.com")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("Drop table clients")
	s.db.Exec("Drop table accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
	s.db.Exec("INSERT INTO clients(id, name, email, created_at) VALUES(?, ?, ?, ?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)

	account := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
	account, err = s.accountDB.FindByID(account.ID)
	s.Nil(err)

	s.Equal(account.ID, account.ID)
	s.Equal(account.Client.ID, account.Client.ID)
	s.Equal(account.Client.Name, account.Client.Name)
	s.Equal(account.Client.Email, account.Client.Email)
	s.Equal(account.Client.CreatedAt, account.Client.CreatedAt)
	s.Equal(account.Balance, account.Balance)
	s.Equal(account.CreatedAt, account.CreatedAt)
}

package main

import (
	"database/sql"
	"fmt"

	"github.com.br/vyctor/fc-microsservicos/internal/database"
	"github.com.br/vyctor/fc-microsservicos/internal/event"
	CreateAccount "github.com.br/vyctor/fc-microsservicos/internal/usecase/create_account"
	CreateClient "github.com.br/vyctor/fc-microsservicos/internal/usecase/create_client"
	CreateTransaction "github.com.br/vyctor/fc-microsservicos/internal/usecase/create_transaction"
	"github.com.br/vyctor/fc-microsservicos/internal/web"
	"github.com.br/vyctor/fc-microsservicos/internal/web/webserver"
	"github.com.br/vyctor/fc-microsservicos/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root,", "root", "localhost", "3306", "wallet"))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	// eventDispatcher.Register("TransactionCreated", handler);

	transactionCreatedEvent := event.NewTransactionCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUsecase := CreateClient.NewCreateClientUsecase(clientDb)
	createAccountUsecase := CreateAccount.NewCreateAccountUsecase(accountDb, clientDb)
	createTransactionUsecase := CreateTransaction.NewCreateTransactionUsecase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUsecase)
	accountHandler := web.NewWebAccountHandler(*createAccountUsecase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}

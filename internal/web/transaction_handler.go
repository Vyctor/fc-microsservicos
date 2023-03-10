package web

import (
	"encoding/json"
	"net/http"

	CreateTransaction "github.com.br/vyctor/fc-microsservicos/internal/usecase/create_transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUsecase CreateTransaction.CreateTransactionUsecase
}

func NewWebTransactionHandler(createTransactionUsecase CreateTransaction.CreateTransactionUsecase) *WebTransactionHandler {
	return &WebTransactionHandler{
		CreateTransactionUsecase: createTransactionUsecase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto CreateTransaction.CreateTransactionInputDto

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateTransactionUsecase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

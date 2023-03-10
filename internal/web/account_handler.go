package web

import (
	"encoding/json"
	"net/http"

	CreateAccount "github.com.br/vyctor/fc-microsservicos/internal/usecase/create_account"
)

type WebAccountHandler struct {
	CreateClientUsecase CreateAccount.CreateAccountUsecase
}

func NewWebAccountHandler(createAccountUsecase CreateAccount.CreateAccountUsecase) *WebAccountHandler {
	return &WebAccountHandler{
		CreateClientUsecase: createAccountUsecase,
	}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto CreateAccount.CreateAccountInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.CreateClientUsecase.Execute(dto)
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

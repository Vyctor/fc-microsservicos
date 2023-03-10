package web

import (
	"encoding/json"
	"net/http"

	CreateClient "github.com.br/vyctor/fc-microsservicos/internal/usecase/create_client"
)

type WebClientHandler struct {
	CreateClientUsecase CreateClient.CreateClientUseCase
}

func NewWebClientHandler(createClientUsecase CreateClient.CreateClientUseCase) WebClientHandler {
	return WebClientHandler{
		CreateClientUsecase: createClientUsecase,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto CreateClient.CreateClientInputDTO

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

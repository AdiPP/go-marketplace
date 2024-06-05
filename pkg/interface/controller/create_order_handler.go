package controller

import (
	"encoding/json"
	"fmt"
	"github.com/AdiPP/go-marketplace/pkg/usecase"
	"net/http"
)

type CreateOrderHandler struct {
	createOrderUseCase *usecase.CreateOrderUseCase
}

func NewCreateOrderHandler(createOrderUseCase *usecase.CreateOrderUseCase) *CreateOrderHandler {
	return &CreateOrderHandler{createOrderUseCase: createOrderUseCase}
}

func (h *CreateOrderHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var requestData usecase.CreateOrderDto

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	order, err := h.createOrderUseCase.Execute(r.Context(), requestData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("order created: %s", order.Id())))
	return
}

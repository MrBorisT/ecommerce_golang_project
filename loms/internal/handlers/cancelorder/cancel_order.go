package cancelorder

import (
	"context"
	"errors"
	"log"
	"route256/loms/internal/domain"
)

type Request struct {
	OrderID int64 `json:"orderID"`
}

func (r Request) Validate() error {
	switch {
	case r.OrderID == 0:
		return ErrEmptyOrder
	}
	return nil
}

var (
	ErrEmptyOrder = errors.New("empty order")
)

type Response struct {
}

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("cancel order: %+v", request)

	var response Response
	err := h.businessLogic.CancelOrder(ctx, request.OrderID)
	if err != nil {
		return response, err
	}

	return response, nil
}

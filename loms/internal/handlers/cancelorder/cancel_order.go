package cancelorder

import (
	"context"
	"errors"
	"log"
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

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("cancel order: %+v", request)

	return Response{}, nil
}

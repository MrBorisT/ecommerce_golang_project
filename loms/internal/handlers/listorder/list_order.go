package listorder

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

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	Status string `json:"status"`
	User   int64  `json:"user"`
	Items  []Item `json:"items"`
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("list order: %+v", request)

	return Response{
		Status: "new",
		User:   1,
		Items: []Item{
			{
				SKU:   1,
				Count: 1,
			},
		},
	}, nil
}

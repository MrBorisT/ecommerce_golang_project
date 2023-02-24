package stockshandler

import (
	"context"
	"errors"
	"log"
)

type Request struct {
	SKU uint32 `json:"sku"`
}

func (r Request) Validate() error {
	switch {
	case r.SKU == 0:
		return ErrEmptySKU
	}
	return nil
}

var (
	ErrEmptySKU = errors.New("empty sku")
)

type Item struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type Response struct {
	Stocks []Item `json:"stocks"`
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("stocks: %+v", request)
	return Response{
		Stocks: []Item{
			{
				WarehouseID: 123,
				Count:       5,
			},
		},
	}, nil
}

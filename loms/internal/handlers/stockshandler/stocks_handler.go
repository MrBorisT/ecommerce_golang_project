package stockshandler

import (
	"context"
	"log"
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
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

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}
func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("stocks: %+v", request)

	var response Response
	err := h.businessLogic.Stocks(ctx, request.SKU)
	if err != nil {
		return response, err
	}

	return response, nil

}

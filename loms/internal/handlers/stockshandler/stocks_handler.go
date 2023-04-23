package stockshandler

import (
	"context"
	"route256/loms/internal/domain"
	"route256/loms/internal/model"

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

type Response struct {
	Stocks []model.Stock `json:"stocks"`
}

type Handler struct {
	businessLogic domain.Service
}

func New(businessLogic domain.Service) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}
func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	var response Response
	stocks, err := h.businessLogic.Stocks(ctx, request.SKU)
	if err != nil {
		return response, err
	}

	response.Stocks = stocks

	return response, nil

}

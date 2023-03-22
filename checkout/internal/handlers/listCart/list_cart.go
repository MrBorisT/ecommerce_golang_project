package listcart

import (
	"context"
	"log"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/model"

	"github.com/pkg/errors"
)

type Handler struct {
	businessLogic domain.CheckoutService
}

func New(businessLogic domain.CheckoutService) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

type Request struct {
	User int64 `json:"user"`
}

var (
	ErrEmptyUser = errors.New("empty user")
)

func (r Request) Validate() error {
	switch {
	case r.User == 0:
		return ErrEmptyUser
	}

	return nil
}

type Response struct {
	Items      []model.Item `json:"items"`
	TotalPrice uint32       `json:"total_price"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("listCart: %+v", req)

	var response Response

	items, totalPrice, err := h.businessLogic.ListCart(ctx, req.User)
	if err != nil {
		return response, err
	}

	response.Items = items
	response.TotalPrice = totalPrice

	return response, nil
}

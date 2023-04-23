package purchase

import (
	"context"
	"route256/checkout/internal/domain"

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
	OrderID string `json:"orderID"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	var response Response

	err := h.businessLogic.Purchase(ctx, req.User)
	if err != nil {
		return response, err
	}

	return response, nil
}

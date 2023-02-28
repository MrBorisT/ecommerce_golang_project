package purchase

import (
	"context"
	"errors"
	"log"
	"route256/checkout/internal/domain"
)

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
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
	log.Printf("purchase: %+v", req)

	var response Response

	err := h.businessLogic.Purchase(ctx, req.User)
	if err != nil {
		return response, err
	}

	return response, nil
}

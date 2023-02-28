package createorder

import (
	"context"
	"log"
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
)

type Request struct {
	User  int64  `json:"user"`
	Items []Item `json:"items"`
}

func (r Request) Validate() error {

	switch {
	case r.User == 0:
		return ErrEmptyUser
	case len(r.Items) == 0:
		return ErrEmptyItems
	}
	return nil
}

var (
	ErrEmptyUser  = errors.New("empty user")
	ErrEmptyItems = errors.New("empty items")
)

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	OrderID int64 `json:"orderID"`
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
	log.Printf("create order: %+v", request)

	var response Response
	err := h.businessLogic.CreateOrder(ctx, request.User, request.Items)
	if err != nil {
		return response, err
	}

	return response, nil
}

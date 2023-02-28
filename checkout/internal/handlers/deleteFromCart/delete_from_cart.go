package deletefromcart

import (
	"context"
	"log"
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
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
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

var (
	ErrEmptyUser  = errors.New("empty user")
	ErrEmptySKU   = errors.New("empty sku")
	ErrEmptyCount = errors.New("empty count")
)

func (r Request) Validate() error {
	switch {
	case r.User == 0:
		return ErrEmptyUser
	case r.Sku == 0:
		return ErrEmptySKU
	case r.Count == 0:
		return ErrEmptyCount
	}

	return nil
}

type Response struct {
	Test string `json:"test"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("deleteFromCart: %+v", req)

	var response Response

	err := h.businessLogic.DeleteFromCart(ctx, req.User, req.Sku, req.Count)
	if err != nil {
		return response, err
	}

	return response, nil
}

package listorder

import (
	"context"
	"log"
	"route256/loms/internal/domain"

	"github.com/pkg/errors"
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

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}
func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("list order: %+v", request)

	var response Response
	err := h.businessLogic.ListOrder(ctx, request.OrderID)
	if err != nil {
		return response, err
	}

	return response, nil

}

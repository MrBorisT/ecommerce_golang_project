package listorder

import (
	"context"
	"route256/loms/internal/domain"
	"route256/loms/internal/model"

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

type Response struct {
	Status string       `json:"status"`
	User   int64        `json:"user"`
	Items  []model.Item `json:"items"`
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
	status, user, items, err := h.businessLogic.ListOrder(ctx, request.OrderID)

	if err != nil {
		return response, err
	}

	response.Status = status
	response.User = user
	response.Items = items

	return response, nil

}

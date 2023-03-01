package createorder

import (
	"context"
	"log"
	"route256/loms/internal/domain"
	"route256/loms/internal/model"

	"github.com/pkg/errors"
)

type Request struct {
	User  int64        `json:"user"`
	Items []model.Item `json:"items"`
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

type Response struct {
	OrderID int64 `json:"orderID"`
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
	log.Printf("create order: %+v", request)

	var response Response
	orderID, err := h.businessLogic.CreateOrder(ctx, request.User, request.Items)
	if err != nil {
		return response, err
	}

	response.OrderID = orderID

	return response, nil
}

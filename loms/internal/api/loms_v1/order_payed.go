package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implementation) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*desc.OrderPayedResponse, error) {
	err := i.lomsService.OrderPayed(ctx, req.GetOrderID())
	if err != nil {
		return nil, err
	}

	return &desc.OrderPayedResponse{}, nil
}

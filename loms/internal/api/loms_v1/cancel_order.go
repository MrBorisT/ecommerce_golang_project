package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implementation) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*desc.CancelOrderResponse, error) {
	err := i.lomsService.CancelOrder(ctx, req.GetOrderID())
	if err != nil {
		return nil, err
	}

	return &desc.CancelOrderResponse{}, nil
}

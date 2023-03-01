package loms

import (
	"context"
	lomsServiceAPI "route256/checkout/pkg/loms_v1"

	"github.com/pkg/errors"
)

func (c *client) CancelOrder(ctx context.Context, orderID int64) error {
	// response is empty
	_, err := c.lomsClient.CancelOrder(ctx, &lomsServiceAPI.CancelOrderRequest{
		OrderID: orderID,
	})

	if err != nil {
		return errors.WithMessage(err, "rpc cancel order")
	}

	return nil
}

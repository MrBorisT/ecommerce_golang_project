package loms

import (
	"context"
	lomsServiceAPI "route256/checkout/pkg/loms_v1"

	"github.com/pkg/errors"
)

func (c *client) OrderPayed(ctx context.Context, orderID int64) error {
	// response is empty
	_, err := c.lomsClient.OrderPayed(ctx, &lomsServiceAPI.OrderPayedRequest{
		OrderID: orderID,
	})

	if err != nil {
		return errors.WithMessage(err, "rpc order payed")
	}

	return nil
}

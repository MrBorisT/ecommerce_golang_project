package product

import (
	"context"
	productsClient "route256/checkout/internal/clients/grpc/products"
	"route256/checkout/internal/model"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) Product(ctx context.Context, sku uint32) (*model.Product, error) {
	conn, err := grpc.Dial(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.WithMessage(err, "grpc dial")
	}

	defer conn.Close()

	client := productsClient.NewClient(conn, c.token)
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	name, price, err := client.GetProduct(ctxWithTimeout, sku)
	if err != nil {
		return nil, err
	}

	product := model.Product{
		Name:  name,
		Price: price,
	}

	return &product, nil
}

package product

import (
	"context"
	productsClient "route256/checkout/internal/clients/grpc/products"
	"route256/checkout/internal/domain"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	address string
	token   string
}

func New(address, token string) *Client {
	return &Client{
		address: address,
		token:   token,
	}
}

func (c *Client) Product(ctx context.Context, sku uint32) (*domain.Product, error) {
	conn, err := grpc.Dial(c.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := productsClient.NewClient(conn, c.token)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	name, price, err := client.GetProduct(ctx, sku)
	if err != nil {
		return nil, err
	}

	product := domain.Product{
		Name:  name,
		Price: price,
	}

	return &product, nil
}

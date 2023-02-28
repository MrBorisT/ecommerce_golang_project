package product

import (
	"context"
	"fmt"
	"route256/checkout/internal/domain"
	clientwrapper "route256/libs/client"
)

type Client struct {
	url string
}

func New(url string) *Client {
	return &Client{
		url: url,
	}
}

type ProductRequest struct {
	Token string `json:"token"`
	SKU   uint32 `json:"sku"`
}

type ProductItem struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type ProductResponse struct {
	Product ProductItem `json:"product"`
}

func (c *Client) Product(ctx context.Context, sku uint32) (*domain.Product, error) {
	request := ProductRequest{SKU: sku}

	//TODO maybe change Sprint?
	response, err := clientwrapper.SendRequest[ProductRequest, ProductResponse](ctx, request, fmt.Sprint(sku))
	if err != nil {
		return nil, err
	}

	product := domain.Product{
		Name:  response.Product.Name,
		Price: response.Product.Price,
	}

	return &product, nil
}

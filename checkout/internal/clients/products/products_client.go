package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"route256/checkout/internal/domain"

	"github.com/pkg/errors"
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

	rawJSON, err := json.Marshal(request)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling json")
	}

	//TODO maybe change Sprint?
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprint(sku), bytes.NewBuffer(rawJSON))
	if err != nil {
		return nil, errors.Wrap(err, "creating http request")
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, errors.Wrap(err, "calling http")
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", httpResponse.StatusCode)
	}

	var response ProductResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return nil, errors.Wrap(err, "decoding json")
	}

	product := domain.Product{
		Name:  response.Product.Name,
		Price: response.Product.Price,
	}

	return &product, nil
}

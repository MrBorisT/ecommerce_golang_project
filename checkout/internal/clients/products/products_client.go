package product

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

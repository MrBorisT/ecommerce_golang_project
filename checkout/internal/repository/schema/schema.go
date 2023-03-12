package schema

type Cart struct {
	User  int64      `db:"user"`
	Items []CartItem `db:"items"`
}

type CartItem struct {
	SKU   uint32 `db:"sku"`
	Count uint16 `db:"count"`
}

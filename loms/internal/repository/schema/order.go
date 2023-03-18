package schema

type Order struct {
	OrderID int64  `db:"order_id"`
	Status  string `db:"status"` // (new | awaiting payment | failed | payed | cancelled)
	UserID  int64  `db:"user_id"`
}

type OrderItem struct {
	OrderID int64  `db:"order_id"`
	SKU     uint32 `db:"sku"`
	Count   uint16 `db:"count"`
}

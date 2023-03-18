package model

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Stock struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type Product struct {
	Name  string
	Price uint32
}

type Cart struct {
	User  int64
	Items []CartItem
}

type CartItem struct {
	SKU   uint32
	Count uint16
}

func BindCartItemToItem(cartItems []CartItem) []Item {
	result := make([]Item, 0, len(cartItems))
	for _, cartItem := range cartItems {
		result = append(result, Item{
			SKU:   cartItem.SKU,
			Count: cartItem.Count,
		})
	}
	return result
}

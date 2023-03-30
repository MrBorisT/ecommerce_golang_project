package domain

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/model"
	"route256/checkout/internal/repository/schema"
	"route256/libs/pool/batch"
	"sync"
	"sync/atomic"
)

const DEFAULT_AMOUNT_WORKERS = 5

type ItemResponse struct {
	item *model.Item
	err  error
}

type ItemRequest struct {
	ctx  context.Context
	item model.CartItem
}

func (m *service) ListCart(ctx context.Context, user int64) ([]model.Item, uint32, error) {
	// Getting cart of the user (by user's ID)
	cartItems, err := m.CartRepository.ListCart(ctx, user)
	if err != nil {
		return nil, 0, err
	}

	items := make([]model.Item, 0, len(cartItems))

	// Setting workers amount from config.yml
	amountWorkers := config.ConfigData.MaxWorkers
	// if set incorrectly - set to DEFAULT_AMOUNT_WORKERS workers
	if amountWorkers <= 0 {
		amountWorkers = DEFAULT_AMOUNT_WORKERS
	}

	// For task generics (in Task[] brackets):
	// - ItemRequest: context and cart item (with sku and count info)
	// - ItemResponse: full item info (sku, count, name and price) and an error
	tasks := make([]batch.Task[ItemRequest, ItemResponse], 0, len(cartItems))

	// Creating every a task for every cart item
	for _, cartItem := range cartItems {
		itemReq := ItemRequest{
			ctx,
			convertSchemaToModel_CartItem(cartItem),
		}
		tasks = append(tasks, batch.Task[ItemRequest, ItemResponse]{
			Callback: m.getItemInfo,
			InArgs:   itemReq,
		})
	}

	// creating worker pool
	// and a result channel
	batchingPool, results := batch.NewPool[ItemRequest, ItemResponse](ctx, amountWorkers)

	var wg sync.WaitGroup
	wg.Add(1)
	var totalPrice uint32

	go func() {
		defer wg.Done()
		for res := range results {
			if err != nil {
				continue
			}
			if res.err != nil {
				err = res.err
				continue
			}
			items = append(items, *res.item)

			// for race safety
			atomic.AddUint32(&totalPrice, res.item.Price*uint32(res.item.Count))
		}
	}()

	// run tasks in worker pool
	batchingPool.SubmitThenClose(ctx, tasks)

	wg.Wait()

	if err != nil {
		return nil, 0, err
	}
	return items, totalPrice, nil
}

func (m service) getItemInfo(req ItemRequest) ItemResponse {
	productName, productPrice, err := m.ProductChecker.GetProduct(req.ctx, req.item.SKU)
	if err != nil {
		req.ctx.Done()
		return ItemResponse{nil, err}
	}
	return ItemResponse{
		&model.Item{
			SKU:   req.item.SKU,
			Count: req.item.Count,
			Name:  productName,
			Price: productPrice,
		},
		nil,
	}
}

func convertSchemaToModel_CartItem(cartItem schema.CartItems) model.CartItem {
	res := model.CartItem{}
	res.SKU = cartItem.SKU
	res.Count = cartItem.Count

	return res
}

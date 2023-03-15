package domain

import (
	"context"
	"route256/checkout/internal/config"
	"route256/checkout/internal/model"
	"route256/libs/pool/batch"
	"sync"
	"sync/atomic"
)

type ItemResponse struct {
	item *model.Item
	err  error
}

type ItemRequest struct {
	ctx  context.Context
	item model.CartItem
}

func (m *CheckoutService) ListCart(ctx context.Context, user int64) ([]model.Item, uint32, error) {
	// Getting cart of the user (by user's ID)
	cart, err := m.CartRepository.ListCart(ctx, user)
	if err != nil {
		return nil, 0, err
	}

	items := make([]model.Item, 0, len(cart.Items))

	// Setting workers amount from config.yml
	amountWorkers := config.ConfigData.MaxWorkers
	// if set incorrectly - set to 5 workers
	if amountWorkers <= 0 {
		amountWorkers = 5
	}

	// For task generics (in Task[] brackets):
	// - ItemRequest: context and cart item (with sku and count info)
	// - ItemResponse: full item info (sku, count, name and price) and an error
	tasks := make([]batch.Task[ItemRequest, ItemResponse], 0, len(cart.Items))

	// Creating every a task for every cart item
	for _, cartItem := range cart.Items {
		itemReq := ItemRequest{
			ctx,
			cartItem,
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
		for res := range results {
			if res.err != nil && err == nil {
				err = res.err
			}
			items = append(items, *res.item)

			// for race safety
			atomic.AddUint32(&totalPrice, res.item.Price*uint32(res.item.Count))
		}
	}()

	// run tasks in worker pool
	batchingPool.Submit(ctx, tasks)

	<-ctx.Done()
	batchingPool.Close()
	wg.Wait()

	if err != nil {
		return nil, 0, err
	}
	return items, totalPrice, nil
}

func (m *CheckoutService) getItemInfo(req ItemRequest) ItemResponse {
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

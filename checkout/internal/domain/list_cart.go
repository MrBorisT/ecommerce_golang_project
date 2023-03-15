package domain

import (
	"context"
	"route256/checkout/internal/model"
	"route256/libs/pool/batch"
	"sync"
)

type TaskResponse struct {
	*model.Item
	error
}

func (m *CheckoutService) ListCart(ctx context.Context, user int64) ([]model.Item, uint32, error) {
	cart, err := m.CartRepository.ListCart(ctx, user)
	if err != nil {
		return nil, 0, err
	}

	items := make([]model.Item, 0, len(cart.Items))

	amountWorkers := 5

	tasks := make([]batch.Task[model.CartItem, TaskResponse], 0, len(cart.Items))

	for _, cartItem := range cart.Items {
		tasks = append(tasks, batch.Task[model.CartItem, TaskResponse]{
			Callback: func(cartItem model.CartItem) TaskResponse {
				productName, productPrice, err := m.ProductChecker.GetProduct(ctx, cartItem.SKU)
				if err != nil {
					return TaskResponse{nil, err}
				}
				return TaskResponse{
					&model.Item{
						SKU:   cartItem.SKU,
						Count: cartItem.Count,
						Name:  productName,
						Price: productPrice,
					},
					nil,
				}
			},
			InArgs: cartItem,
		})
	}

	batchingPool, results := batch.NewPool[model.CartItem, TaskResponse](ctx, amountWorkers)

	var wg sync.WaitGroup
	wg.Add(1)
	var totalPrice uint32

	go func() {
		for res := range results {
			if res.error != nil {
				err = res.error
			}
			items = append(items, *res.Item)
			totalPrice += res.Price * uint32(res.Count)
		}
	}()

	batchingPool.Submit(ctx, tasks)

	<-ctx.Done()
	batchingPool.Close()
	wg.Wait()

	if err != nil {
		return nil, 0, err
	}
	return items, totalPrice, nil
}

package batch

import (
	"context"
	"sync"
)

type Task[In, Out any] struct {
	Callback func(In) Out
	InArgs   In
}

type Pool[In, Out any] interface {
	Submit(context.Context, []Task[In, Out])
	SubmitThenClose(context.Context, []Task[In, Out])
	Close()
}

var _ Pool[any, any] = &p[any, any]{}

type p[In, Out any] struct {
	amountWorkers int

	wg sync.WaitGroup

	taskSource chan Task[In, Out]
	outSink    chan Out
}

func NewPool[In, Out any](ctx context.Context, amountWorkers int) (Pool[In, Out], <-chan Out) {
	pool := &p[In, Out]{
		amountWorkers: amountWorkers,
	}

	pool.bootstrap(ctx)

	return pool, pool.outSink
}

func (pool *p[In, Out]) Close() {
	// Больше задач не будет
	close(pool.taskSource)

	// Дожидаемся, пока все воркеры закончат работы
	pool.wg.Wait()

	// Закрываем канал на выход, чтобы потребители могли выйти из := range
	close(pool.outSink)
}

// Submit implements Pool
func (pool *p[In, Out]) Submit(ctx context.Context, tasks []Task[In, Out]) {

	// будет запущено pool.amountWorkers горутин
	go func() {
		for _, task := range tasks {
			select {
			case <-ctx.Done():
				return

			case pool.taskSource <- task:
			}
		}
	}()
}

// Submit then close pool
func (pool *p[In, Out]) SubmitThenClose(ctx context.Context, tasks []Task[In, Out]) {
	var wg sync.WaitGroup
	wg.Add(1)
	// будет запущено pool.amountWorkers горутин
	go func() {
		defer wg.Done()
		for _, task := range tasks {
			select {
			case <-ctx.Done():
				return

			case pool.taskSource <- task:
			}
		}
	}()
	wg.Wait()
	pool.Close()
}

func (pool *p[In, Out]) bootstrap(ctx context.Context) {
	pool.taskSource = make(chan Task[In, Out], pool.amountWorkers)
	pool.outSink = make(chan Out, pool.amountWorkers)

	for i := 0; i < pool.amountWorkers; i++ {
		pool.wg.Add(1)
		go func() {
			defer pool.wg.Done()
			worker(ctx, pool.taskSource, pool.outSink)
		}()

	}
}

func worker[In, Out any](
	ctx context.Context,
	taskSource <-chan Task[In, Out],
	resultSink chan<- Out,
) {
	for task := range taskSource {
		select {
		case <-ctx.Done():
			return
		case resultSink <- task.Callback(task.InArgs):
		}
	}
}

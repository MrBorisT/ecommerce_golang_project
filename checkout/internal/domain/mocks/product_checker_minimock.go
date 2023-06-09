package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/domain.ProductChecker -o ./mocks/product_checker_minimock.go -n ProductCheckerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ProductCheckerMock implements domain.ProductChecker
type ProductCheckerMock struct {
	t minimock.Tester

	funcGetProduct          func(ctx context.Context, sku uint32) (s1 string, u1 uint32, err error)
	inspectFuncGetProduct   func(ctx context.Context, sku uint32)
	afterGetProductCounter  uint64
	beforeGetProductCounter uint64
	GetProductMock          mProductCheckerMockGetProduct
}

// NewProductCheckerMock returns a mock for domain.ProductChecker
func NewProductCheckerMock(t minimock.Tester) *ProductCheckerMock {
	m := &ProductCheckerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetProductMock = mProductCheckerMockGetProduct{mock: m}
	m.GetProductMock.callArgs = []*ProductCheckerMockGetProductParams{}

	return m
}

type mProductCheckerMockGetProduct struct {
	mock               *ProductCheckerMock
	defaultExpectation *ProductCheckerMockGetProductExpectation
	expectations       []*ProductCheckerMockGetProductExpectation

	callArgs []*ProductCheckerMockGetProductParams
	mutex    sync.RWMutex
}

// ProductCheckerMockGetProductExpectation specifies expectation struct of the ProductChecker.GetProduct
type ProductCheckerMockGetProductExpectation struct {
	mock    *ProductCheckerMock
	params  *ProductCheckerMockGetProductParams
	results *ProductCheckerMockGetProductResults
	Counter uint64
}

// ProductCheckerMockGetProductParams contains parameters of the ProductChecker.GetProduct
type ProductCheckerMockGetProductParams struct {
	ctx context.Context
	sku uint32
}

// ProductCheckerMockGetProductResults contains results of the ProductChecker.GetProduct
type ProductCheckerMockGetProductResults struct {
	s1  string
	u1  uint32
	err error
}

// Expect sets up expected params for ProductChecker.GetProduct
func (mmGetProduct *mProductCheckerMockGetProduct) Expect(ctx context.Context, sku uint32) *mProductCheckerMockGetProduct {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductCheckerMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductCheckerMockGetProductExpectation{}
	}

	mmGetProduct.defaultExpectation.params = &ProductCheckerMockGetProductParams{ctx, sku}
	for _, e := range mmGetProduct.expectations {
		if minimock.Equal(e.params, mmGetProduct.defaultExpectation.params) {
			mmGetProduct.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetProduct.defaultExpectation.params)
		}
	}

	return mmGetProduct
}

// Inspect accepts an inspector function that has same arguments as the ProductChecker.GetProduct
func (mmGetProduct *mProductCheckerMockGetProduct) Inspect(f func(ctx context.Context, sku uint32)) *mProductCheckerMockGetProduct {
	if mmGetProduct.mock.inspectFuncGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("Inspect function is already set for ProductCheckerMock.GetProduct")
	}

	mmGetProduct.mock.inspectFuncGetProduct = f

	return mmGetProduct
}

// Return sets up results that will be returned by ProductChecker.GetProduct
func (mmGetProduct *mProductCheckerMockGetProduct) Return(s1 string, u1 uint32, err error) *ProductCheckerMock {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductCheckerMock.GetProduct mock is already set by Set")
	}

	if mmGetProduct.defaultExpectation == nil {
		mmGetProduct.defaultExpectation = &ProductCheckerMockGetProductExpectation{mock: mmGetProduct.mock}
	}
	mmGetProduct.defaultExpectation.results = &ProductCheckerMockGetProductResults{s1, u1, err}
	return mmGetProduct.mock
}

// Set uses given function f to mock the ProductChecker.GetProduct method
func (mmGetProduct *mProductCheckerMockGetProduct) Set(f func(ctx context.Context, sku uint32) (s1 string, u1 uint32, err error)) *ProductCheckerMock {
	if mmGetProduct.defaultExpectation != nil {
		mmGetProduct.mock.t.Fatalf("Default expectation is already set for the ProductChecker.GetProduct method")
	}

	if len(mmGetProduct.expectations) > 0 {
		mmGetProduct.mock.t.Fatalf("Some expectations are already set for the ProductChecker.GetProduct method")
	}

	mmGetProduct.mock.funcGetProduct = f
	return mmGetProduct.mock
}

// When sets expectation for the ProductChecker.GetProduct which will trigger the result defined by the following
// Then helper
func (mmGetProduct *mProductCheckerMockGetProduct) When(ctx context.Context, sku uint32) *ProductCheckerMockGetProductExpectation {
	if mmGetProduct.mock.funcGetProduct != nil {
		mmGetProduct.mock.t.Fatalf("ProductCheckerMock.GetProduct mock is already set by Set")
	}

	expectation := &ProductCheckerMockGetProductExpectation{
		mock:   mmGetProduct.mock,
		params: &ProductCheckerMockGetProductParams{ctx, sku},
	}
	mmGetProduct.expectations = append(mmGetProduct.expectations, expectation)
	return expectation
}

// Then sets up ProductChecker.GetProduct return parameters for the expectation previously defined by the When method
func (e *ProductCheckerMockGetProductExpectation) Then(s1 string, u1 uint32, err error) *ProductCheckerMock {
	e.results = &ProductCheckerMockGetProductResults{s1, u1, err}
	return e.mock
}

// GetProduct implements domain.ProductChecker
func (mmGetProduct *ProductCheckerMock) GetProduct(ctx context.Context, sku uint32) (s1 string, u1 uint32, err error) {
	mm_atomic.AddUint64(&mmGetProduct.beforeGetProductCounter, 1)
	defer mm_atomic.AddUint64(&mmGetProduct.afterGetProductCounter, 1)

	if mmGetProduct.inspectFuncGetProduct != nil {
		mmGetProduct.inspectFuncGetProduct(ctx, sku)
	}

	mm_params := &ProductCheckerMockGetProductParams{ctx, sku}

	// Record call args
	mmGetProduct.GetProductMock.mutex.Lock()
	mmGetProduct.GetProductMock.callArgs = append(mmGetProduct.GetProductMock.callArgs, mm_params)
	mmGetProduct.GetProductMock.mutex.Unlock()

	for _, e := range mmGetProduct.GetProductMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.u1, e.results.err
		}
	}

	if mmGetProduct.GetProductMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetProduct.GetProductMock.defaultExpectation.Counter, 1)
		mm_want := mmGetProduct.GetProductMock.defaultExpectation.params
		mm_got := ProductCheckerMockGetProductParams{ctx, sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetProduct.t.Errorf("ProductCheckerMock.GetProduct got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetProduct.GetProductMock.defaultExpectation.results
		if mm_results == nil {
			mmGetProduct.t.Fatal("No results are set for the ProductCheckerMock.GetProduct")
		}
		return (*mm_results).s1, (*mm_results).u1, (*mm_results).err
	}
	if mmGetProduct.funcGetProduct != nil {
		return mmGetProduct.funcGetProduct(ctx, sku)
	}
	mmGetProduct.t.Fatalf("Unexpected call to ProductCheckerMock.GetProduct. %v %v", ctx, sku)
	return
}

// GetProductAfterCounter returns a count of finished ProductCheckerMock.GetProduct invocations
func (mmGetProduct *ProductCheckerMock) GetProductAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.afterGetProductCounter)
}

// GetProductBeforeCounter returns a count of ProductCheckerMock.GetProduct invocations
func (mmGetProduct *ProductCheckerMock) GetProductBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetProduct.beforeGetProductCounter)
}

// Calls returns a list of arguments used in each call to ProductCheckerMock.GetProduct.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetProduct *mProductCheckerMockGetProduct) Calls() []*ProductCheckerMockGetProductParams {
	mmGetProduct.mutex.RLock()

	argCopy := make([]*ProductCheckerMockGetProductParams, len(mmGetProduct.callArgs))
	copy(argCopy, mmGetProduct.callArgs)

	mmGetProduct.mutex.RUnlock()

	return argCopy
}

// MinimockGetProductDone returns true if the count of the GetProduct invocations corresponds
// the number of defined expectations
func (m *ProductCheckerMock) MinimockGetProductDone() bool {
	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetProduct != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetProductInspect logs each unmet expectation
func (m *ProductCheckerMock) MinimockGetProductInspect() {
	for _, e := range m.GetProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ProductCheckerMock.GetProduct with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetProductMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		if m.GetProductMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ProductCheckerMock.GetProduct")
		} else {
			m.t.Errorf("Expected call to ProductCheckerMock.GetProduct with params: %#v", *m.GetProductMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetProduct != nil && mm_atomic.LoadUint64(&m.afterGetProductCounter) < 1 {
		m.t.Error("Expected call to ProductCheckerMock.GetProduct")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ProductCheckerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetProductInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ProductCheckerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ProductCheckerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetProductDone()
}

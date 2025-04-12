// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package addproducts

//go:generate minimock -i github.com/nabishec/avito_pvz_service/internal/http_server/handlers/add_products.PostProducts -o post_products_mock_test.go -n PostProductsMock -p addproducts

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

// PostProductsMock implements PostProducts
type PostProductsMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAddProduct          func(pvzID uuid.UUID, productType string) (pp1 *model.ProductsResp, err error)
	funcAddProductOrigin    string
	inspectFuncAddProduct   func(pvzID uuid.UUID, productType string)
	afterAddProductCounter  uint64
	beforeAddProductCounter uint64
	AddProductMock          mPostProductsMockAddProduct
}

// NewPostProductsMock returns a mock for PostProducts
func NewPostProductsMock(t minimock.Tester) *PostProductsMock {
	m := &PostProductsMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddProductMock = mPostProductsMockAddProduct{mock: m}
	m.AddProductMock.callArgs = []*PostProductsMockAddProductParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mPostProductsMockAddProduct struct {
	optional           bool
	mock               *PostProductsMock
	defaultExpectation *PostProductsMockAddProductExpectation
	expectations       []*PostProductsMockAddProductExpectation

	callArgs []*PostProductsMockAddProductParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// PostProductsMockAddProductExpectation specifies expectation struct of the PostProducts.AddProduct
type PostProductsMockAddProductExpectation struct {
	mock               *PostProductsMock
	params             *PostProductsMockAddProductParams
	paramPtrs          *PostProductsMockAddProductParamPtrs
	expectationOrigins PostProductsMockAddProductExpectationOrigins
	results            *PostProductsMockAddProductResults
	returnOrigin       string
	Counter            uint64
}

// PostProductsMockAddProductParams contains parameters of the PostProducts.AddProduct
type PostProductsMockAddProductParams struct {
	pvzID       uuid.UUID
	productType string
}

// PostProductsMockAddProductParamPtrs contains pointers to parameters of the PostProducts.AddProduct
type PostProductsMockAddProductParamPtrs struct {
	pvzID       *uuid.UUID
	productType *string
}

// PostProductsMockAddProductResults contains results of the PostProducts.AddProduct
type PostProductsMockAddProductResults struct {
	pp1 *model.ProductsResp
	err error
}

// PostProductsMockAddProductOrigins contains origins of expectations of the PostProducts.AddProduct
type PostProductsMockAddProductExpectationOrigins struct {
	origin            string
	originPvzID       string
	originProductType string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmAddProduct *mPostProductsMockAddProduct) Optional() *mPostProductsMockAddProduct {
	mmAddProduct.optional = true
	return mmAddProduct
}

// Expect sets up expected params for PostProducts.AddProduct
func (mmAddProduct *mPostProductsMockAddProduct) Expect(pvzID uuid.UUID, productType string) *mPostProductsMockAddProduct {
	if mmAddProduct.mock.funcAddProduct != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by Set")
	}

	if mmAddProduct.defaultExpectation == nil {
		mmAddProduct.defaultExpectation = &PostProductsMockAddProductExpectation{}
	}

	if mmAddProduct.defaultExpectation.paramPtrs != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by ExpectParams functions")
	}

	mmAddProduct.defaultExpectation.params = &PostProductsMockAddProductParams{pvzID, productType}
	mmAddProduct.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmAddProduct.expectations {
		if minimock.Equal(e.params, mmAddProduct.defaultExpectation.params) {
			mmAddProduct.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddProduct.defaultExpectation.params)
		}
	}

	return mmAddProduct
}

// ExpectPvzIDParam1 sets up expected param pvzID for PostProducts.AddProduct
func (mmAddProduct *mPostProductsMockAddProduct) ExpectPvzIDParam1(pvzID uuid.UUID) *mPostProductsMockAddProduct {
	if mmAddProduct.mock.funcAddProduct != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by Set")
	}

	if mmAddProduct.defaultExpectation == nil {
		mmAddProduct.defaultExpectation = &PostProductsMockAddProductExpectation{}
	}

	if mmAddProduct.defaultExpectation.params != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by Expect")
	}

	if mmAddProduct.defaultExpectation.paramPtrs == nil {
		mmAddProduct.defaultExpectation.paramPtrs = &PostProductsMockAddProductParamPtrs{}
	}
	mmAddProduct.defaultExpectation.paramPtrs.pvzID = &pvzID
	mmAddProduct.defaultExpectation.expectationOrigins.originPvzID = minimock.CallerInfo(1)

	return mmAddProduct
}

// ExpectProductTypeParam2 sets up expected param productType for PostProducts.AddProduct
func (mmAddProduct *mPostProductsMockAddProduct) ExpectProductTypeParam2(productType string) *mPostProductsMockAddProduct {
	if mmAddProduct.mock.funcAddProduct != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by Set")
	}

	if mmAddProduct.defaultExpectation == nil {
		mmAddProduct.defaultExpectation = &PostProductsMockAddProductExpectation{}
	}

	if mmAddProduct.defaultExpectation.params != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by Expect")
	}

	if mmAddProduct.defaultExpectation.paramPtrs == nil {
		mmAddProduct.defaultExpectation.paramPtrs = &PostProductsMockAddProductParamPtrs{}
	}
	mmAddProduct.defaultExpectation.paramPtrs.productType = &productType
	mmAddProduct.defaultExpectation.expectationOrigins.originProductType = minimock.CallerInfo(1)

	return mmAddProduct
}

// Inspect accepts an inspector function that has same arguments as the PostProducts.AddProduct
func (mmAddProduct *mPostProductsMockAddProduct) Inspect(f func(pvzID uuid.UUID, productType string)) *mPostProductsMockAddProduct {
	if mmAddProduct.mock.inspectFuncAddProduct != nil {
		mmAddProduct.mock.t.Fatalf("Inspect function is already set for PostProductsMock.AddProduct")
	}

	mmAddProduct.mock.inspectFuncAddProduct = f

	return mmAddProduct
}

// Return sets up results that will be returned by PostProducts.AddProduct
func (mmAddProduct *mPostProductsMockAddProduct) Return(pp1 *model.ProductsResp, err error) *PostProductsMock {
	if mmAddProduct.mock.funcAddProduct != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by Set")
	}

	if mmAddProduct.defaultExpectation == nil {
		mmAddProduct.defaultExpectation = &PostProductsMockAddProductExpectation{mock: mmAddProduct.mock}
	}
	mmAddProduct.defaultExpectation.results = &PostProductsMockAddProductResults{pp1, err}
	mmAddProduct.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmAddProduct.mock
}

// Set uses given function f to mock the PostProducts.AddProduct method
func (mmAddProduct *mPostProductsMockAddProduct) Set(f func(pvzID uuid.UUID, productType string) (pp1 *model.ProductsResp, err error)) *PostProductsMock {
	if mmAddProduct.defaultExpectation != nil {
		mmAddProduct.mock.t.Fatalf("Default expectation is already set for the PostProducts.AddProduct method")
	}

	if len(mmAddProduct.expectations) > 0 {
		mmAddProduct.mock.t.Fatalf("Some expectations are already set for the PostProducts.AddProduct method")
	}

	mmAddProduct.mock.funcAddProduct = f
	mmAddProduct.mock.funcAddProductOrigin = minimock.CallerInfo(1)
	return mmAddProduct.mock
}

// When sets expectation for the PostProducts.AddProduct which will trigger the result defined by the following
// Then helper
func (mmAddProduct *mPostProductsMockAddProduct) When(pvzID uuid.UUID, productType string) *PostProductsMockAddProductExpectation {
	if mmAddProduct.mock.funcAddProduct != nil {
		mmAddProduct.mock.t.Fatalf("PostProductsMock.AddProduct mock is already set by Set")
	}

	expectation := &PostProductsMockAddProductExpectation{
		mock:               mmAddProduct.mock,
		params:             &PostProductsMockAddProductParams{pvzID, productType},
		expectationOrigins: PostProductsMockAddProductExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmAddProduct.expectations = append(mmAddProduct.expectations, expectation)
	return expectation
}

// Then sets up PostProducts.AddProduct return parameters for the expectation previously defined by the When method
func (e *PostProductsMockAddProductExpectation) Then(pp1 *model.ProductsResp, err error) *PostProductsMock {
	e.results = &PostProductsMockAddProductResults{pp1, err}
	return e.mock
}

// Times sets number of times PostProducts.AddProduct should be invoked
func (mmAddProduct *mPostProductsMockAddProduct) Times(n uint64) *mPostProductsMockAddProduct {
	if n == 0 {
		mmAddProduct.mock.t.Fatalf("Times of PostProductsMock.AddProduct mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmAddProduct.expectedInvocations, n)
	mmAddProduct.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmAddProduct
}

func (mmAddProduct *mPostProductsMockAddProduct) invocationsDone() bool {
	if len(mmAddProduct.expectations) == 0 && mmAddProduct.defaultExpectation == nil && mmAddProduct.mock.funcAddProduct == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmAddProduct.mock.afterAddProductCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmAddProduct.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// AddProduct implements PostProducts
func (mmAddProduct *PostProductsMock) AddProduct(pvzID uuid.UUID, productType string) (pp1 *model.ProductsResp, err error) {
	mm_atomic.AddUint64(&mmAddProduct.beforeAddProductCounter, 1)
	defer mm_atomic.AddUint64(&mmAddProduct.afterAddProductCounter, 1)

	mmAddProduct.t.Helper()

	if mmAddProduct.inspectFuncAddProduct != nil {
		mmAddProduct.inspectFuncAddProduct(pvzID, productType)
	}

	mm_params := PostProductsMockAddProductParams{pvzID, productType}

	// Record call args
	mmAddProduct.AddProductMock.mutex.Lock()
	mmAddProduct.AddProductMock.callArgs = append(mmAddProduct.AddProductMock.callArgs, &mm_params)
	mmAddProduct.AddProductMock.mutex.Unlock()

	for _, e := range mmAddProduct.AddProductMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.pp1, e.results.err
		}
	}

	if mmAddProduct.AddProductMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddProduct.AddProductMock.defaultExpectation.Counter, 1)
		mm_want := mmAddProduct.AddProductMock.defaultExpectation.params
		mm_want_ptrs := mmAddProduct.AddProductMock.defaultExpectation.paramPtrs

		mm_got := PostProductsMockAddProductParams{pvzID, productType}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.pvzID != nil && !minimock.Equal(*mm_want_ptrs.pvzID, mm_got.pvzID) {
				mmAddProduct.t.Errorf("PostProductsMock.AddProduct got unexpected parameter pvzID, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAddProduct.AddProductMock.defaultExpectation.expectationOrigins.originPvzID, *mm_want_ptrs.pvzID, mm_got.pvzID, minimock.Diff(*mm_want_ptrs.pvzID, mm_got.pvzID))
			}

			if mm_want_ptrs.productType != nil && !minimock.Equal(*mm_want_ptrs.productType, mm_got.productType) {
				mmAddProduct.t.Errorf("PostProductsMock.AddProduct got unexpected parameter productType, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAddProduct.AddProductMock.defaultExpectation.expectationOrigins.originProductType, *mm_want_ptrs.productType, mm_got.productType, minimock.Diff(*mm_want_ptrs.productType, mm_got.productType))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddProduct.t.Errorf("PostProductsMock.AddProduct got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmAddProduct.AddProductMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddProduct.AddProductMock.defaultExpectation.results
		if mm_results == nil {
			mmAddProduct.t.Fatal("No results are set for the PostProductsMock.AddProduct")
		}
		return (*mm_results).pp1, (*mm_results).err
	}
	if mmAddProduct.funcAddProduct != nil {
		return mmAddProduct.funcAddProduct(pvzID, productType)
	}
	mmAddProduct.t.Fatalf("Unexpected call to PostProductsMock.AddProduct. %v %v", pvzID, productType)
	return
}

// AddProductAfterCounter returns a count of finished PostProductsMock.AddProduct invocations
func (mmAddProduct *PostProductsMock) AddProductAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddProduct.afterAddProductCounter)
}

// AddProductBeforeCounter returns a count of PostProductsMock.AddProduct invocations
func (mmAddProduct *PostProductsMock) AddProductBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddProduct.beforeAddProductCounter)
}

// Calls returns a list of arguments used in each call to PostProductsMock.AddProduct.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddProduct *mPostProductsMockAddProduct) Calls() []*PostProductsMockAddProductParams {
	mmAddProduct.mutex.RLock()

	argCopy := make([]*PostProductsMockAddProductParams, len(mmAddProduct.callArgs))
	copy(argCopy, mmAddProduct.callArgs)

	mmAddProduct.mutex.RUnlock()

	return argCopy
}

// MinimockAddProductDone returns true if the count of the AddProduct invocations corresponds
// the number of defined expectations
func (m *PostProductsMock) MinimockAddProductDone() bool {
	if m.AddProductMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.AddProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.AddProductMock.invocationsDone()
}

// MinimockAddProductInspect logs each unmet expectation
func (m *PostProductsMock) MinimockAddProductInspect() {
	for _, e := range m.AddProductMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PostProductsMock.AddProduct at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterAddProductCounter := mm_atomic.LoadUint64(&m.afterAddProductCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.AddProductMock.defaultExpectation != nil && afterAddProductCounter < 1 {
		if m.AddProductMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to PostProductsMock.AddProduct at\n%s", m.AddProductMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to PostProductsMock.AddProduct at\n%s with params: %#v", m.AddProductMock.defaultExpectation.expectationOrigins.origin, *m.AddProductMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddProduct != nil && afterAddProductCounter < 1 {
		m.t.Errorf("Expected call to PostProductsMock.AddProduct at\n%s", m.funcAddProductOrigin)
	}

	if !m.AddProductMock.invocationsDone() && afterAddProductCounter > 0 {
		m.t.Errorf("Expected %d calls to PostProductsMock.AddProduct at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.AddProductMock.expectedInvocations), m.AddProductMock.expectedInvocationsOrigin, afterAddProductCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PostProductsMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddProductInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PostProductsMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *PostProductsMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddProductDone()
}

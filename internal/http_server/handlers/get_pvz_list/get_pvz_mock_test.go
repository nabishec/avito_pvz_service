// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package getpvzlist

//go:generate minimock -i github.com/nabishec/avito_pvz_service/internal/http_server/handlers/get_pvz_list.GetPVZ -o get_pvz_mock_test.go -n GetPVZMock -p getpvzlist

import (
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

// GetPVZMock implements GetPVZ
type GetPVZMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcGetPVZListWithRecep          func(startDate time.Time, endDate time.Time, page int, limit int) (ppa1 []*model.PVZWithRecep, err error)
	funcGetPVZListWithRecepOrigin    string
	inspectFuncGetPVZListWithRecep   func(startDate time.Time, endDate time.Time, page int, limit int)
	afterGetPVZListWithRecepCounter  uint64
	beforeGetPVZListWithRecepCounter uint64
	GetPVZListWithRecepMock          mGetPVZMockGetPVZListWithRecep
}

// NewGetPVZMock returns a mock for GetPVZ
func NewGetPVZMock(t minimock.Tester) *GetPVZMock {
	m := &GetPVZMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetPVZListWithRecepMock = mGetPVZMockGetPVZListWithRecep{mock: m}
	m.GetPVZListWithRecepMock.callArgs = []*GetPVZMockGetPVZListWithRecepParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGetPVZMockGetPVZListWithRecep struct {
	optional           bool
	mock               *GetPVZMock
	defaultExpectation *GetPVZMockGetPVZListWithRecepExpectation
	expectations       []*GetPVZMockGetPVZListWithRecepExpectation

	callArgs []*GetPVZMockGetPVZListWithRecepParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// GetPVZMockGetPVZListWithRecepExpectation specifies expectation struct of the GetPVZ.GetPVZListWithRecep
type GetPVZMockGetPVZListWithRecepExpectation struct {
	mock               *GetPVZMock
	params             *GetPVZMockGetPVZListWithRecepParams
	paramPtrs          *GetPVZMockGetPVZListWithRecepParamPtrs
	expectationOrigins GetPVZMockGetPVZListWithRecepExpectationOrigins
	results            *GetPVZMockGetPVZListWithRecepResults
	returnOrigin       string
	Counter            uint64
}

// GetPVZMockGetPVZListWithRecepParams contains parameters of the GetPVZ.GetPVZListWithRecep
type GetPVZMockGetPVZListWithRecepParams struct {
	startDate time.Time
	endDate   time.Time
	page      int
	limit     int
}

// GetPVZMockGetPVZListWithRecepParamPtrs contains pointers to parameters of the GetPVZ.GetPVZListWithRecep
type GetPVZMockGetPVZListWithRecepParamPtrs struct {
	startDate *time.Time
	endDate   *time.Time
	page      *int
	limit     *int
}

// GetPVZMockGetPVZListWithRecepResults contains results of the GetPVZ.GetPVZListWithRecep
type GetPVZMockGetPVZListWithRecepResults struct {
	ppa1 []*model.PVZWithRecep
	err  error
}

// GetPVZMockGetPVZListWithRecepOrigins contains origins of expectations of the GetPVZ.GetPVZListWithRecep
type GetPVZMockGetPVZListWithRecepExpectationOrigins struct {
	origin          string
	originStartDate string
	originEndDate   string
	originPage      string
	originLimit     string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) Optional() *mGetPVZMockGetPVZListWithRecep {
	mmGetPVZListWithRecep.optional = true
	return mmGetPVZListWithRecep
}

// Expect sets up expected params for GetPVZ.GetPVZListWithRecep
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) Expect(startDate time.Time, endDate time.Time, page int, limit int) *mGetPVZMockGetPVZListWithRecep {
	if mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Set")
	}

	if mmGetPVZListWithRecep.defaultExpectation == nil {
		mmGetPVZListWithRecep.defaultExpectation = &GetPVZMockGetPVZListWithRecepExpectation{}
	}

	if mmGetPVZListWithRecep.defaultExpectation.paramPtrs != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by ExpectParams functions")
	}

	mmGetPVZListWithRecep.defaultExpectation.params = &GetPVZMockGetPVZListWithRecepParams{startDate, endDate, page, limit}
	mmGetPVZListWithRecep.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmGetPVZListWithRecep.expectations {
		if minimock.Equal(e.params, mmGetPVZListWithRecep.defaultExpectation.params) {
			mmGetPVZListWithRecep.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetPVZListWithRecep.defaultExpectation.params)
		}
	}

	return mmGetPVZListWithRecep
}

// ExpectStartDateParam1 sets up expected param startDate for GetPVZ.GetPVZListWithRecep
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) ExpectStartDateParam1(startDate time.Time) *mGetPVZMockGetPVZListWithRecep {
	if mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Set")
	}

	if mmGetPVZListWithRecep.defaultExpectation == nil {
		mmGetPVZListWithRecep.defaultExpectation = &GetPVZMockGetPVZListWithRecepExpectation{}
	}

	if mmGetPVZListWithRecep.defaultExpectation.params != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Expect")
	}

	if mmGetPVZListWithRecep.defaultExpectation.paramPtrs == nil {
		mmGetPVZListWithRecep.defaultExpectation.paramPtrs = &GetPVZMockGetPVZListWithRecepParamPtrs{}
	}
	mmGetPVZListWithRecep.defaultExpectation.paramPtrs.startDate = &startDate
	mmGetPVZListWithRecep.defaultExpectation.expectationOrigins.originStartDate = minimock.CallerInfo(1)

	return mmGetPVZListWithRecep
}

// ExpectEndDateParam2 sets up expected param endDate for GetPVZ.GetPVZListWithRecep
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) ExpectEndDateParam2(endDate time.Time) *mGetPVZMockGetPVZListWithRecep {
	if mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Set")
	}

	if mmGetPVZListWithRecep.defaultExpectation == nil {
		mmGetPVZListWithRecep.defaultExpectation = &GetPVZMockGetPVZListWithRecepExpectation{}
	}

	if mmGetPVZListWithRecep.defaultExpectation.params != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Expect")
	}

	if mmGetPVZListWithRecep.defaultExpectation.paramPtrs == nil {
		mmGetPVZListWithRecep.defaultExpectation.paramPtrs = &GetPVZMockGetPVZListWithRecepParamPtrs{}
	}
	mmGetPVZListWithRecep.defaultExpectation.paramPtrs.endDate = &endDate
	mmGetPVZListWithRecep.defaultExpectation.expectationOrigins.originEndDate = minimock.CallerInfo(1)

	return mmGetPVZListWithRecep
}

// ExpectPageParam3 sets up expected param page for GetPVZ.GetPVZListWithRecep
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) ExpectPageParam3(page int) *mGetPVZMockGetPVZListWithRecep {
	if mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Set")
	}

	if mmGetPVZListWithRecep.defaultExpectation == nil {
		mmGetPVZListWithRecep.defaultExpectation = &GetPVZMockGetPVZListWithRecepExpectation{}
	}

	if mmGetPVZListWithRecep.defaultExpectation.params != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Expect")
	}

	if mmGetPVZListWithRecep.defaultExpectation.paramPtrs == nil {
		mmGetPVZListWithRecep.defaultExpectation.paramPtrs = &GetPVZMockGetPVZListWithRecepParamPtrs{}
	}
	mmGetPVZListWithRecep.defaultExpectation.paramPtrs.page = &page
	mmGetPVZListWithRecep.defaultExpectation.expectationOrigins.originPage = minimock.CallerInfo(1)

	return mmGetPVZListWithRecep
}

// ExpectLimitParam4 sets up expected param limit for GetPVZ.GetPVZListWithRecep
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) ExpectLimitParam4(limit int) *mGetPVZMockGetPVZListWithRecep {
	if mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Set")
	}

	if mmGetPVZListWithRecep.defaultExpectation == nil {
		mmGetPVZListWithRecep.defaultExpectation = &GetPVZMockGetPVZListWithRecepExpectation{}
	}

	if mmGetPVZListWithRecep.defaultExpectation.params != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Expect")
	}

	if mmGetPVZListWithRecep.defaultExpectation.paramPtrs == nil {
		mmGetPVZListWithRecep.defaultExpectation.paramPtrs = &GetPVZMockGetPVZListWithRecepParamPtrs{}
	}
	mmGetPVZListWithRecep.defaultExpectation.paramPtrs.limit = &limit
	mmGetPVZListWithRecep.defaultExpectation.expectationOrigins.originLimit = minimock.CallerInfo(1)

	return mmGetPVZListWithRecep
}

// Inspect accepts an inspector function that has same arguments as the GetPVZ.GetPVZListWithRecep
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) Inspect(f func(startDate time.Time, endDate time.Time, page int, limit int)) *mGetPVZMockGetPVZListWithRecep {
	if mmGetPVZListWithRecep.mock.inspectFuncGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("Inspect function is already set for GetPVZMock.GetPVZListWithRecep")
	}

	mmGetPVZListWithRecep.mock.inspectFuncGetPVZListWithRecep = f

	return mmGetPVZListWithRecep
}

// Return sets up results that will be returned by GetPVZ.GetPVZListWithRecep
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) Return(ppa1 []*model.PVZWithRecep, err error) *GetPVZMock {
	if mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Set")
	}

	if mmGetPVZListWithRecep.defaultExpectation == nil {
		mmGetPVZListWithRecep.defaultExpectation = &GetPVZMockGetPVZListWithRecepExpectation{mock: mmGetPVZListWithRecep.mock}
	}
	mmGetPVZListWithRecep.defaultExpectation.results = &GetPVZMockGetPVZListWithRecepResults{ppa1, err}
	mmGetPVZListWithRecep.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmGetPVZListWithRecep.mock
}

// Set uses given function f to mock the GetPVZ.GetPVZListWithRecep method
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) Set(f func(startDate time.Time, endDate time.Time, page int, limit int) (ppa1 []*model.PVZWithRecep, err error)) *GetPVZMock {
	if mmGetPVZListWithRecep.defaultExpectation != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("Default expectation is already set for the GetPVZ.GetPVZListWithRecep method")
	}

	if len(mmGetPVZListWithRecep.expectations) > 0 {
		mmGetPVZListWithRecep.mock.t.Fatalf("Some expectations are already set for the GetPVZ.GetPVZListWithRecep method")
	}

	mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep = f
	mmGetPVZListWithRecep.mock.funcGetPVZListWithRecepOrigin = minimock.CallerInfo(1)
	return mmGetPVZListWithRecep.mock
}

// When sets expectation for the GetPVZ.GetPVZListWithRecep which will trigger the result defined by the following
// Then helper
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) When(startDate time.Time, endDate time.Time, page int, limit int) *GetPVZMockGetPVZListWithRecepExpectation {
	if mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.mock.t.Fatalf("GetPVZMock.GetPVZListWithRecep mock is already set by Set")
	}

	expectation := &GetPVZMockGetPVZListWithRecepExpectation{
		mock:               mmGetPVZListWithRecep.mock,
		params:             &GetPVZMockGetPVZListWithRecepParams{startDate, endDate, page, limit},
		expectationOrigins: GetPVZMockGetPVZListWithRecepExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmGetPVZListWithRecep.expectations = append(mmGetPVZListWithRecep.expectations, expectation)
	return expectation
}

// Then sets up GetPVZ.GetPVZListWithRecep return parameters for the expectation previously defined by the When method
func (e *GetPVZMockGetPVZListWithRecepExpectation) Then(ppa1 []*model.PVZWithRecep, err error) *GetPVZMock {
	e.results = &GetPVZMockGetPVZListWithRecepResults{ppa1, err}
	return e.mock
}

// Times sets number of times GetPVZ.GetPVZListWithRecep should be invoked
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) Times(n uint64) *mGetPVZMockGetPVZListWithRecep {
	if n == 0 {
		mmGetPVZListWithRecep.mock.t.Fatalf("Times of GetPVZMock.GetPVZListWithRecep mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGetPVZListWithRecep.expectedInvocations, n)
	mmGetPVZListWithRecep.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmGetPVZListWithRecep
}

func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) invocationsDone() bool {
	if len(mmGetPVZListWithRecep.expectations) == 0 && mmGetPVZListWithRecep.defaultExpectation == nil && mmGetPVZListWithRecep.mock.funcGetPVZListWithRecep == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGetPVZListWithRecep.mock.afterGetPVZListWithRecepCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGetPVZListWithRecep.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// GetPVZListWithRecep implements GetPVZ
func (mmGetPVZListWithRecep *GetPVZMock) GetPVZListWithRecep(startDate time.Time, endDate time.Time, page int, limit int) (ppa1 []*model.PVZWithRecep, err error) {
	mm_atomic.AddUint64(&mmGetPVZListWithRecep.beforeGetPVZListWithRecepCounter, 1)
	defer mm_atomic.AddUint64(&mmGetPVZListWithRecep.afterGetPVZListWithRecepCounter, 1)

	mmGetPVZListWithRecep.t.Helper()

	if mmGetPVZListWithRecep.inspectFuncGetPVZListWithRecep != nil {
		mmGetPVZListWithRecep.inspectFuncGetPVZListWithRecep(startDate, endDate, page, limit)
	}

	mm_params := GetPVZMockGetPVZListWithRecepParams{startDate, endDate, page, limit}

	// Record call args
	mmGetPVZListWithRecep.GetPVZListWithRecepMock.mutex.Lock()
	mmGetPVZListWithRecep.GetPVZListWithRecepMock.callArgs = append(mmGetPVZListWithRecep.GetPVZListWithRecepMock.callArgs, &mm_params)
	mmGetPVZListWithRecep.GetPVZListWithRecepMock.mutex.Unlock()

	for _, e := range mmGetPVZListWithRecep.GetPVZListWithRecepMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ppa1, e.results.err
		}
	}

	if mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.Counter, 1)
		mm_want := mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.params
		mm_want_ptrs := mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.paramPtrs

		mm_got := GetPVZMockGetPVZListWithRecepParams{startDate, endDate, page, limit}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.startDate != nil && !minimock.Equal(*mm_want_ptrs.startDate, mm_got.startDate) {
				mmGetPVZListWithRecep.t.Errorf("GetPVZMock.GetPVZListWithRecep got unexpected parameter startDate, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.expectationOrigins.originStartDate, *mm_want_ptrs.startDate, mm_got.startDate, minimock.Diff(*mm_want_ptrs.startDate, mm_got.startDate))
			}

			if mm_want_ptrs.endDate != nil && !minimock.Equal(*mm_want_ptrs.endDate, mm_got.endDate) {
				mmGetPVZListWithRecep.t.Errorf("GetPVZMock.GetPVZListWithRecep got unexpected parameter endDate, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.expectationOrigins.originEndDate, *mm_want_ptrs.endDate, mm_got.endDate, minimock.Diff(*mm_want_ptrs.endDate, mm_got.endDate))
			}

			if mm_want_ptrs.page != nil && !minimock.Equal(*mm_want_ptrs.page, mm_got.page) {
				mmGetPVZListWithRecep.t.Errorf("GetPVZMock.GetPVZListWithRecep got unexpected parameter page, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.expectationOrigins.originPage, *mm_want_ptrs.page, mm_got.page, minimock.Diff(*mm_want_ptrs.page, mm_got.page))
			}

			if mm_want_ptrs.limit != nil && !minimock.Equal(*mm_want_ptrs.limit, mm_got.limit) {
				mmGetPVZListWithRecep.t.Errorf("GetPVZMock.GetPVZListWithRecep got unexpected parameter limit, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.expectationOrigins.originLimit, *mm_want_ptrs.limit, mm_got.limit, minimock.Diff(*mm_want_ptrs.limit, mm_got.limit))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetPVZListWithRecep.t.Errorf("GetPVZMock.GetPVZListWithRecep got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetPVZListWithRecep.GetPVZListWithRecepMock.defaultExpectation.results
		if mm_results == nil {
			mmGetPVZListWithRecep.t.Fatal("No results are set for the GetPVZMock.GetPVZListWithRecep")
		}
		return (*mm_results).ppa1, (*mm_results).err
	}
	if mmGetPVZListWithRecep.funcGetPVZListWithRecep != nil {
		return mmGetPVZListWithRecep.funcGetPVZListWithRecep(startDate, endDate, page, limit)
	}
	mmGetPVZListWithRecep.t.Fatalf("Unexpected call to GetPVZMock.GetPVZListWithRecep. %v %v %v %v", startDate, endDate, page, limit)
	return
}

// GetPVZListWithRecepAfterCounter returns a count of finished GetPVZMock.GetPVZListWithRecep invocations
func (mmGetPVZListWithRecep *GetPVZMock) GetPVZListWithRecepAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPVZListWithRecep.afterGetPVZListWithRecepCounter)
}

// GetPVZListWithRecepBeforeCounter returns a count of GetPVZMock.GetPVZListWithRecep invocations
func (mmGetPVZListWithRecep *GetPVZMock) GetPVZListWithRecepBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPVZListWithRecep.beforeGetPVZListWithRecepCounter)
}

// Calls returns a list of arguments used in each call to GetPVZMock.GetPVZListWithRecep.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetPVZListWithRecep *mGetPVZMockGetPVZListWithRecep) Calls() []*GetPVZMockGetPVZListWithRecepParams {
	mmGetPVZListWithRecep.mutex.RLock()

	argCopy := make([]*GetPVZMockGetPVZListWithRecepParams, len(mmGetPVZListWithRecep.callArgs))
	copy(argCopy, mmGetPVZListWithRecep.callArgs)

	mmGetPVZListWithRecep.mutex.RUnlock()

	return argCopy
}

// MinimockGetPVZListWithRecepDone returns true if the count of the GetPVZListWithRecep invocations corresponds
// the number of defined expectations
func (m *GetPVZMock) MinimockGetPVZListWithRecepDone() bool {
	if m.GetPVZListWithRecepMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetPVZListWithRecepMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetPVZListWithRecepMock.invocationsDone()
}

// MinimockGetPVZListWithRecepInspect logs each unmet expectation
func (m *GetPVZMock) MinimockGetPVZListWithRecepInspect() {
	for _, e := range m.GetPVZListWithRecepMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GetPVZMock.GetPVZListWithRecep at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterGetPVZListWithRecepCounter := mm_atomic.LoadUint64(&m.afterGetPVZListWithRecepCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetPVZListWithRecepMock.defaultExpectation != nil && afterGetPVZListWithRecepCounter < 1 {
		if m.GetPVZListWithRecepMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to GetPVZMock.GetPVZListWithRecep at\n%s", m.GetPVZListWithRecepMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to GetPVZMock.GetPVZListWithRecep at\n%s with params: %#v", m.GetPVZListWithRecepMock.defaultExpectation.expectationOrigins.origin, *m.GetPVZListWithRecepMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPVZListWithRecep != nil && afterGetPVZListWithRecepCounter < 1 {
		m.t.Errorf("Expected call to GetPVZMock.GetPVZListWithRecep at\n%s", m.funcGetPVZListWithRecepOrigin)
	}

	if !m.GetPVZListWithRecepMock.invocationsDone() && afterGetPVZListWithRecepCounter > 0 {
		m.t.Errorf("Expected %d calls to GetPVZMock.GetPVZListWithRecep at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.GetPVZListWithRecepMock.expectedInvocations), m.GetPVZListWithRecepMock.expectedInvocationsOrigin, afterGetPVZListWithRecepCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GetPVZMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockGetPVZListWithRecepInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GetPVZMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *GetPVZMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetPVZListWithRecepDone()
}

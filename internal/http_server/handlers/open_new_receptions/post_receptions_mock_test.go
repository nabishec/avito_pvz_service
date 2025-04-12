// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package opennewreceptions

//go:generate minimock -i github.com/nabishec/avito_pvz_service/internal/http_server/handlers/open_new_receptions.PostReceptions -o post_receptions_mock_test.go -n PostReceptionsMock -p opennewreceptions

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

// PostReceptionsMock implements PostReceptions
type PostReceptionsMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAddReception          func(pvzID uuid.UUID) (rp1 *model.ReceptionsResp, err error)
	funcAddReceptionOrigin    string
	inspectFuncAddReception   func(pvzID uuid.UUID)
	afterAddReceptionCounter  uint64
	beforeAddReceptionCounter uint64
	AddReceptionMock          mPostReceptionsMockAddReception
}

// NewPostReceptionsMock returns a mock for PostReceptions
func NewPostReceptionsMock(t minimock.Tester) *PostReceptionsMock {
	m := &PostReceptionsMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddReceptionMock = mPostReceptionsMockAddReception{mock: m}
	m.AddReceptionMock.callArgs = []*PostReceptionsMockAddReceptionParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mPostReceptionsMockAddReception struct {
	optional           bool
	mock               *PostReceptionsMock
	defaultExpectation *PostReceptionsMockAddReceptionExpectation
	expectations       []*PostReceptionsMockAddReceptionExpectation

	callArgs []*PostReceptionsMockAddReceptionParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// PostReceptionsMockAddReceptionExpectation specifies expectation struct of the PostReceptions.AddReception
type PostReceptionsMockAddReceptionExpectation struct {
	mock               *PostReceptionsMock
	params             *PostReceptionsMockAddReceptionParams
	paramPtrs          *PostReceptionsMockAddReceptionParamPtrs
	expectationOrigins PostReceptionsMockAddReceptionExpectationOrigins
	results            *PostReceptionsMockAddReceptionResults
	returnOrigin       string
	Counter            uint64
}

// PostReceptionsMockAddReceptionParams contains parameters of the PostReceptions.AddReception
type PostReceptionsMockAddReceptionParams struct {
	pvzID uuid.UUID
}

// PostReceptionsMockAddReceptionParamPtrs contains pointers to parameters of the PostReceptions.AddReception
type PostReceptionsMockAddReceptionParamPtrs struct {
	pvzID *uuid.UUID
}

// PostReceptionsMockAddReceptionResults contains results of the PostReceptions.AddReception
type PostReceptionsMockAddReceptionResults struct {
	rp1 *model.ReceptionsResp
	err error
}

// PostReceptionsMockAddReceptionOrigins contains origins of expectations of the PostReceptions.AddReception
type PostReceptionsMockAddReceptionExpectationOrigins struct {
	origin      string
	originPvzID string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmAddReception *mPostReceptionsMockAddReception) Optional() *mPostReceptionsMockAddReception {
	mmAddReception.optional = true
	return mmAddReception
}

// Expect sets up expected params for PostReceptions.AddReception
func (mmAddReception *mPostReceptionsMockAddReception) Expect(pvzID uuid.UUID) *mPostReceptionsMockAddReception {
	if mmAddReception.mock.funcAddReception != nil {
		mmAddReception.mock.t.Fatalf("PostReceptionsMock.AddReception mock is already set by Set")
	}

	if mmAddReception.defaultExpectation == nil {
		mmAddReception.defaultExpectation = &PostReceptionsMockAddReceptionExpectation{}
	}

	if mmAddReception.defaultExpectation.paramPtrs != nil {
		mmAddReception.mock.t.Fatalf("PostReceptionsMock.AddReception mock is already set by ExpectParams functions")
	}

	mmAddReception.defaultExpectation.params = &PostReceptionsMockAddReceptionParams{pvzID}
	mmAddReception.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmAddReception.expectations {
		if minimock.Equal(e.params, mmAddReception.defaultExpectation.params) {
			mmAddReception.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddReception.defaultExpectation.params)
		}
	}

	return mmAddReception
}

// ExpectPvzIDParam1 sets up expected param pvzID for PostReceptions.AddReception
func (mmAddReception *mPostReceptionsMockAddReception) ExpectPvzIDParam1(pvzID uuid.UUID) *mPostReceptionsMockAddReception {
	if mmAddReception.mock.funcAddReception != nil {
		mmAddReception.mock.t.Fatalf("PostReceptionsMock.AddReception mock is already set by Set")
	}

	if mmAddReception.defaultExpectation == nil {
		mmAddReception.defaultExpectation = &PostReceptionsMockAddReceptionExpectation{}
	}

	if mmAddReception.defaultExpectation.params != nil {
		mmAddReception.mock.t.Fatalf("PostReceptionsMock.AddReception mock is already set by Expect")
	}

	if mmAddReception.defaultExpectation.paramPtrs == nil {
		mmAddReception.defaultExpectation.paramPtrs = &PostReceptionsMockAddReceptionParamPtrs{}
	}
	mmAddReception.defaultExpectation.paramPtrs.pvzID = &pvzID
	mmAddReception.defaultExpectation.expectationOrigins.originPvzID = minimock.CallerInfo(1)

	return mmAddReception
}

// Inspect accepts an inspector function that has same arguments as the PostReceptions.AddReception
func (mmAddReception *mPostReceptionsMockAddReception) Inspect(f func(pvzID uuid.UUID)) *mPostReceptionsMockAddReception {
	if mmAddReception.mock.inspectFuncAddReception != nil {
		mmAddReception.mock.t.Fatalf("Inspect function is already set for PostReceptionsMock.AddReception")
	}

	mmAddReception.mock.inspectFuncAddReception = f

	return mmAddReception
}

// Return sets up results that will be returned by PostReceptions.AddReception
func (mmAddReception *mPostReceptionsMockAddReception) Return(rp1 *model.ReceptionsResp, err error) *PostReceptionsMock {
	if mmAddReception.mock.funcAddReception != nil {
		mmAddReception.mock.t.Fatalf("PostReceptionsMock.AddReception mock is already set by Set")
	}

	if mmAddReception.defaultExpectation == nil {
		mmAddReception.defaultExpectation = &PostReceptionsMockAddReceptionExpectation{mock: mmAddReception.mock}
	}
	mmAddReception.defaultExpectation.results = &PostReceptionsMockAddReceptionResults{rp1, err}
	mmAddReception.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmAddReception.mock
}

// Set uses given function f to mock the PostReceptions.AddReception method
func (mmAddReception *mPostReceptionsMockAddReception) Set(f func(pvzID uuid.UUID) (rp1 *model.ReceptionsResp, err error)) *PostReceptionsMock {
	if mmAddReception.defaultExpectation != nil {
		mmAddReception.mock.t.Fatalf("Default expectation is already set for the PostReceptions.AddReception method")
	}

	if len(mmAddReception.expectations) > 0 {
		mmAddReception.mock.t.Fatalf("Some expectations are already set for the PostReceptions.AddReception method")
	}

	mmAddReception.mock.funcAddReception = f
	mmAddReception.mock.funcAddReceptionOrigin = minimock.CallerInfo(1)
	return mmAddReception.mock
}

// When sets expectation for the PostReceptions.AddReception which will trigger the result defined by the following
// Then helper
func (mmAddReception *mPostReceptionsMockAddReception) When(pvzID uuid.UUID) *PostReceptionsMockAddReceptionExpectation {
	if mmAddReception.mock.funcAddReception != nil {
		mmAddReception.mock.t.Fatalf("PostReceptionsMock.AddReception mock is already set by Set")
	}

	expectation := &PostReceptionsMockAddReceptionExpectation{
		mock:               mmAddReception.mock,
		params:             &PostReceptionsMockAddReceptionParams{pvzID},
		expectationOrigins: PostReceptionsMockAddReceptionExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmAddReception.expectations = append(mmAddReception.expectations, expectation)
	return expectation
}

// Then sets up PostReceptions.AddReception return parameters for the expectation previously defined by the When method
func (e *PostReceptionsMockAddReceptionExpectation) Then(rp1 *model.ReceptionsResp, err error) *PostReceptionsMock {
	e.results = &PostReceptionsMockAddReceptionResults{rp1, err}
	return e.mock
}

// Times sets number of times PostReceptions.AddReception should be invoked
func (mmAddReception *mPostReceptionsMockAddReception) Times(n uint64) *mPostReceptionsMockAddReception {
	if n == 0 {
		mmAddReception.mock.t.Fatalf("Times of PostReceptionsMock.AddReception mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmAddReception.expectedInvocations, n)
	mmAddReception.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmAddReception
}

func (mmAddReception *mPostReceptionsMockAddReception) invocationsDone() bool {
	if len(mmAddReception.expectations) == 0 && mmAddReception.defaultExpectation == nil && mmAddReception.mock.funcAddReception == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmAddReception.mock.afterAddReceptionCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmAddReception.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// AddReception implements PostReceptions
func (mmAddReception *PostReceptionsMock) AddReception(pvzID uuid.UUID) (rp1 *model.ReceptionsResp, err error) {
	mm_atomic.AddUint64(&mmAddReception.beforeAddReceptionCounter, 1)
	defer mm_atomic.AddUint64(&mmAddReception.afterAddReceptionCounter, 1)

	mmAddReception.t.Helper()

	if mmAddReception.inspectFuncAddReception != nil {
		mmAddReception.inspectFuncAddReception(pvzID)
	}

	mm_params := PostReceptionsMockAddReceptionParams{pvzID}

	// Record call args
	mmAddReception.AddReceptionMock.mutex.Lock()
	mmAddReception.AddReceptionMock.callArgs = append(mmAddReception.AddReceptionMock.callArgs, &mm_params)
	mmAddReception.AddReceptionMock.mutex.Unlock()

	for _, e := range mmAddReception.AddReceptionMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.rp1, e.results.err
		}
	}

	if mmAddReception.AddReceptionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddReception.AddReceptionMock.defaultExpectation.Counter, 1)
		mm_want := mmAddReception.AddReceptionMock.defaultExpectation.params
		mm_want_ptrs := mmAddReception.AddReceptionMock.defaultExpectation.paramPtrs

		mm_got := PostReceptionsMockAddReceptionParams{pvzID}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.pvzID != nil && !minimock.Equal(*mm_want_ptrs.pvzID, mm_got.pvzID) {
				mmAddReception.t.Errorf("PostReceptionsMock.AddReception got unexpected parameter pvzID, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAddReception.AddReceptionMock.defaultExpectation.expectationOrigins.originPvzID, *mm_want_ptrs.pvzID, mm_got.pvzID, minimock.Diff(*mm_want_ptrs.pvzID, mm_got.pvzID))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddReception.t.Errorf("PostReceptionsMock.AddReception got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmAddReception.AddReceptionMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddReception.AddReceptionMock.defaultExpectation.results
		if mm_results == nil {
			mmAddReception.t.Fatal("No results are set for the PostReceptionsMock.AddReception")
		}
		return (*mm_results).rp1, (*mm_results).err
	}
	if mmAddReception.funcAddReception != nil {
		return mmAddReception.funcAddReception(pvzID)
	}
	mmAddReception.t.Fatalf("Unexpected call to PostReceptionsMock.AddReception. %v", pvzID)
	return
}

// AddReceptionAfterCounter returns a count of finished PostReceptionsMock.AddReception invocations
func (mmAddReception *PostReceptionsMock) AddReceptionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddReception.afterAddReceptionCounter)
}

// AddReceptionBeforeCounter returns a count of PostReceptionsMock.AddReception invocations
func (mmAddReception *PostReceptionsMock) AddReceptionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddReception.beforeAddReceptionCounter)
}

// Calls returns a list of arguments used in each call to PostReceptionsMock.AddReception.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddReception *mPostReceptionsMockAddReception) Calls() []*PostReceptionsMockAddReceptionParams {
	mmAddReception.mutex.RLock()

	argCopy := make([]*PostReceptionsMockAddReceptionParams, len(mmAddReception.callArgs))
	copy(argCopy, mmAddReception.callArgs)

	mmAddReception.mutex.RUnlock()

	return argCopy
}

// MinimockAddReceptionDone returns true if the count of the AddReception invocations corresponds
// the number of defined expectations
func (m *PostReceptionsMock) MinimockAddReceptionDone() bool {
	if m.AddReceptionMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.AddReceptionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.AddReceptionMock.invocationsDone()
}

// MinimockAddReceptionInspect logs each unmet expectation
func (m *PostReceptionsMock) MinimockAddReceptionInspect() {
	for _, e := range m.AddReceptionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PostReceptionsMock.AddReception at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterAddReceptionCounter := mm_atomic.LoadUint64(&m.afterAddReceptionCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.AddReceptionMock.defaultExpectation != nil && afterAddReceptionCounter < 1 {
		if m.AddReceptionMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to PostReceptionsMock.AddReception at\n%s", m.AddReceptionMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to PostReceptionsMock.AddReception at\n%s with params: %#v", m.AddReceptionMock.defaultExpectation.expectationOrigins.origin, *m.AddReceptionMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddReception != nil && afterAddReceptionCounter < 1 {
		m.t.Errorf("Expected call to PostReceptionsMock.AddReception at\n%s", m.funcAddReceptionOrigin)
	}

	if !m.AddReceptionMock.invocationsDone() && afterAddReceptionCounter > 0 {
		m.t.Errorf("Expected %d calls to PostReceptionsMock.AddReception at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.AddReceptionMock.expectedInvocations), m.AddReceptionMock.expectedInvocationsOrigin, afterAddReceptionCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PostReceptionsMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddReceptionInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PostReceptionsMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *PostReceptionsMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddReceptionDone()
}

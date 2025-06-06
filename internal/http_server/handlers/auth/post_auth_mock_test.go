// Code generated by http://github.com/gojuno/minimock (v3.4.5). DO NOT EDIT.

package auth

//go:generate minimock -i github.com/nabishec/avito_pvz_service/internal/http_server/handlers/auth.PostAuth -o post_auth_mock_test.go -n PostAuthMock -p auth

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/nabishec/avito_pvz_service/internal/model"
)

// PostAuthMock implements PostAuth
type PostAuthMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateUser          func(ctx context.Context, email string, password string, role string) (rp1 *model.RegisterResp, err error)
	funcCreateUserOrigin    string
	inspectFuncCreateUser   func(ctx context.Context, email string, password string, role string)
	afterCreateUserCounter  uint64
	beforeCreateUserCounter uint64
	CreateUserMock          mPostAuthMockCreateUser
}

// NewPostAuthMock returns a mock for PostAuth
func NewPostAuthMock(t minimock.Tester) *PostAuthMock {
	m := &PostAuthMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateUserMock = mPostAuthMockCreateUser{mock: m}
	m.CreateUserMock.callArgs = []*PostAuthMockCreateUserParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mPostAuthMockCreateUser struct {
	optional           bool
	mock               *PostAuthMock
	defaultExpectation *PostAuthMockCreateUserExpectation
	expectations       []*PostAuthMockCreateUserExpectation

	callArgs []*PostAuthMockCreateUserParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// PostAuthMockCreateUserExpectation specifies expectation struct of the PostAuth.CreateUser
type PostAuthMockCreateUserExpectation struct {
	mock               *PostAuthMock
	params             *PostAuthMockCreateUserParams
	paramPtrs          *PostAuthMockCreateUserParamPtrs
	expectationOrigins PostAuthMockCreateUserExpectationOrigins
	results            *PostAuthMockCreateUserResults
	returnOrigin       string
	Counter            uint64
}

// PostAuthMockCreateUserParams contains parameters of the PostAuth.CreateUser
type PostAuthMockCreateUserParams struct {
	ctx      context.Context
	email    string
	password string
	role     string
}

// PostAuthMockCreateUserParamPtrs contains pointers to parameters of the PostAuth.CreateUser
type PostAuthMockCreateUserParamPtrs struct {
	ctx      *context.Context
	email    *string
	password *string
	role     *string
}

// PostAuthMockCreateUserResults contains results of the PostAuth.CreateUser
type PostAuthMockCreateUserResults struct {
	rp1 *model.RegisterResp
	err error
}

// PostAuthMockCreateUserOrigins contains origins of expectations of the PostAuth.CreateUser
type PostAuthMockCreateUserExpectationOrigins struct {
	origin         string
	originCtx      string
	originEmail    string
	originPassword string
	originRole     string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreateUser *mPostAuthMockCreateUser) Optional() *mPostAuthMockCreateUser {
	mmCreateUser.optional = true
	return mmCreateUser
}

// Expect sets up expected params for PostAuth.CreateUser
func (mmCreateUser *mPostAuthMockCreateUser) Expect(ctx context.Context, email string, password string, role string) *mPostAuthMockCreateUser {
	if mmCreateUser.mock.funcCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Set")
	}

	if mmCreateUser.defaultExpectation == nil {
		mmCreateUser.defaultExpectation = &PostAuthMockCreateUserExpectation{}
	}

	if mmCreateUser.defaultExpectation.paramPtrs != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by ExpectParams functions")
	}

	mmCreateUser.defaultExpectation.params = &PostAuthMockCreateUserParams{ctx, email, password, role}
	mmCreateUser.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreateUser.expectations {
		if minimock.Equal(e.params, mmCreateUser.defaultExpectation.params) {
			mmCreateUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateUser.defaultExpectation.params)
		}
	}

	return mmCreateUser
}

// ExpectCtxParam1 sets up expected param ctx for PostAuth.CreateUser
func (mmCreateUser *mPostAuthMockCreateUser) ExpectCtxParam1(ctx context.Context) *mPostAuthMockCreateUser {
	if mmCreateUser.mock.funcCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Set")
	}

	if mmCreateUser.defaultExpectation == nil {
		mmCreateUser.defaultExpectation = &PostAuthMockCreateUserExpectation{}
	}

	if mmCreateUser.defaultExpectation.params != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Expect")
	}

	if mmCreateUser.defaultExpectation.paramPtrs == nil {
		mmCreateUser.defaultExpectation.paramPtrs = &PostAuthMockCreateUserParamPtrs{}
	}
	mmCreateUser.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreateUser.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreateUser
}

// ExpectEmailParam2 sets up expected param email for PostAuth.CreateUser
func (mmCreateUser *mPostAuthMockCreateUser) ExpectEmailParam2(email string) *mPostAuthMockCreateUser {
	if mmCreateUser.mock.funcCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Set")
	}

	if mmCreateUser.defaultExpectation == nil {
		mmCreateUser.defaultExpectation = &PostAuthMockCreateUserExpectation{}
	}

	if mmCreateUser.defaultExpectation.params != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Expect")
	}

	if mmCreateUser.defaultExpectation.paramPtrs == nil {
		mmCreateUser.defaultExpectation.paramPtrs = &PostAuthMockCreateUserParamPtrs{}
	}
	mmCreateUser.defaultExpectation.paramPtrs.email = &email
	mmCreateUser.defaultExpectation.expectationOrigins.originEmail = minimock.CallerInfo(1)

	return mmCreateUser
}

// ExpectPasswordParam3 sets up expected param password for PostAuth.CreateUser
func (mmCreateUser *mPostAuthMockCreateUser) ExpectPasswordParam3(password string) *mPostAuthMockCreateUser {
	if mmCreateUser.mock.funcCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Set")
	}

	if mmCreateUser.defaultExpectation == nil {
		mmCreateUser.defaultExpectation = &PostAuthMockCreateUserExpectation{}
	}

	if mmCreateUser.defaultExpectation.params != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Expect")
	}

	if mmCreateUser.defaultExpectation.paramPtrs == nil {
		mmCreateUser.defaultExpectation.paramPtrs = &PostAuthMockCreateUserParamPtrs{}
	}
	mmCreateUser.defaultExpectation.paramPtrs.password = &password
	mmCreateUser.defaultExpectation.expectationOrigins.originPassword = minimock.CallerInfo(1)

	return mmCreateUser
}

// ExpectRoleParam4 sets up expected param role for PostAuth.CreateUser
func (mmCreateUser *mPostAuthMockCreateUser) ExpectRoleParam4(role string) *mPostAuthMockCreateUser {
	if mmCreateUser.mock.funcCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Set")
	}

	if mmCreateUser.defaultExpectation == nil {
		mmCreateUser.defaultExpectation = &PostAuthMockCreateUserExpectation{}
	}

	if mmCreateUser.defaultExpectation.params != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Expect")
	}

	if mmCreateUser.defaultExpectation.paramPtrs == nil {
		mmCreateUser.defaultExpectation.paramPtrs = &PostAuthMockCreateUserParamPtrs{}
	}
	mmCreateUser.defaultExpectation.paramPtrs.role = &role
	mmCreateUser.defaultExpectation.expectationOrigins.originRole = minimock.CallerInfo(1)

	return mmCreateUser
}

// Inspect accepts an inspector function that has same arguments as the PostAuth.CreateUser
func (mmCreateUser *mPostAuthMockCreateUser) Inspect(f func(ctx context.Context, email string, password string, role string)) *mPostAuthMockCreateUser {
	if mmCreateUser.mock.inspectFuncCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("Inspect function is already set for PostAuthMock.CreateUser")
	}

	mmCreateUser.mock.inspectFuncCreateUser = f

	return mmCreateUser
}

// Return sets up results that will be returned by PostAuth.CreateUser
func (mmCreateUser *mPostAuthMockCreateUser) Return(rp1 *model.RegisterResp, err error) *PostAuthMock {
	if mmCreateUser.mock.funcCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Set")
	}

	if mmCreateUser.defaultExpectation == nil {
		mmCreateUser.defaultExpectation = &PostAuthMockCreateUserExpectation{mock: mmCreateUser.mock}
	}
	mmCreateUser.defaultExpectation.results = &PostAuthMockCreateUserResults{rp1, err}
	mmCreateUser.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreateUser.mock
}

// Set uses given function f to mock the PostAuth.CreateUser method
func (mmCreateUser *mPostAuthMockCreateUser) Set(f func(ctx context.Context, email string, password string, role string) (rp1 *model.RegisterResp, err error)) *PostAuthMock {
	if mmCreateUser.defaultExpectation != nil {
		mmCreateUser.mock.t.Fatalf("Default expectation is already set for the PostAuth.CreateUser method")
	}

	if len(mmCreateUser.expectations) > 0 {
		mmCreateUser.mock.t.Fatalf("Some expectations are already set for the PostAuth.CreateUser method")
	}

	mmCreateUser.mock.funcCreateUser = f
	mmCreateUser.mock.funcCreateUserOrigin = minimock.CallerInfo(1)
	return mmCreateUser.mock
}

// When sets expectation for the PostAuth.CreateUser which will trigger the result defined by the following
// Then helper
func (mmCreateUser *mPostAuthMockCreateUser) When(ctx context.Context, email string, password string, role string) *PostAuthMockCreateUserExpectation {
	if mmCreateUser.mock.funcCreateUser != nil {
		mmCreateUser.mock.t.Fatalf("PostAuthMock.CreateUser mock is already set by Set")
	}

	expectation := &PostAuthMockCreateUserExpectation{
		mock:               mmCreateUser.mock,
		params:             &PostAuthMockCreateUserParams{ctx, email, password, role},
		expectationOrigins: PostAuthMockCreateUserExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreateUser.expectations = append(mmCreateUser.expectations, expectation)
	return expectation
}

// Then sets up PostAuth.CreateUser return parameters for the expectation previously defined by the When method
func (e *PostAuthMockCreateUserExpectation) Then(rp1 *model.RegisterResp, err error) *PostAuthMock {
	e.results = &PostAuthMockCreateUserResults{rp1, err}
	return e.mock
}

// Times sets number of times PostAuth.CreateUser should be invoked
func (mmCreateUser *mPostAuthMockCreateUser) Times(n uint64) *mPostAuthMockCreateUser {
	if n == 0 {
		mmCreateUser.mock.t.Fatalf("Times of PostAuthMock.CreateUser mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreateUser.expectedInvocations, n)
	mmCreateUser.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreateUser
}

func (mmCreateUser *mPostAuthMockCreateUser) invocationsDone() bool {
	if len(mmCreateUser.expectations) == 0 && mmCreateUser.defaultExpectation == nil && mmCreateUser.mock.funcCreateUser == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreateUser.mock.afterCreateUserCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreateUser.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CreateUser implements PostAuth
func (mmCreateUser *PostAuthMock) CreateUser(ctx context.Context, email string, password string, role string) (rp1 *model.RegisterResp, err error) {
	mm_atomic.AddUint64(&mmCreateUser.beforeCreateUserCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateUser.afterCreateUserCounter, 1)

	mmCreateUser.t.Helper()

	if mmCreateUser.inspectFuncCreateUser != nil {
		mmCreateUser.inspectFuncCreateUser(ctx, email, password, role)
	}

	mm_params := PostAuthMockCreateUserParams{ctx, email, password, role}

	// Record call args
	mmCreateUser.CreateUserMock.mutex.Lock()
	mmCreateUser.CreateUserMock.callArgs = append(mmCreateUser.CreateUserMock.callArgs, &mm_params)
	mmCreateUser.CreateUserMock.mutex.Unlock()

	for _, e := range mmCreateUser.CreateUserMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.rp1, e.results.err
		}
	}

	if mmCreateUser.CreateUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateUser.CreateUserMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateUser.CreateUserMock.defaultExpectation.params
		mm_want_ptrs := mmCreateUser.CreateUserMock.defaultExpectation.paramPtrs

		mm_got := PostAuthMockCreateUserParams{ctx, email, password, role}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreateUser.t.Errorf("PostAuthMock.CreateUser got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateUser.CreateUserMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.email != nil && !minimock.Equal(*mm_want_ptrs.email, mm_got.email) {
				mmCreateUser.t.Errorf("PostAuthMock.CreateUser got unexpected parameter email, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateUser.CreateUserMock.defaultExpectation.expectationOrigins.originEmail, *mm_want_ptrs.email, mm_got.email, minimock.Diff(*mm_want_ptrs.email, mm_got.email))
			}

			if mm_want_ptrs.password != nil && !minimock.Equal(*mm_want_ptrs.password, mm_got.password) {
				mmCreateUser.t.Errorf("PostAuthMock.CreateUser got unexpected parameter password, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateUser.CreateUserMock.defaultExpectation.expectationOrigins.originPassword, *mm_want_ptrs.password, mm_got.password, minimock.Diff(*mm_want_ptrs.password, mm_got.password))
			}

			if mm_want_ptrs.role != nil && !minimock.Equal(*mm_want_ptrs.role, mm_got.role) {
				mmCreateUser.t.Errorf("PostAuthMock.CreateUser got unexpected parameter role, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateUser.CreateUserMock.defaultExpectation.expectationOrigins.originRole, *mm_want_ptrs.role, mm_got.role, minimock.Diff(*mm_want_ptrs.role, mm_got.role))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateUser.t.Errorf("PostAuthMock.CreateUser got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreateUser.CreateUserMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateUser.CreateUserMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateUser.t.Fatal("No results are set for the PostAuthMock.CreateUser")
		}
		return (*mm_results).rp1, (*mm_results).err
	}
	if mmCreateUser.funcCreateUser != nil {
		return mmCreateUser.funcCreateUser(ctx, email, password, role)
	}
	mmCreateUser.t.Fatalf("Unexpected call to PostAuthMock.CreateUser. %v %v %v %v", ctx, email, password, role)
	return
}

// CreateUserAfterCounter returns a count of finished PostAuthMock.CreateUser invocations
func (mmCreateUser *PostAuthMock) CreateUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateUser.afterCreateUserCounter)
}

// CreateUserBeforeCounter returns a count of PostAuthMock.CreateUser invocations
func (mmCreateUser *PostAuthMock) CreateUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateUser.beforeCreateUserCounter)
}

// Calls returns a list of arguments used in each call to PostAuthMock.CreateUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateUser *mPostAuthMockCreateUser) Calls() []*PostAuthMockCreateUserParams {
	mmCreateUser.mutex.RLock()

	argCopy := make([]*PostAuthMockCreateUserParams, len(mmCreateUser.callArgs))
	copy(argCopy, mmCreateUser.callArgs)

	mmCreateUser.mutex.RUnlock()

	return argCopy
}

// MinimockCreateUserDone returns true if the count of the CreateUser invocations corresponds
// the number of defined expectations
func (m *PostAuthMock) MinimockCreateUserDone() bool {
	if m.CreateUserMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateUserMock.invocationsDone()
}

// MinimockCreateUserInspect logs each unmet expectation
func (m *PostAuthMock) MinimockCreateUserInspect() {
	for _, e := range m.CreateUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PostAuthMock.CreateUser at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateUserCounter := mm_atomic.LoadUint64(&m.afterCreateUserCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateUserMock.defaultExpectation != nil && afterCreateUserCounter < 1 {
		if m.CreateUserMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to PostAuthMock.CreateUser at\n%s", m.CreateUserMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to PostAuthMock.CreateUser at\n%s with params: %#v", m.CreateUserMock.defaultExpectation.expectationOrigins.origin, *m.CreateUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateUser != nil && afterCreateUserCounter < 1 {
		m.t.Errorf("Expected call to PostAuthMock.CreateUser at\n%s", m.funcCreateUserOrigin)
	}

	if !m.CreateUserMock.invocationsDone() && afterCreateUserCounter > 0 {
		m.t.Errorf("Expected %d calls to PostAuthMock.CreateUser at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateUserMock.expectedInvocations), m.CreateUserMock.expectedInvocationsOrigin, afterCreateUserCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PostAuthMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateUserInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PostAuthMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *PostAuthMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateUserDone()
}

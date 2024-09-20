// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i chat-server/internal/repository.ChatUserRepository -o chat_user_repository_minimock.go -n ChatUserRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ChatUserRepositoryMock implements repository.ChatUserRepository
type ChatUserRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, chat_id int64, ids []int64) (err error)
	inspectFuncCreate   func(ctx context.Context, chat_id int64, ids []int64)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mChatUserRepositoryMockCreate

	funcDelete          func(ctx context.Context, chat_id int64, user_id int64) (err error)
	inspectFuncDelete   func(ctx context.Context, chat_id int64, user_id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mChatUserRepositoryMockDelete
}

// NewChatUserRepositoryMock returns a mock for repository.ChatUserRepository
func NewChatUserRepositoryMock(t minimock.Tester) *ChatUserRepositoryMock {
	m := &ChatUserRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mChatUserRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*ChatUserRepositoryMockCreateParams{}

	m.DeleteMock = mChatUserRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*ChatUserRepositoryMockDeleteParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mChatUserRepositoryMockCreate struct {
	mock               *ChatUserRepositoryMock
	defaultExpectation *ChatUserRepositoryMockCreateExpectation
	expectations       []*ChatUserRepositoryMockCreateExpectation

	callArgs []*ChatUserRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// ChatUserRepositoryMockCreateExpectation specifies expectation struct of the ChatUserRepository.Create
type ChatUserRepositoryMockCreateExpectation struct {
	mock    *ChatUserRepositoryMock
	params  *ChatUserRepositoryMockCreateParams
	results *ChatUserRepositoryMockCreateResults
	Counter uint64
}

// ChatUserRepositoryMockCreateParams contains parameters of the ChatUserRepository.Create
type ChatUserRepositoryMockCreateParams struct {
	ctx     context.Context
	chat_id int64
	ids     []int64
}

// ChatUserRepositoryMockCreateResults contains results of the ChatUserRepository.Create
type ChatUserRepositoryMockCreateResults struct {
	err error
}

// Expect sets up expected params for ChatUserRepository.Create
func (mmCreate *mChatUserRepositoryMockCreate) Expect(ctx context.Context, chat_id int64, ids []int64) *mChatUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatUserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatUserRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &ChatUserRepositoryMockCreateParams{ctx, chat_id, ids}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the ChatUserRepository.Create
func (mmCreate *mChatUserRepositoryMockCreate) Inspect(f func(ctx context.Context, chat_id int64, ids []int64)) *mChatUserRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for ChatUserRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by ChatUserRepository.Create
func (mmCreate *mChatUserRepositoryMockCreate) Return(err error) *ChatUserRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatUserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatUserRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &ChatUserRepositoryMockCreateResults{err}
	return mmCreate.mock
}

// Set uses given function f to mock the ChatUserRepository.Create method
func (mmCreate *mChatUserRepositoryMockCreate) Set(f func(ctx context.Context, chat_id int64, ids []int64) (err error)) *ChatUserRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the ChatUserRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the ChatUserRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the ChatUserRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mChatUserRepositoryMockCreate) When(ctx context.Context, chat_id int64, ids []int64) *ChatUserRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatUserRepositoryMock.Create mock is already set by Set")
	}

	expectation := &ChatUserRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &ChatUserRepositoryMockCreateParams{ctx, chat_id, ids},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up ChatUserRepository.Create return parameters for the expectation previously defined by the When method
func (e *ChatUserRepositoryMockCreateExpectation) Then(err error) *ChatUserRepositoryMock {
	e.results = &ChatUserRepositoryMockCreateResults{err}
	return e.mock
}

// Create implements repository.ChatUserRepository
func (mmCreate *ChatUserRepositoryMock) Create(ctx context.Context, chat_id int64, ids []int64) (err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, chat_id, ids)
	}

	mm_params := ChatUserRepositoryMockCreateParams{ctx, chat_id, ids}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := ChatUserRepositoryMockCreateParams{ctx, chat_id, ids}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("ChatUserRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the ChatUserRepositoryMock.Create")
		}
		return (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, chat_id, ids)
	}
	mmCreate.t.Fatalf("Unexpected call to ChatUserRepositoryMock.Create. %v %v %v", ctx, chat_id, ids)
	return
}

// CreateAfterCounter returns a count of finished ChatUserRepositoryMock.Create invocations
func (mmCreate *ChatUserRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of ChatUserRepositoryMock.Create invocations
func (mmCreate *ChatUserRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to ChatUserRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mChatUserRepositoryMockCreate) Calls() []*ChatUserRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*ChatUserRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *ChatUserRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *ChatUserRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatUserRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatUserRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to ChatUserRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to ChatUserRepositoryMock.Create")
	}
}

type mChatUserRepositoryMockDelete struct {
	mock               *ChatUserRepositoryMock
	defaultExpectation *ChatUserRepositoryMockDeleteExpectation
	expectations       []*ChatUserRepositoryMockDeleteExpectation

	callArgs []*ChatUserRepositoryMockDeleteParams
	mutex    sync.RWMutex
}

// ChatUserRepositoryMockDeleteExpectation specifies expectation struct of the ChatUserRepository.Delete
type ChatUserRepositoryMockDeleteExpectation struct {
	mock    *ChatUserRepositoryMock
	params  *ChatUserRepositoryMockDeleteParams
	results *ChatUserRepositoryMockDeleteResults
	Counter uint64
}

// ChatUserRepositoryMockDeleteParams contains parameters of the ChatUserRepository.Delete
type ChatUserRepositoryMockDeleteParams struct {
	ctx     context.Context
	chat_id int64
	user_id int64
}

// ChatUserRepositoryMockDeleteResults contains results of the ChatUserRepository.Delete
type ChatUserRepositoryMockDeleteResults struct {
	err error
}

// Expect sets up expected params for ChatUserRepository.Delete
func (mmDelete *mChatUserRepositoryMockDelete) Expect(ctx context.Context, chat_id int64, user_id int64) *mChatUserRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatUserRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatUserRepositoryMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &ChatUserRepositoryMockDeleteParams{ctx, chat_id, user_id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the ChatUserRepository.Delete
func (mmDelete *mChatUserRepositoryMockDelete) Inspect(f func(ctx context.Context, chat_id int64, user_id int64)) *mChatUserRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for ChatUserRepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by ChatUserRepository.Delete
func (mmDelete *mChatUserRepositoryMockDelete) Return(err error) *ChatUserRepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatUserRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatUserRepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &ChatUserRepositoryMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the ChatUserRepository.Delete method
func (mmDelete *mChatUserRepositoryMockDelete) Set(f func(ctx context.Context, chat_id int64, user_id int64) (err error)) *ChatUserRepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the ChatUserRepository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the ChatUserRepository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the ChatUserRepository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mChatUserRepositoryMockDelete) When(ctx context.Context, chat_id int64, user_id int64) *ChatUserRepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatUserRepositoryMock.Delete mock is already set by Set")
	}

	expectation := &ChatUserRepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &ChatUserRepositoryMockDeleteParams{ctx, chat_id, user_id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up ChatUserRepository.Delete return parameters for the expectation previously defined by the When method
func (e *ChatUserRepositoryMockDeleteExpectation) Then(err error) *ChatUserRepositoryMock {
	e.results = &ChatUserRepositoryMockDeleteResults{err}
	return e.mock
}

// Delete implements repository.ChatUserRepository
func (mmDelete *ChatUserRepositoryMock) Delete(ctx context.Context, chat_id int64, user_id int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, chat_id, user_id)
	}

	mm_params := ChatUserRepositoryMockDeleteParams{ctx, chat_id, user_id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, &mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := ChatUserRepositoryMockDeleteParams{ctx, chat_id, user_id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("ChatUserRepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the ChatUserRepositoryMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, chat_id, user_id)
	}
	mmDelete.t.Fatalf("Unexpected call to ChatUserRepositoryMock.Delete. %v %v %v", ctx, chat_id, user_id)
	return
}

// DeleteAfterCounter returns a count of finished ChatUserRepositoryMock.Delete invocations
func (mmDelete *ChatUserRepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of ChatUserRepositoryMock.Delete invocations
func (mmDelete *ChatUserRepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to ChatUserRepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mChatUserRepositoryMockDelete) Calls() []*ChatUserRepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*ChatUserRepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *ChatUserRepositoryMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *ChatUserRepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatUserRepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatUserRepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to ChatUserRepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to ChatUserRepositoryMock.Delete")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatUserRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockDeleteInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatUserRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ChatUserRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone()
}
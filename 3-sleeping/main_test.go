package main

import (
	"sync"
	"testing"
	"time"
)

type publisherMock1 struct {
	mu  sync.RWMutex
	got []Foo
}

func (p *publisherMock1) Publish(got []Foo) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.got = got
}

func (p *publisherMock1) Get() []Foo {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.got
}

func TestGetBestFoo(t *testing.T) {
	mock := publisherMock1{}
	h := Handler{
		publisher: &mock,
		n:         2,
	}

	foo := h.getBestFoo(42)
	// Check foo
	_ = foo

	time.Sleep(1000 * time.Millisecond)
	published := mock.Get()
	if len(published) != 2 {
		t.Fatalf("expected 2, got %d", published)
	}
}

func TestGetBestFoo2(t *testing.T) {
	mock := publisherMock1{}
	h := Handler{
		publisher: &mock,
		n:         2,
	}

	foo := h.getBestFoo(42)
	// Check foo
	_ = foo

	assert(t, func() bool {
		return len(mock.Get()) == 2
	}, 30, time.Millisecond)

	published := mock.Get()
	// Check published
	_ = published
}

func assert(t *testing.T, assertion func() bool, maxRetry int, waitTime time.Duration) {
	for i := 0; i < maxRetry; i++ {
		if assertion() {
			return
		}
		time.Sleep(waitTime)
	}
	t.Fail()
}

type publisherMock2 struct {
	ch chan []Foo
}

func (p *publisherMock2) Publish(got []Foo) {
	p.ch <- got
}

func TestGetBestFoo3(t *testing.T) {
	mock := publisherMock2{
		ch: make(chan []Foo),
	}
	defer close(mock.ch)

	h := Handler{
		publisher: &mock,
		n:         2,
	}
	foo := h.getBestFoo(42)
	// Check foo
	_ = foo

	if v := len(<-mock.ch); v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}
}

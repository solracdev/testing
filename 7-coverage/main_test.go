package main

import "testing"

func TestSum(t *testing.T) {
	var (
		a        = 5
		b        = 10
		expected = 15
	)

	r := sum(a, b)
	if r != expected {
		t.Errorf("sum: expected 15, got %v", r)
	}
}

func TestSub(t *testing.T) {
	var (
		a        = 5
		b        = 4
		expected = 1
	)

	r := sub(a, b)
	if r != expected {
		t.Errorf("sub: expected 1, got %v", r)
	}
}

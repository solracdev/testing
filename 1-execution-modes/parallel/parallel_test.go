package parallel

import "testing"

func TestA(t *testing.T) {
	t.Parallel()
	// ...
}

func TestB(t *testing.T) {
	t.Parallel()
	// ...
}
func TestC(t *testing.T) {
	// ...
}

func TestD(t *testing.T) {
	t.Parallel()
}

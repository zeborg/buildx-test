package calculator

import "testing"

func TestDivide(t *testing.T) {
	if got, want := Divide(4, 2), 2; got != want {
		t.Errorf("multiply method produced wrong result. expected: %d, got: %d ", want, got)
	}
}

package dynamic_programming

import "testing"

func TestFibo(t *testing.T) {
	f := fibo(5)

	if f != 5 {
		t.Errorf("expect 5, but got %v", f)
	}
}

func TestFiboDp(t *testing.T) {
	f := fibo_dp(5)

	if f != 5 {
		t.Errorf("expect 5, but got %v", f)
	}
}

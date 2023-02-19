package matrix

import (
	"reflect"
	"testing"
)

func TestBFT(t *testing.T) {
	data := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	matrix := New(data)
	values := matrix.BreathFirstTraverse()

	if reflect.DeepEqual([]int{'1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '0', '1', '0', '0', '0', '0', '0', '0', '0', '0'}, values) {
		t.Errorf("expected [1 1 1 1 1 1 1 0 1 0 0 1 0 0 0 0 0 0 0 0], but got %v", values)
	}
}

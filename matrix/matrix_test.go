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

	if !reflect.DeepEqual([]byte{'1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '0', '1', '0', '0', '0', '0', '0', '0', '0', '0'}, values) {
		t.Errorf("expected [1 1 1 1 1 1 1 0 1 0 0 1 0 0 0 0 0 0 0 0], but got %v", values)
	}
}

func TestNumsOfIslands(t *testing.T) {
	data := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	matrix := New(data)
	count := matrix.NumsOfIslands('1', '0')

	if count != 1 {
		t.Errorf("expected 1, but got %v", count)
	}

	data2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	matrix2 := New(data2)
	count2 := matrix2.NumsOfIslands('1', '0')
	if count2 != 3 {
		t.Errorf("expected 3, but got %v", count2)
	}
}

func TestDFT(t *testing.T) {
	data := [][]int{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 2},
	}

	matrix := New(data)
	values := matrix.DepthFirstTraverse()

	if !reflect.DeepEqual([]int{0, 0, 0, 1, 2, 2, 1, 1, 2}, values) {
		t.Errorf("expected [0, 0, 0, 1, 2, 2, 1, 1, 2], but got %v", values)
	}
}

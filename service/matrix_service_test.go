package service

import (
	"reflect"
	"testing"
)

func TestEcho(t *testing.T) {
	echoTests := []struct {
		matrix [][]int
		expected string
	}{
		{ [][]int{ {1,2,3}, {4,5,6}, {7,8,9} }, "1,2,3\n4,5,6\n7,8,9\n" },
		{ [][]int{ {1,2}, {3,4} }, "1,2\n3,4\n" },
	}

	ms := NewMatrixService()

	for _, tt := range echoTests {
        got := ms.Echo(tt.matrix)
        if got != tt.expected {
			t.Errorf("echo returned wrong result: got %v expected %v",
            got, tt.expected)
        }
    }
}

func TestInvert(t *testing.T) {
	invertTests := []struct {
		matrix [][]int
		expected string
	}{
		{ [][]int{ {1,2,3}, {4,5,6}, {7,8,9} }, "1,4,7\n2,5,8\n3,6,9\n" },
		{ [][]int{ {1,2}, {3,4} }, "1,3\n2,4\n" },
	}

	ms := NewMatrixService()

	for _, tt := range invertTests {
        got := ms.Invert(tt.matrix)
        if reflect.DeepEqual(got, tt.expected) == false {
			t.Errorf("invert returned wrong result: got %v expected %v",
            got, tt.expected)
        }
    }
}

func TestFlatten(t *testing.T) {
	flattenTests := []struct {
		matrix [][]int
		expected string
	}{
		{ [][]int{ {1,2,3}, {4,5,6}, {7,8,9} }, "1,2,3,4,5,6,7,8,9" },
		{ [][]int{ {1,2}, {3,4} }, "1,2,3,4" },
	}

	ms := NewMatrixService()

	for _, tt := range flattenTests {
        got := ms.Flatten(tt.matrix)
        if got != tt.expected {
			t.Errorf("flatten returned wrong result: got %v expected %v",
            got, tt.expected)
        }
    }
}

func TestSum(t *testing.T) {
	sumTests := []struct {
		matrix [][]int
		expected string
	}{
		{ [][]int{ {1,2,3}, {4,5,6}, {7,8,9} }, "45" },
		{ [][]int{ {1,2}, {3,4} }, "10" },
		{ [][]int{ {0,0}, {0,0} }, "0" },
		{ [][]int{ {0,0}, {1,2,3}, {0,0} }, "6" },
	}

	ms := NewMatrixService()

	for _, tt := range sumTests {
        got := ms.Sum(tt.matrix)
        if got != tt.expected {
			t.Errorf("sum returned wrong result: got %v expected %v",
            got, tt.expected)
        }
    }
}

func TestMultiply(t *testing.T) {
	multiplyTests := []struct {
		matrix [][]int
		expected string
	}{
		{ [][]int{ {1,2,3}, {4,5,6}, {7,8,9} }, "362880" },
		{ [][]int{ {1,2}, {3,4} }, "24" },
		{ [][]int{ {1,2}, {3,4}, {0} }, "0" },
		{ [][]int{ {1,2}, {-3,4} }, "-24" },
	}

	ms := NewMatrixService()

	for _, tt := range multiplyTests {
        got := ms.Multiply(tt.matrix)
        if got != tt.expected {
			t.Errorf("multiply returned wrong result: got %v expected %v",
            got, tt.expected)
        }
    }
}
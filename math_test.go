package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -2, -3},
		{-1, 1, 0},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{3, 2, 1},
		{0, 0, 0},
		{-1, -2, 1},
		{-1, 1, -2},
	}

	for _, tc := range testCases {
		result := Subtract(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Subtract(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 0, 0},
		{-1, -2, 2},
		{-1, 1, -1},
	}

	for _, tc := range testCases {
		result := Multiply(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Multiply(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{6, 3, 2},
		{0, 1, 0},
		{-6, -2, 3},
		{-6, 2, -3},
		{5, 0, 0}, // 测试除数为0的情况
	}

	for _, tc := range testCases {
		result := Divide(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Divide(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestPower(t *testing.T) {
	testCases := []struct {
		base, exponent, expected int
	}{
		{2, 3, 8},
		{3, 2, 9},
		{5, 0, 1},
		{0, 5, 0},
	}

	for _, tc := range testCases {
		result := Power(tc.base, tc.exponent)
		if result != tc.expected {
			t.Errorf("Power(%d, %d) = %d; expected %d", tc.base, tc.exponent, result, tc.expected)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{5, 3, 5},
		{2, 8, 8},
		{0, 0, 0},
		{-1, -5, -1},
	}

	for _, tc := range testCases {
		result := Max(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Max(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

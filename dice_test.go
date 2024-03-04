package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected int
	}{
		{"positive number", 5, 5},
		{"negative number", -5, 5},
		{"zero", 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Abs(test.number)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestIsSuccess(t *testing.T) {
	tests := []struct {
		name        string
		threshold   int
		successNum  int
		combination []int
		expected    bool
	}{
		{"all above threshold", 3, 2, []int{3, 4, 5}, true},
		{"not enough above threshold", 4, 3, []int{3, 4, 5}, false},
		{"exact match", 4, 2, []int{4, 4, 2}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isSuccess(test.threshold, test.successNum, test.combination)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestSummSlices(t *testing.T) {

}

func TestProcessCombination(t *testing.T) {

}

package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, false},
		{3, true},
		{4, false},
		{11, true},
		{12, false},
		{13, true},
		{14, false},
		{21, false},
		{22, false},
		{23, true},
	}

	for _, test := range tests {
		if result := IsPrime(test.input); result != test.expected {
			t.Errorf("IsPrime(%d) = %v; want %v", test.input, result, test.expected)
		}
	}

}

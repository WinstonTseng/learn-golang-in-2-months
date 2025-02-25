package mathutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) return %d, want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Substract(10, 5)
	expected := 5
	if result != expected {
		t.Errorf("Substract(10, 5) return %d, want %d", result, expected)
	}
}

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{5, 5, 10},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) return %d, want %d", test.a, test.b, result, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

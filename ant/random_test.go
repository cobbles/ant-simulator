package ant

import (
	"testing"
)

// MockRandomGenerator is a test double that implements RandomGenerator
type MockRandomGenerator struct {
	values []int
	index  int
}

// NewMockRandomGenerator creates a mock with predefined values
func NewMockRandomGenerator(values []int) *MockRandomGenerator {
	return &MockRandomGenerator{
		values: values,
		index:  0,
	}
}

// Intn returns the next predefined value
func (m *MockRandomGenerator) Intn(n int) int {
	if m.index >= len(m.values) {
		return 0
	}
	val := m.values[m.index]
	m.index++
	return val
}

func TestRealRandomGenerator_Intn(t *testing.T) {
	// Test with deterministic seed
	r := NewRealRandomGeneratorWithSeed(42)

	// Test that Intn returns values in expected range
	for i := 0; i < 100; i++ {
		val := r.Intn(10)
		if val < 0 || val >= 10 {
			t.Errorf("Intn(10) returned %d, expected range [0, 9]", val)
		}
	}
}

func TestRealRandomGenerator_Deterministic(t *testing.T) {
	// Test that same seed produces same sequence
	seed := int64(12345)
	r1 := NewRealRandomGeneratorWithSeed(seed)
	r2 := NewRealRandomGeneratorWithSeed(seed)

	// Generate sequence of values
	values1 := make([]int, 10)
	values2 := make([]int, 10)

	for i := 0; i < 10; i++ {
		values1[i] = r1.Intn(100)
		values2[i] = r2.Intn(100)
	}

	// Compare sequences
	for i := 0; i < 10; i++ {
		if values1[i] != values2[i] {
			t.Errorf("Sequences differ at index %d: %d vs %d", i, values1[i], values2[i])
		}
	}
}

func TestRealRandomGenerator_DifferentSeeds(t *testing.T) {
	// Test that different seeds produce different sequences
	r1 := NewRealRandomGeneratorWithSeed(123)
	r2 := NewRealRandomGeneratorWithSeed(456)

	// Generate first few values
	same := true
	for i := 0; i < 10; i++ {
		if r1.Intn(100) != r2.Intn(100) {
			same = false
			break
		}
	}

	if same {
		t.Error("Different seeds produced same sequence")
	}
}

func TestMockRandomGenerator_Intn(t *testing.T) {
	// Test mock returns predefined values
	expected := []int{1, 2, 3, 4, 5}
	mock := NewMockRandomGenerator(expected)

	for i, val := range expected {
		result := mock.Intn(10) // n parameter should be ignored by mock
		if result != val {
			t.Errorf("Mock returned %d at index %d, expected %d", result, i, val)
		}
	}
}

func TestMockRandomGenerator_Exhausted(t *testing.T) {
	// Test mock behavior when values are exhausted
	mock := NewMockRandomGenerator([]int{42})

	// Get the predefined value
	if mock.Intn(10) != 42 {
		t.Error("Mock should return predefined value")
	}

	// Subsequent calls should return 0
	if mock.Intn(10) != 0 {
		t.Error("Exhausted mock should return 0")
	}
}

func TestRandomGeneratorInterface(t *testing.T) {
	// Test that both implementations satisfy the interface
	var rg RandomGenerator

	// RealRandomGenerator
	rg = NewRealRandomGeneratorWithSeed(42)
	_ = rg.Intn(10) // Should compile and run

	// MockRandomGenerator
	rg = NewMockRandomGenerator([]int{5})
	_ = rg.Intn(10) // Should compile and run
}

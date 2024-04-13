package test_utils

import "testing"

func AssertMapEquals(t *testing.T, expected, actual map[string]int) {
	t.Helper()

	if len(expected) != len(actual) {
		t.Errorf("Expected map length %d, got %d", len(expected), len(actual))
		return
	}

	for key, value1 := range expected {
		value2, ok := actual[key]
		if !ok || value1 != value2 {
			t.Errorf("Value mismatch for key %q: expected %d, got %d", key, value1, value2)
			return
		}
	}
}

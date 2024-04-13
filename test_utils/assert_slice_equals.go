package test_utils

import "testing"

func AssertSlicesEquals(t *testing.T, expected, actual []string) {
	t.Helper()

	if len(expected) != len(actual) {
		t.Errorf("Expected slice length %d, got %d", len(expected), len(actual))
		return
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Value mismatch for index %d: expected %v, got %v", i, expected, actual)
		}
	}
}

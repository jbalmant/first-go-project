package test_utils

import (
	"testing"
)

func AssertEquals(t *testing.T, expected, actual interface{}) {
	t.Helper()

	switch expectedValue := expected.(type) {
	case int:
		actualValue, ok := actual.(int)
		if !ok {
			t.Errorf("Expected int, got %T", actual)
			return
		}
		if expectedValue != actualValue {
			t.Errorf("Expected %d, got %d", expectedValue, actualValue)
		}
	case string:
		actualValue, ok := actual.(string)
		if !ok {
			t.Errorf("Expected string, got %T", actual)
			return
		}
		if expectedValue != actualValue {
			t.Errorf("Expected %q, got %q", expectedValue, actualValue)
		}
	default:
		t.Errorf("unsupported type %T", expected)
	}
}

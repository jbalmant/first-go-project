package test_utils

import (
	"reflect"
	"testing"
)

func AssertObjectEquals(t *testing.T, expected, actual interface{}) {
	t.Helper()

	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Errorf("Expected type %T, got %T", expected, actual)
		return
	}

	if reflect.TypeOf(expected).Kind() == reflect.Ptr {
		expected = reflect.ValueOf(expected).Elem().Interface()
	}

	if reflect.TypeOf(actual).Kind() == reflect.Ptr {
		actual = reflect.ValueOf(actual).Elem().Interface()
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	for i := 0; i < expectedValue.NumField(); i++ {
		fieldName := expectedValue.Type().Field(i).Name
		expectedFieldValue := expectedValue.Field(i)
		actualFieldValue := actualValue.Field(i)

		if !reflect.DeepEqual(expectedFieldValue.Interface(), actualFieldValue.Interface()) {
			t.Errorf("Field %s: expected %v, got %v", fieldName, expectedFieldValue.Interface(), actualFieldValue.Interface())
		}
	}
}

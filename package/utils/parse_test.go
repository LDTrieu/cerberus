package utils

import "testing"

func TestParse(t *testing.T) {

	// Test case 1: Parsing int to int
	var inInt int = 42
	outInt, err := Parse[int](inInt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if inInt != outInt {
		t.Errorf("Expected outInt to be %d, but got %d", inInt, outInt)
	}

	// Test case 2: Parsing float to float
	var inFloat float64 = 3.14
	outFloat, err := Parse[float64](inFloat)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if inFloat != outFloat {
		t.Errorf("Expected outFloat to be %f, but got %f", inFloat, outFloat)
	}

	// Test case 3: Parsing int to float
	var inIntToFloat int = 42
	outFloatFromInt, err := Parse[float64](inIntToFloat)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedFloatFromInt := float64(inIntToFloat)
	if outFloatFromInt != expectedFloatFromInt {
		t.Errorf("Expected outFloatFromInt to be %f, but got %f", expectedFloatFromInt, outFloatFromInt)
	}

	// Test case 4: Parsing float to int
	var inFloatToInt float64 = 3.14
	outIntFromFloat, err := Parse[int](inFloatToInt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedIntFromFloat := int(inFloatToInt)
	if outIntFromFloat != expectedIntFromFloat {
		t.Errorf("Expected outIntFromFloat to be %d, but got %d", expectedIntFromFloat, outIntFromFloat)
	}

	// Test case 5: Parsing string to int
	var inStringToInt string = "42"
	outIntFromString, err := Parse[int](inStringToInt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedIntFromString := 42
	if outIntFromString != expectedIntFromString {
		t.Errorf("Expected outIntFromString to be %d, but got %d", expectedIntFromString, outIntFromString)
	}

	// Test case 6: Parsing string to float
	var inStringToFloat string = "3.14"
	outFloatFromString, err := Parse[float64](inStringToFloat)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedFloatFromString := 3.14
	if outFloatFromString != expectedFloatFromString {
		t.Errorf("Expected outFloatFromString to be %f, but got %f", expectedFloatFromString, outFloatFromString)
	}

	// Test case 7: Parsing int to string
	var inIntToString int = 42
	outStringFromInt, err := Parse[string](inIntToString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedStringFromInt := "42"
	if outStringFromInt != expectedStringFromInt {
		t.Errorf("Expected outStringFromInt to be %s, but got %s", expectedStringFromInt, outStringFromInt)
	}

	// Test case 8: Parsing float to string
	var inFloatToString float64 = 3.14
	outStringFromFloat, err := Parse[string](inFloatToString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedStringFromFloat := "3.14"
	if outStringFromFloat != expectedStringFromFloat {
		t.Errorf("Expected outStringFromFloat to be %s, but got %s", expectedStringFromFloat, outStringFromFloat)
	}

	// Test case 9: Parsing string to string
	var inStringToString string = "hello"
	outStringFromString, err := Parse[string](inStringToString)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedStringFromString := "hello"
	if outStringFromString != expectedStringFromString {
		t.Errorf("Expected outStringFromString to be %s, but got %s", expectedStringFromString, outStringFromString)
	}

	// Test case 10: Parsing int64 to int64
	var inInt64 int64 = 1234567890
	outInt64, err := Parse[int64](inInt64)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if inInt64 != outInt64 {
		t.Errorf("Expected outInt64 to be %d, but got %d", inInt64, outInt64)
	}

	// Test case 11: Parsing float64 to float64
	var inFloat64 float64 = 3.14159
	outFloat64, err := Parse[float64](inFloat64)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if inFloat64 != outFloat64 {
		t.Errorf("Expected outFloat64 to be %f, but got %f", inFloat64, outFloat64)
	}

	// Test case 12: Parsing string to int64
	var inStringToInt64 string = "9876543210"
	outInt64FromString, err := Parse[int64](inStringToInt64)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedInt64FromString := int64(9876543210)
	if outInt64FromString != expectedInt64FromString {
		t.Errorf("Expected outInt64FromString to be %d, but got %d", expectedInt64FromString, outInt64FromString)
	}

	// Test case 13: Parsing string to float64
	var inStringToFloat64 string = "3.14159"
	outFloat64FromString, err := Parse[float64](inStringToFloat64)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedFloat64FromString := 3.14159
	if outFloat64FromString != expectedFloat64FromString {
		t.Errorf("Expected outFloat64FromString to be %f, but got %f", expectedFloat64FromString, outFloat64FromString)
	}

	// Test case 14: Parsing int64 to float64
	var inInt64ToFloat64 int64 = 1234567890
	outFloat64FromInt64, err := Parse[float64](inInt64ToFloat64)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedFloat64FromInt64 := float64(inInt64ToFloat64)
	if outFloat64FromInt64 != expectedFloat64FromInt64 {
		t.Errorf("Expected outFloat64FromInt64 to be %f, but got %f", expectedFloat64FromInt64, outFloat64FromInt64)
	}

	// Test case 15: Parsing float64 to int64
	var inFloat64ToInt64 float64 = 3.14159
	outInt64FromFloat64, err := Parse[int64](inFloat64ToInt64)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedInt64FromFloat64 := int64(inFloat64ToInt64)
	if outInt64FromFloat64 != expectedInt64FromFloat64 {
		t.Errorf("Expected outInt64FromFloat64 to be %d, but got %d", expectedInt64FromFloat64, outInt64FromFloat64)
	}
}

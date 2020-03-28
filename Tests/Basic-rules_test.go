package Tests

import (
	"fmt"
	"pwd-safety/Utils"
	"strconv"
	"testing"
)

//Test if two slices of bytes are equal
func testEqByte(a, b []byte) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//Test the correct operation of GenerateSetString func
func TestGenerateSetString(t *testing.T) {
	var tests = []struct {
		input    string
		expected []byte
	}{
		{"hellohello", []byte{'h', 'e', 'l', 'o'}},
		{"", []byte{}},
		{"abcdefghijklmnopqrstuvwxyz", []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}},
		{"blablablablabla", []byte{'b', 'l', 'a'}},
		{".,:;-_à#ç@", []byte{'.', ',', ':', ';', '-', '_', 'à', '#', 'ç', '@'}},
	}

	for _, test := range tests {
		if output := Utils.GenerateSetString(test.input); !testEqByte(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %d expected, received: %d", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

//Test the correct operation of HowManyDifferents func
func TestHowManyDifferents(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"hellohello", 4},
		{"", 0},
		{"abcdefghijklmnopqrstuvwxyz", 26},
		{"blablablablabla", 3},
		{".,:;-_à#ç@", 10},
	}

	for _, test := range tests {
		if output := Utils.HowManyDifferents(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %d expected, received: %d", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

//Test the correct operation of IsThereUpperCase func
func TestIsThereUpperCase(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"hellohello", false},
		{"", false},
		{"abcdefghijklmnopqrstuvwxyz", false},
		{"blablaBlablablB", true},
		{".,:;-A_à#ç@", true},
	}

	for _, test := range tests {
		if output := Utils.IsThereUpperCase(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %s expected, received: %s", test.input, strconv.FormatBool(test.expected), strconv.FormatBool(output))
			t.Error(errorString)
		}
	}
}

//Test the correct operation of IsThereLowerCase func
func TestIsThereLowerCase(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"hellohello", true},
		{"", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", false},
		{"blablaBlablablB", true},
		{".,:;-A_à#ç@", false},
	}

	for _, test := range tests {
		if output := Utils.IsThereLowerCase(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %s expected, received: %s", test.input, strconv.FormatBool(test.expected), strconv.FormatBool(output))
			t.Error(errorString)
		}
	}
}

//Test the correct operation of IsThereSymbol func
func TestIsThereSymbol(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"hellohello", false},
		{"", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", false},
		{"blablaBlabl,ablB", true},
		{".,:;-A_à#ç@", true},
	}

	for _, test := range tests {
		if output := Utils.IsThereSymbol(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %s expected, received: %s", test.input, strconv.FormatBool(test.expected), strconv.FormatBool(output))
			t.Error(errorString)
		}
	}
}

//Test the correct operation of IsThereNumber func
func TestIsThereNumber(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"hellohello", false},
		{"", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", false},
		{"blabla0Blabl,ablB", true},
		{".,:;-A_à#ç@", false},
	}

	for _, test := range tests {
		if output := Utils.IsThereNumber(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %s expected, received: %s", test.input, strconv.FormatBool(test.expected), strconv.FormatBool(output))
			t.Error(errorString)
		}
	}
}

//Test the correct operation of HowManyTypes func
func TestHowManyTypes(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"hellohello", 1},
		{"", 0},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1},
		{"blabla0Blabl,ablB", 4},
		{".,:;-A1_à#ç@", 3},
	}

	for _, test := range tests {
		if output := Utils.HowManyTypes(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %d expected, received: %d", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

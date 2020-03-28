package Tests

import (
	"fmt"
	"pwd-safety/Utils"
	"strconv"
	"testing"
)

//Test the correct operation of Reverse func
func TestReverse(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"hellohello", "olleholleh"},
		{"", ""},
		{"ABCDEFGHJKILMNOPQRSTUVWXYZ", "ZYXWVUTSRQPONMLIKJHGFEDCBA"},
		{"blabla0Blabl,ablB", "Blba,lbalB0albalb"},
		{".,:;-A1_à#ç@", "@ç#à_1A-;:,."},
	}

	for _, test := range tests {
		if output := Utils.Reverse(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %s expected, received: %s", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

//Test the correct operation of FindExactly func
func TestFindExactly(t *testing.T) {
	var tests = []struct {
		words    []string
		password string
		expected bool
	}{
		{[]string{"pwd", "password", "hello", "password"}, "password", true},
		{[]string{"pwd", "password", "hello", "password"}, "pass", false},
		{[]string{""}, "", true},
		{[]string{"iwfcaheuhwilehli38ry2r8RYBW8RYBYELYT8AOVWYÒV8LOylR-ù-,à.èù-àq.vàfqè+"}, "iwfcaheuhwilehli38ry2r8RYBW8RYBYELYT8AOVWYÒV8LOylR-ù-,à.èù-àq.vàfqè+", true},
		{[]string{"A", "B", "C", "._-"}, "._-", true},
	}

	for _, test := range tests {
		if output := Utils.FindExactly(test.words, test.password); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %v words inputted, %s password inputted, %s expected, received: %s", test.words, test.password, strconv.FormatBool(test.expected), strconv.FormatBool(output))
			t.Error(errorString)
		}
	}
}

//Test the correct operation of FindExactlyReversed func
func TestFindExactlyReversed(t *testing.T) {
	var tests = []struct {
		words    []string
		password string
		expected bool
	}{
		{[]string{"pwd", "password", "hello", "password"}, "drowssap", true},
		{[]string{"pwd", "password", "hello", "password"}, "pass", false},
		{[]string{""}, "", true},
		{[]string{"iwfcaheuhwilehli38ry2r8RYBW8RYBYELYT8AOVWYÒV8LOylR-ù-,à.èù-àq.vàfqè+"}, "àq.vàfqè+", false},
		{[]string{"A", "B", "C", "._-"}, "-_.", true},
	}

	for _, test := range tests {
		if output := Utils.FindExactlyReversed(test.words, test.password); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %v words inputted, %s password inputted, %s expected, received: %s", test.words, test.password, strconv.FormatBool(test.expected), strconv.FormatBool(output))
			t.Error(errorString)
		}
	}
}

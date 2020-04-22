package tests

import (
	"fmt"
	"pwd-safety/utils"
	"testing"
)

//Test if two slices of strings are equal
func testEqString(a, b []string) bool {

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

//Test the correct operation of ReadWords func
func TestReadWords(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{"TestFiles/Pwd1.txt", []string{"pass", "passwor", "password", "password", "passwrdpassword", "passsssssssword"}},
		{"TestFiles/Pwd2.txt", []string{"main", "main", "main", "mainpassmain"}},
		{"TestFiles/Pwd3.txt", nil},
		{"TestFiles/Pwd4.txt", []string{"iwfcaheuhwilehli38ry2r8RYBW8RYBYELYT8AOVWYÒV8LOylR-ù-,à.èù-àq.vàfqè+"}},
		{"TestFiles/Pwd5.txt", []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}},
	}

	for _, test := range tests {
		if output := utils.ReadWords(test.input); !testEqString(test.expected, output) {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

/*
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Author:
 *      Edoardo Ottavianelli <edoardott@gmail.com>
 */

package scoring

import (
	"fmt"
	"strconv"
	"testing"
)

//Test if two slices of bytes are equal
func testEqRune(a, b []rune) bool {

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
		expected []rune
	}{
		{"hellohello", []rune{'h', 'e', 'l', 'o'}},
		{"", []rune{}},
		{"abcdefghijklmnopqrstuvwxyz", []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}},
		{"blablablablabla", []rune{'b', 'l', 'a'}},
		{" !#$%&()*", []rune{' ', '!', '#', '$', '%', '&', '(', ')', '*'}},
		{"+,-./:;<?@[]^", []rune{'+', ',', '-', '.', '/', ':', ';', '<', '?', '@', '[', ']', '^'}},
		{"=>`{|}~_", []rune{'=', '>', '`', '{', '|', '}', '~', '_'}},
	}

	for _, test := range tests {
		if output := GenerateSetString(test.input); !testEqRune(test.expected, output) {
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
		if output := HowManyDifferents(test.input); output != test.expected {
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
		if output := IsThereUpperCase(test.input); output != test.expected {
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
		if output := IsThereLowerCase(test.input); output != test.expected {
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
		if output := IsThereSymbol(test.input); output != test.expected {
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
		if output := IsThereNumber(test.input); output != test.expected {
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
		if output := HowManyTypes(test.input); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %d expected, received: %d", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

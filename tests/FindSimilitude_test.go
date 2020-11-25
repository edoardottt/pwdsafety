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

package tests

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/edoardottt/pwdsafety/utils"
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
		if output := utils.Reverse(test.input); output != test.expected {
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
		if output := utils.FindExactly(test.words, test.password); output != test.expected {
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
		if output := utils.FindExactlyReversed(test.words, test.password); output != test.expected {
			errorString := fmt.Sprintf("Test Failed: %v words inputted, %s password inputted, %s expected, received: %s", test.words, test.password, strconv.FormatBool(test.expected), strconv.FormatBool(output))
			t.Error(errorString)
		}
	}
}

/*
 *    Copyright (C) 2020 Edoardo Ottavianelli
 *
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
	"github.com/edoardottt/pwd-safety/utils"
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

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
	"github.com/edoardottt/pwdsafety/pwds/utilspwds"
	"github.com/edoardottt/pwdsafety/tests/utilstests"
	"testing"
)

//Test the correct operation of Reverse func
func TestGenerateSetFiles(t *testing.T) {
	var tests = []struct {
		input    []string
		expected []string
	}{
		{[]string{"hellohello"}, []string{"hellohello"}},
		{[]string{}, []string{}},
		{[]string{"ABC","ABC","DEF","GHJ","KIL","MNO","MNO","PQR","STU","VWX","YZ"}, []string{"ABC","DEF","GHJ","KIL","MNO","PQR","STU","VWX","YZ"}},
		{[]string{"1","2","2","2","2","2","3"}, []string{"1","2","3"}},
		{[]string{".",":",":",":",";","'","?","?","!","%","="}, []string{".",":",";","'","?","!","%","="}},
	}

	for _, test := range tests {
		if output := utilspwds.GenerateSetFiles(test.input); !utilstests.TestEqString(output, test.expected) {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %s expected, received: %s", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

//Test the correct operation of CheckDuplicatePwd func
func TestCheckDuplicatePwd(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"testFiles/Pwd1.txt", false},
		{"testFiles/Pwd2.txt", false},
		{"testFiles/Pwd3.txt", true},
		{"testFiles/Pwd4.txt", true},
		{"testFiles/Pwd5.txt", true},
	}

	for _, test := range tests {
		if output := utilspwds.CheckDuplicatePwd(test.input); ! output == test.expected {
			errorString := fmt.Sprintf("Test Failed: %s inputted, %v expected, received: %v", test.input, test.expected, output)
			t.Error(errorString)
		}
	}
}

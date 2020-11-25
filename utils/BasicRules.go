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

package utils

import (
	"unicode/utf8"
)

//GenerateSetString : Generates a set of unique characters in the input string
func GenerateSetString(input string) []rune {
	set := map[rune]bool{}
	result := []rune{}
	for len(input) > 0 {
		char, size := utf8.DecodeRuneInString(input)
		existence := set[char]
		if !existence {
			set[char] = true // add element
			result = append(result, char)
		}
		input = input[size:]
	}
	return result
}

//HowManyDifferents : Returns the number of differents characters used in password
func HowManyDifferents(password string) int {
	return len(GenerateSetString(password))
}

//IsThereUpperCase : Checks if there is at least one UPPERCASE character
func IsThereUpperCase(password string) bool {
	for _, r := range password {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

//IsThereLowerCase : Checks if there is at least one lowercase character
func IsThereLowerCase(password string) bool {
	for _, r := range password {
		if r >= 'a' && r <= 'z' {
			return true
		}
	}
	return false
}

//IsThereSymbol : Checks if there is at least one symbol
func IsThereSymbol(password string) bool {
	for _, r := range password {
		if (r < 'A' || r > 'z') && (r < '0' || r > '9') {
			return true
		}
	}
	return false
}

//IsThereNumber : Checks if there is at least one number
func IsThereNumber(password string) bool {
	for _, r := range password {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

//HowManyTypes : Returns how many different types there are in the password
func HowManyTypes(password string) int {
	var howMany int
	if IsThereNumber(password) {
		howMany++
	}
	if IsThereUpperCase(password) {
		howMany++
	}
	if IsThereLowerCase(password) {
		howMany++
	}
	if IsThereSymbol(password) {
		howMany++
	}
	return howMany
}

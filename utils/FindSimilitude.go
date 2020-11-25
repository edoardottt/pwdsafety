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

//FindExactly : Check if the password is equal to one known password
func FindExactly(words []string, password string) bool {
	for _, s := range words {
		if password == s {
			return true
		}
	}
	return false
}

//Reverse : Reverse the input string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//FindExactlyReversed : Check if the password is equal to one known password reversed
func FindExactlyReversed(words []string, password string) bool {
	for _, s := range words {
		if Reverse(password) == s {
			return true
		}
	}
	return false
}

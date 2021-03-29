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

package main

import (
	"math/rand"
	"time"
)

//GenerateRandom : Generate a strong random password
func GenerateRandom(length int) string {
	var lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
	var upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var specialCharSet = "!@#$%&*-_+-=)("
	var numberSet = "0123456789"
	rand.Seed(time.Now().UnixNano())
	var randomPwd string = ""

	for i := 0; i < length; i++ {

		switch choice := rand.Intn(4); choice {
		case 0:
			leng := len(lowerCharSet)
			index := rand.Intn(leng)
			elem := lowerCharSet[index]
			randomPwd += string(elem)
		case 1:
			leng := len(upperCharSet)
			index := rand.Intn(leng)
			elem := upperCharSet[index]
			randomPwd += string(elem)
		case 2:
			leng := len(specialCharSet)
			index := rand.Intn(leng)
			elem := specialCharSet[index]
			randomPwd += string(elem)
		case 3:
			leng := len(numberSet)
			index := rand.Intn(leng)
			elem := numberSet[index]
			randomPwd += string(elem)
		default:
			leng := len(lowerCharSet)
			index := rand.Intn(leng)
			elem := lowerCharSet[index]
			randomPwd += string(elem)
		}
	}

	return randomPwd
}

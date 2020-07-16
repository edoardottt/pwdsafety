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

package utils

import (
	"bufio"
	"fmt"
	"os"
)

//ReadInput : Reading all inputs from stdin
func ReadInput() []string {
	var result []string
	inputs := [5]string{"name", "surname", "birthday(ddmmyyyy)", "telephone number", "pet's name"}
	for _, value := range inputs {
		result = append(result, ReadSingleInput(value))
	}
	return result
}

//ReadSingleInput : Reading one single input
func ReadSingleInput(input string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[>] Enter " + input + ": ")
	text, _ := reader.ReadString('\n')
	ind := len(text)
	if ind > 0 && text[ind-1] == '\n' {
		text = text[:ind-1]
	}
	return text
}

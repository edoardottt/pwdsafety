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

package utilspwds

import (
	"github.com/edoardottt/pwdsafety/utils"
)

//GenerateSetFiles : It generates a set of the array of strings from input
func GenerateSetFiles(input []string) []string {
	set := map[string]bool{}
	result := []string{}
	for i:= 0; i < len(input);i++ {
		word := input[i]
		existence := set[word]
		if !existence {
			set[word] = true // add element
			result = append(result, word)
		}
	}
	return result
}

//CheckDuplicatePwd : It checks for duplicate passwords in passwords files.
func CheckDuplicatePwd(fileName string) bool {
	words := utils.ReadWords(fileName)
	lengthWords := len(words)
	setWords := GenerateSetFiles(words)
	lengthSet := len(setWords)
	return lengthWords == lengthSet
}
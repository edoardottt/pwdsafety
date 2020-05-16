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
	"strconv"
	"os"
)

//SplitPwds : It splits a single txt file that contains one password per line into n files.
func SplitPwds() {
	file := "known-pwd.txt"
	input := utils.ReadSingleInput("the number of files")
	files, err := strconv.Atoi(input)
	if err != nil {
		print("Insert an integer value.")
		os.Exit(1)
	}
	words := utils.ReadWords(file)
	lengthWords := len(words)
	if 1 < files < lengthWords {

	}
}
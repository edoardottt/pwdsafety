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
	"bufio"
	"fmt"
	"github.com/edoardottt/pwdsafety/utils"
	"os"
	"strconv"
)

//WriteLines : it writes the lines to the given file.
func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

//SplitPwds : It splits a single txt file that contains one password per line into n files.
func SplitPwds(file string) {
	input := utils.ReadSingleInput("the number of files")
	files, err := strconv.Atoi(input)
	if err != nil {
		println("Insert an integer value.")
		os.Exit(1)
	}
	words := utils.ReadWords(file)
	lengthWords := len(words)
	if lengthWords == 0 {
		println("Empty file.")
		os.Exit(1)
	}
	if files <= 1 || files >= lengthWords {
		println("Insert a valid value.")
		os.Exit(1)
	}
	LinesPerFile := lengthWords / files
	baseName := "knownPwd"
	index := 0
	for i := 1; i <= files; i++ {
		fileName := baseName + strconv.Itoa(i) + ".txt"
		end := index + LinesPerFile
		if end <= lengthWords {
			err = WriteLines(words[index:end], fileName)
		} else {
			err = WriteLines(words[index:], fileName)
		}
		index = index + LinesPerFile
	}
}

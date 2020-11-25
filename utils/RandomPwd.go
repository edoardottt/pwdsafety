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
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//GenerateRandom : Generate a strong random password
func GenerateRandom(engWords []string) string {
	rand.Seed(time.Now().UnixNano())
	max := len(engWords) - 1
	rand1 := rand.Intn(max)
	rand2 := rand.Intn(max)
	word1 := engWords[rand1]
	word2 := engWords[rand2]
	rand4 := rand.Intn(100)
	if math.Mod(float64(rand1), 2) == 0 {
		word2 = strings.ToUpper(word2)
	} else {
		word1 = strings.ToUpper(word1)
	}
	symbols := []string{"-", ".", "_", "#", "$", "&"}
	rand3 := math.Mod(float64(rand1+rand2), float64(len(symbols)))
	newPwd := word1 + symbols[int(rand3)] + word2 + symbols[int(rand3)] + strconv.Itoa(rand4)
	return newPwd
}

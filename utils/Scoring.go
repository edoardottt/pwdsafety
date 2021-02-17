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
	"fmt"
	"os"
)

/*
LengthScore :
Scores password's length
	total = 30
	length<=7 = 0
	length==8 = 8
	length==9 = 12
	10 <= length <= 15 = 20
	16 <= length <= 19 = 25
	20 <= length <= 24 = 28
	length >=25 = 30
*/
func LengthScore(password string) float64 {
	length := len(password)
	if length <= 7 {
		return 0
	}
	if length == 8 {
		return 7
	}
	if length == 9 {
		return 12
	}
	if length > 9 && length < 16 {
		return 20
	}
	if length > 15 && length < 20 {
		return 25
	}
	if length > 19 && length < 25 {
		return 28
	}
	return 30
}

/*
CompositionPwdScore :
Scores password's composition
	total = 20
	There is numbers = 5
	There is symbol = 5
	There is uppercase = 5
	There is lowercase = 5
*/
func CompositionPwdScore(password string) float64 {
	var result int
	numbers := IsThereNumber(password)
	upper := IsThereUpperCase(password)
	lower := IsThereLowerCase(password)
	symbols := IsThereSymbol(password)
	if numbers {
		result += 5
	}
	if upper {
		result += 5
	}
	if lower {
		result += 5
	}
	if symbols {
		result += 5
	}
	return float64(result)
}

/*
DifferentCharScore :
Scores How many different chars in relation to the length
	total = 15
	n = (different_chars*total)/total_chars
*/
func DifferentCharScore(password string) float64 {
	diffChars := HowManyDifferents(password)
	total := len(password)
	if total == 0 {
		return 0
	}
	return float64((diffChars * 15) / total)
}

/*
EntropyScore :
Scores Entropy's password
	total = 35
	< 28 bits = 3
	28 - 35 bits = 8
	36 - 59 bits = 20
	60 - 80 bits = 24
	81 - 120 bits = 28
	120+ bits = 35
*/
func EntropyScore(password string) float64 {
	entropy := Entropy(password)
	entropyScore := (entropy * 35) / 130
	if entropyScore > 35 {
		return 35
	}
	return entropyScore
}

/*
pwnedPwds returns the scores for pwned password and
reversed password.

Found : -8
Not Found : 0

*/
func pwnedPwds(password string) (float64, float64) {
	var scoreKnownPwd float64
	var scoreKnownPwdReverse float64
	var knownPwd *Result
	var knownPwdReverse *Result
	//check if password is into known leaked passwords

	knownPwd, err := IsPwned(password)
	if err != nil {
		fmt.Println("Error while retrieving data on password...")
		os.Exit(1)
	}

	knownPwdReverse, err = IsPwned(Reverse(password))
	if err != nil {
		fmt.Println("Error while retrieving data on password...")
		os.Exit(1)
	}

	if knownPwd.Pwned {
		scoreKnownPwd = -8
	} else {
		scoreKnownPwd = 0
	}

	if knownPwdReverse.Pwned {
		scoreKnownPwdReverse = -8
	} else {
		scoreKnownPwdReverse = 0
	}
	return scoreKnownPwd, scoreKnownPwdReverse
}

//Grader : Return the score of the password
func Grader(words [][]string, password string) float64 {
	var optimalLength = 27
	var optimalDifferentCharScore float64 = 7
	var knownStr string
	var knownStrReverse string
	var scoreKnownPwd float64
	var scoreKnownPwdReverse float64

	scoreKnownPwd, scoreKnownPwdReverse = pwnedPwds(password)
	lengthScore := LengthScore(password)
	compositionPwdScore := CompositionPwdScore(password)
	differentCharScore := DifferentCharScore(password)
	entropyScore := EntropyScore(password)
	//Printing results
	if scoreKnownPwd != 0 {
		knownStr = "Yes"
	} else {
		knownStr = "No"
	}
	if scoreKnownPwdReverse != 0 {
		knownStrReverse = "Yes"
	} else {
		knownStrReverse = "No"
	}
	fmt.Println("[%] Password found in known leaked passwords: " + knownStr)
	fmt.Println("[%] Password (reversed) found in known leaked passwords: " + knownStrReverse)
	fmt.Println("[%] Length Score: " + fmt.Sprint(lengthScore) + "/30")
	fmt.Println("[%] Composition Score: " + fmt.Sprint(compositionPwdScore) + "/20")
	fmt.Println("[%] Unique chars Score: " + fmt.Sprint(differentCharScore) + "/15")

	entropyRounded := Round(fmt.Sprintf("%.2f", entropyScore))
	fmt.Println("[%] Entropy Score: " + fmt.Sprint(entropyRounded) + "/35")
	score := scoreKnownPwd + scoreKnownPwdReverse + lengthScore + compositionPwdScore + differentCharScore + entropyScore
	//if it's an optimal password by very high length and good different/unique ratio score
	if differentCharScore >= optimalDifferentCharScore && len(password) > optimalLength {
		return 100
	}
	if score > 0 {
		return score
	}
	return 0
}

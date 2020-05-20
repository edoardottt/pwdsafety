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

/*
KnownPwdScore :
Scores known password
	total = -18
	It's present = -18
	There isn't = 0
*/
func KnownPwdScore(words []string, password string) float64 {
	found := FindExactly(words, password)
	if found {
		return -18
	}
	return 0
}

/*
KnownPwdReverseScore :
Scores known reversed password
	total = -8
	It's present = -8
	There isn't = 0
*/
func KnownPwdReverseScore(words []string, password string) float64 {
	foundReverse := FindExactlyReversed(words, password)
	if foundReverse {
		return -8
	}
	return 0
}

/*
LengthScore :
Scores password's length
	total = 30
	length<=7 = 0
	length==8 = 8
	length==9 = 10
	10 <= length <= 12 = 15
	13 <= length <= 15 = 18
	16 <= length <= 19 = 23
	20 <= length <= 24 = 26
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
		return 10
	}
	if length > 9 && length < 13 {
		return 15
	}
	if length > 12 && length < 16 {
		return 18
	}
	if length > 15 && length < 20 {
		return 23
	}
	if length > 19 && length < 25 {
		return 26
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

//Grader : Return the score of the password
func Grader(words [][]string, password string) float64 {
	var optimalLength = 27
	var optimalDifferentCharScore float64 = 7
	var knownPwd float64
	for i := range words {
		knownPwd1 := KnownPwdScore(words[i], password)
		knownPwd = knownPwd1
		if knownPwd < 0 {
			break
		}
	}
	var knownPwdReverse float64
	for i := range words {
		knownPwdReverse1 := KnownPwdReverseScore(words[i], password)
		knownPwdReverse = knownPwdReverse1
		if knownPwdReverse < 0 {
			break
		}
	}
	lengthScore := LengthScore(password)
	compositionPwdScore := CompositionPwdScore(password)
	differentCharScore := DifferentCharScore(password)
	entropyScore := EntropyScore(password)
	score := knownPwd + knownPwdReverse + lengthScore + compositionPwdScore + differentCharScore + entropyScore
	//if it's an optimal password by very high length and good different/unique ratio score
	if differentCharScore >= optimalDifferentCharScore && len(password) > optimalLength {
		return 100
	}
	if score > 0 {
		return score
	}
	return 0
}

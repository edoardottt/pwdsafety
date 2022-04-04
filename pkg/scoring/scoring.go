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

package scoring

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

//Entropy computes the entropy of a password
func Entropy(password string) float64 {
	var E float64
	var pool float64 = 95
	length := float64(len(password))
	E = -(math.Log2(1 / (math.Pow(pool, length))))
	return E
}

//CountTypeElements : Counts the different types in password
func CountTypeElements(input string) map[string]float64 {
	res := map[string]float64{"lower": 0, "number": 0, "symbol": 0, "upper": 0}
	for i := 0; i < len(input); i++ {
		r := input[i]
		if r >= 'A' && r <= 'Z' { //IF UPPERCASE
			res["upper"]++
		}
		if r >= 'a' && r <= 'z' { //if lowercase
			res["lower"]++
		}
		if (r < 'A' || r > 'z') && (r < '0' || r > '9') { //if numb3r
			res["symbol"]++
		}
		if r >= '0' && r <= '9' { //if symbol
			res["number"]++
		}
	}
	return res
}

//Round : Round in a clever way float64 numbers
func Round(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}
	if math.Mod(value, 1) == 0 {
		v, err := strconv.ParseFloat(fmt.Sprintf("%.0f", value), 64)
		if err != nil {
			panic(err)
		}
		value = v
	}
	return value
}

//CrackTime : Returns the seconds needed to crack the password
func CrackTime(password string) float64 {
	const GPU float64 = 1000000000
	var bots float64 = 15000
	var KPS = bots * GPU
	var combinations float64
	length := float64(len(password))
	var pool float64 = 0
	if IsThereLowerCase(password) {
		pool += 26
	}
	if IsThereUpperCase(password) {
		pool += 26
	}
	if IsThereNumber(password) {
		pool += 10
	}
	if IsThereSymbol(password) {
		pool += 33
	}
	combinations = math.Pow(pool, length)
	return combinations / KPS
}

//ShowCrackTime : Beautify the crack time (from seconds to human readable string )
func ShowCrackTime(crackTime float64) string {
	if crackTime <= 1 {
		return "less than one second."
	}
	var Result string
	var remainder float64
	var seconds float64
	var minutes float64
	var hours float64
	var days float64
	var months float64
	var years float64
	var decades float64
	var centuries float64
	var secondsStr string
	var minutesStr string
	var hoursStr string
	var daysStr string
	var monthsStr string
	var yearsStr string
	var decadesStr string
	var centuriesStr string
	centuries = crackTime / 3110400000
	remainder = math.Mod(crackTime, 3110400000)
	decades = remainder / 311040000
	remainder = math.Mod(remainder, 311040000)
	years = remainder / 31104000
	remainder = math.Mod(remainder, 31104000)
	months = remainder / 2592000
	remainder = math.Mod(remainder, 2592000)
	days = remainder / 86400
	remainder = math.Mod(remainder, 86400)
	hours = remainder / 3600
	remainder = math.Mod(remainder, 3600)
	minutes = remainder / 60
	seconds = math.Mod(remainder, 60)
	if centuries > 1 {
		centuriesStr = strconv.Itoa(int(centuries)) + " Centuries, "
		Result += centuriesStr
	}
	if decades > 1 {
		decadesStr = strconv.Itoa(int(decades)) + " Decades, "
		Result += decadesStr
	}
	if years > 1 {
		yearsStr = strconv.Itoa(int(years)) + " Years, "
		Result += yearsStr
	}
	if months > 1 {
		monthsStr = strconv.Itoa(int(months)) + " Months, "
		Result += monthsStr
	}
	if days > 1 {
		daysStr = strconv.Itoa(int(days)) + " Days, "
		Result += daysStr
	}
	if hours > 1 {
		hoursStr = strconv.Itoa(int(hours)) + " Hours, "
		Result += hoursStr
	}
	if minutes > 1 {
		minutesStr = strconv.Itoa(int(minutes)) + " Minutes, "
		Result += minutesStr
	}
	if seconds > 1 {
		secondsStr = strconv.Itoa(int(seconds)) + " Seconds, "
		Result += secondsStr
	}
	runes := []rune(Result)
	Result = string(runes[0 : len(Result)-2])
	return Result + "."
}

// Result describes a result from the Pwned Password service.
type Result struct {
	// Pwned has the password been seen at least once. A value of false doesn't mean the password is any good though.
	Pwned bool
	// TimesObserved the number of times this password has been seen by the pwned password service.
	TimesObserved uint64
}

type pwnedHash struct {
	Hash  string
	Range string
}

// IsPwned will synchronously check if the provided password has been pwned.
func IsPwned(password string) (*Result, error) {
	if password == "" {
		return nil, fmt.Errorf("empty password provided")
	}

	hash, err := getHash(password)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + hash.Range)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(body), "\r\n")
	for _, line := range lines {
		components := strings.Split(line, ":")
		if len(components) != 2 {
			return nil, fmt.Errorf("invalid response from pwned password API")
		}

		resultHash := components[0]
		countStr := components[1]

		if hash.Range+resultHash == hash.Hash {
			count, err := strconv.ParseUint(countStr, 10, 64)
			if err != nil {
				return nil, err
			}

			ret := Result{
				Pwned:         true,
				TimesObserved: count,
			}
			return &ret, nil
		}
	}

	ret := Result{
		Pwned:         false,
		TimesObserved: 0,
	}
	return &ret, nil
}

func getHash(password string) (*pwnedHash, error) {
	h := sha1.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		return nil, err
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	hash = strings.ToUpper(hash)
	if len(hash) < 5 {
		return nil, fmt.Errorf("unable to hash password")
	}

	result := pwnedHash{
		Hash:  hash,
		Range: hash[0:5],
	}

	return &result, nil
}

//Reverse : Reverse the input string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

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

/*
LengthScore :
Scores password's length
	total = 30
	length<=7 = 0
	length==8 = 4
	length==9 = 10
	10 <= length <= 15 = 15
	16 <= length <= 19 = 21
	20 <= length <= 24 = 26
	length >=25 = 30
*/
func LengthScore(password string) float64 {
	length := len(password)
	if length <= 7 {
		return 0
	}
	if length == 8 {
		return 4
	}
	if length == 9 {
		return 10
	}
	if length > 9 && length < 16 {
		return 15
	}
	if length > 15 && length < 20 {
		return 21
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

/*
pwnedPwds returns the scores for pwned password and
reversed password.

Found : -30
Not Found : 0

Found (Reversed): -10
Not Found (Reversed): 0
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
		scoreKnownPwd = -30
	} else {
		scoreKnownPwd = 0
	}

	if knownPwdReverse.Pwned {
		scoreKnownPwdReverse = -10
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
		knownStr = "Yes (-30)"
	} else {
		knownStr = "No"
	}
	if scoreKnownPwdReverse != 0 {
		knownStrReverse = "Yes (-10)"
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

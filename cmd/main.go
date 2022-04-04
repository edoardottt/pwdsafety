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
	"fmt"
	"os"

	beauty "github.com/edoardottt/pwdsafety/internal"
	"github.com/edoardottt/pwdsafety/pkg/hash"
	"github.com/edoardottt/pwdsafety/pkg/scoring"
	"github.com/fatih/color"
)

func main() {
	beauty.Beautify()
	password := ReadSingleInput("Password")
	CheckPwd(password)
	words := ReadAllFiles("pwds")
	score := scoring.Grader(words, password)
	DisplayResult(score)
	if score <= 68 {
		crackTime := scoring.CrackTime(password)
		println("[?] Estimated password cracking time: " + scoring.ShowCrackTime(crackTime))
		fmt.Println("[-] ------------------------------")
		randomPwd := SuggestPwd(words)
		password = randomPwd
	}
	fmt.Println("[&] Hash functions for " + password + " :")
	fmt.Println("[&] MD4 : " + hash.GetMD4Hash(password))
	fmt.Println("[&] MD5 : " + hash.GetMD5Hash(password))
	fmt.Println("[&] SHA1 : " + hash.GetSHA1Hash(password))
	fmt.Println("[&] RIPEMD160 : " + hash.GetRipemd160Hash(password))
	fmt.Println("[&] SHA2-224 : " + hash.GetSHA224Hash(password))
	fmt.Println("[&] SHA2-256 : " + hash.GetSHA256Hash(password))
	fmt.Println("[&] SHA2-384 : " + hash.GetSHA384Hash(password))
	fmt.Println("[&] SHA2-512 : " + hash.GetSHA512Hash(password))
	fmt.Println("[&] SHA3-224 : " + hash.GetSHA3224Hash(password))
	fmt.Println("[&] SHA3-256 : " + hash.GetSHA3256Hash(password))
	fmt.Println("[&] SHA3-384 : " + hash.GetSHA3384Hash(password))
	fmt.Println("[&] SHA3-512 : " + hash.GetSHA3512Hash(password))
	fmt.Println("[&] Blake2b256 : " + hash.GetBlake2b256Hash(password))
	fmt.Println("[&] Blake2b384 : " + hash.GetBlake2b384Hash(password))
	fmt.Println("[&] Blake2b512 : " + hash.GetBlake2b512Hash(password))
}

//DisplayResult : Display the result for a password
func DisplayResult(score float64) {
	scoreRounded := scoring.Round(fmt.Sprintf("%.2f", score))
	fmt.Println("[!] Final Score: " + fmt.Sprint(scoreRounded) + "/100")
	if score <= 35 {
		color.Red("[X] VERY WEAK")
	}
	if score > 35 && score <= 59 {
		color.Red("[X] WEAK")
	}
	if score > 59 && score <= 68 {
		color.Yellow("[.] REASONABLE")
	}
	if score > 68 && score <= 80 {
		color.Green("[!] STRONG")
	}
	if score > 80 {
		color.Green("[!] VERY STRONG")
	}
}

//SuggestPwd : Suggest a new random password
func SuggestPwd(words [][]string) string {
	randomPwd := scoring.GenerateRandom(30)
	println("[!] You should use this instead...")
	color.Green("[>>] " + randomPwd)
	scoreRandomPwd := scoring.Grader(words, randomPwd)
	DisplayResult(scoreRandomPwd)
	return randomPwd
}

//CheckPwd : Check if a password is useless
func CheckPwd(password string) {
	if len(password) <= 5 {
		println("[X] Hey....Do you know what is password cracking?")
		os.Exit(1)
	}
}

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

	"github.com/fatih/color"
)

func main() {
	beauty.Beautify()
	password := ReadSingleInput("Password")
	CheckPwd(password)
	words := ReadAllFiles("pwds")
	score := Grader(words, password)
	DisplayResult(score)
	if score <= 68 {
		crackTime := CrackTime(password)
		println("[?] Estimated password cracking time: " + ShowCrackTime(crackTime))
		fmt.Println("[-] ------------------------------")
		randomPwd := SuggestPwd(words)
		password = randomPwd
	}
	fmt.Println("[&] Hash functions for " + password + " :")
	fmt.Println("[&] MD4 : " + GetMD4Hash(password))
	fmt.Println("[&] MD5 : " + GetMD5Hash(password))
	fmt.Println("[&] SHA1 : " + GetSHA1Hash(password))
	fmt.Println("[&] RIPEMD160 : " + GetRipemd160Hash(password))
	fmt.Println("[&] SHA2-224 : " + GetSHA224Hash(password))
	fmt.Println("[&] SHA2-256 : " + GetSHA256Hash(password))
	fmt.Println("[&] SHA2-384 : " + GetSHA384Hash(password))
	fmt.Println("[&] SHA2-512 : " + GetSHA512Hash(password))
	fmt.Println("[&] SHA3-224 : " + GetSHA3224Hash(password))
	fmt.Println("[&] SHA3-256 : " + GetSHA3256Hash(password))
	fmt.Println("[&] SHA3-384 : " + GetSHA3384Hash(password))
	fmt.Println("[&] SHA3-512 : " + GetSHA3512Hash(password))
	fmt.Println("[&] Blake2b256 : " + GetBlake2b256Hash(password))
	fmt.Println("[&] Blake2b384 : " + GetBlake2b384Hash(password))
	fmt.Println("[&] Blake2b512 : " + GetBlake2b512Hash(password))
}

//DisplayResult : Display the result for a password
func DisplayResult(score float64) {
	scoreRounded := Round(fmt.Sprintf("%.2f", score))
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
	randomPwd := GenerateRandom(30)
	println("[!] You should use this instead...")
	color.Green("[>>] " + randomPwd)
	scoreRandomPwd := Grader(words, randomPwd)
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

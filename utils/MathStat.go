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
	"fmt"
	"math"
	"strconv"
)

//Entropy : Returns the entropy of password
func Entropy(password string) float64 {
	var E float64
	var pool float64 = 95
	length := float64(len(password))
	E = math.Log2(math.Pow(pool, length))
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
	const GPU float64 = 2000000000
	var bots float64 = 40000
	var KPS = bots * GPU
	var combinations float64
	length := float64(len(password))
	var pool float64 = 95
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

package utils

import (
	"math"
)

//Returns the entropy of password
func Entropy(password string) float64 {
	var E float64
	var pool float64 = 95
	length := float64(len(password))
	E = math.Log2(math.Pow(pool, length))
	return E
}

//Counts the different types in password
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

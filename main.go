package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"pwd-safety/Utils"
)

func main() {
	password := Utils.ReadSingleInput("Password")
	CheckPwd(password)
	words := Utils.ReadWords("known-pwd.txt")
	score := Utils.Grader(words, password)
	DisplayResult(score)
	if score <= 68 {
		SuggestPwd(words)
	}
}

//Display the result for a password
func DisplayResult(score float64) {
	scoreRounded := fmt.Sprintf("%.2f", score)
	fmt.Printf("Score:%s/100", scoreRounded)
	print("\n")
	if score <= 35 {
		color.Red("VERY WEAK")
	}
	if score > 35 && score <= 59 {
		color.Red("WEAK")
	}
	if score > 59 && score <= 68 {
		color.Yellow("REASONABLE")
	}
	if score > 68 && score <= 80 {
		color.Green("STRONG")
	}
	if score > 80 {
		color.Green("VERY STRONG")
	}
}

//Suggest a new random password
func SuggestPwd(words []string) {
	engWords := Utils.ReadWords("eng_words_list.txt")
	randomPwd := Utils.GenerateRandom(engWords)
	scoreRandomPwd := Utils.Grader(words, randomPwd)
	println("You should use this instead...")
	println(randomPwd)
	DisplayResult(scoreRandomPwd)
}

//Check if a password is blank
func CheckPwd(password string) {
	if len(password) == 0 {
		println("Hey....Do you know what is password cracking?")
		os.Exit(1)
	}
}

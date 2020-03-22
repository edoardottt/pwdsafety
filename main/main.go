package main

import (
	"fmt"
	"github.com/fatih/color"
	"pwd-safety/Utils"
)

func main() {
	password := Utils.ReadSingleInput("Password")
	words := Utils.ReadWords("known-pwd.txt")
	score := Utils.Grader(words, password)
	DisplayResult(score)
}

func DisplayResult(score float64) {
	scoreRounded := fmt.Sprintf("%.2f", score)
	fmt.Printf("Score:%s/100",scoreRounded)
	print("\n")
	if score <= 35 { color.Red("VERY WEAK") }
	if score > 35 && score <= 59 { color.Red("WEAK") }
	if score > 59 && score <= 68 { color.Yellow("REASONABLE") }
	if score > 68 && score <= 80 { color.Green("STRONG") }
	if score > 80 { color.Green("VERY STRONG") }
}
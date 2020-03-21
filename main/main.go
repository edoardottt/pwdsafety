package main

import (
	"fmt"
	"github.com/fatih/color"
	"pwd-safety/Utils"
)

func main() {
	password := Utils.ReadSingleInput("Password")
	words := Utils.ReadWords("known-pwd.txt")
	color.Green("Prints %s in green.", "text")
	color.Red("We have red")
	println(password)
	score := Utils.Grader(words, password)
	fmt.Println("Score:", score)
}
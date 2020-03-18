package main

import (
	"fmt"
	"github.com/edoardottt/pwd-safety/Utils"
	"github.com/fatih/color"
)

func main() {
	password := Utils.ReadSingleInput("Password")
	words := Utils.ReadWords("words.txt")
	color.Green("Prints %s in green.", "text")
	color.Red("We have red")
	println(password)
	boo := Utils.FindExactly(words,password)
	fmt.Printf("%t", boo)
}
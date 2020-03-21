package main

import (
	"fmt"
	"github.com/fatih/color"
	"pwd-safety/Utils"
)

func main() {
	password := Utils.ReadSingleInput("Password")
	words := Utils.ReadWords("words.txt")
	color.Green("Prints %s in green.", "text")
	color.Red("We have red")
	println(password)
	boo := Utils.FindExactly(words,password)
	numbers := Utils.IsThereNumber(password)
	upper := Utils.IsThereUpperCase(password)
	lower := Utils.IsThereLowerCase(password)
	howMany := Utils.HowManyDifferents(password)
	fmt.Printf("%s %t", "Found in the famous passwords:",boo)
	println("")
	fmt.Println("There is numbers:",numbers)
	fmt.Println("There is UPPERCASE:", upper)
	fmt.Println("There is lowercase:", lower)
	fmt.Println("How many different character: ", howMany)
}
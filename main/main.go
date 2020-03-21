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
	boo := Utils.FindExactly(words,password)
	boa := Utils.FindExactlyReversed(words,password)
	length := len(password)
	numbers := Utils.IsThereNumber(password)
	upper := Utils.IsThereUpperCase(password)
	lower := Utils.IsThereLowerCase(password)
	symbols := Utils.IsThereSymbol(password)
	howMany := Utils.HowManyDifferents(password)
	types := Utils.HowManyTypes(password)
	count := Utils.CountTypeElements(password)
	entropy := Utils.Entropy(password)
	fmt.Printf("%s %t", "Found in the famous passwords:",boo)
	println("")
	fmt.Printf("%s %t", "Found reversed in the famous passwords:",boa)
	println("")
	fmt.Println("Password's length:", length)
	fmt.Println("There is numbers:", numbers)
	fmt.Println("There is UPPERCASE:", upper)
	fmt.Println("There is lowercase:", lower)
	fmt.Println("There is symbols:", symbols)
	fmt.Println("How many different character:", howMany)
	fmt.Println("How many different types:", types)
	fmt.Println("Count elements:", count)
	fmt.Println("Entropy:",entropy)
}
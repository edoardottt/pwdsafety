package main

import (
	"fmt"
	"github.com/edoardottt/pwd-safety/Utils"
	"github.com/fatih/color"
)

func main() {
	inputs := Utils.ReadInput()
	words := Utils.ReadWords("words.txt")
	color.Green("Prints %s in green.", "text")
	color.Red("We have red")
	fmt.Printf("%v", words)
	fmt.Printf("%v", inputs)
}
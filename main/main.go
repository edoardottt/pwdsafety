package main

import (
	"github.com/edoardottt/pwd-safety/Utils"
	"github.com/fatih/color"
)

func main() {
	words := ReadWords("words.txt")
	color.Green("Prints %s in green.", "text")
	color.Red("We have red")
}
package Utils

import (
	"bufio"
	"fmt"
	"os"
)

//Reading all inputs from stdin
func ReadInput() []string {
	var result []string
	inputs := [5]string{"name", "surname", "birthday(ddmmyyyy)", "telephone number", "pet's name"}
	for _, value := range inputs {
		result = append(result, ReadSingleInput(value))
	}
	return result
}

//Reading one single input
func ReadSingleInput(input string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter " + input + ": ")
	text, _ := reader.ReadString('\n')
	ind := len(text)
	if ind > 0 && text[ind-1] == '\n' {
		text = text[:ind-1]
	}
	return text
}

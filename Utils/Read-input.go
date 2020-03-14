package Utils

import (
	"bufio"
	"fmt"
	"os"
)


//Reading one single input
func ReadSingleInput(input string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter "+ input + ": ")
	text, _ := reader.ReadString('\n')
	return text
}
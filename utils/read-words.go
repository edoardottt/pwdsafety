package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readWords(fileInput string) {
	file, err := os.Open(fileInput)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	result []string
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}
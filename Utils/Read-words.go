package Utils

import (
	"bufio"
	"log"
	"os"
)

func ReadWords(fileInput string) []string {
	file, err := os.Open(fileInput)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines
}
package Utils

import (
	"bufio"
	"log"
	"os"
)

//Read words from a file
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
	err = file.Close()
	if err != nil {
		log.Fatalf("Failed closing file: %s", err)
	}
	return txtlines
}
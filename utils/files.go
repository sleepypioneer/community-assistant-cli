package utils

import (
	"io/ioutil"
	"os"
)

func ReadFileText(inputFile string) (string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	inputBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	inputText := string(inputBytes)
	return inputText, nil
}

func WriteTextToFile(outputFile string, outputText string) {
	ioutil.WriteFile(outputFile, []byte(outputText), 0644)
}

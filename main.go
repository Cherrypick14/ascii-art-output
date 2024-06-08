package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if  len(os.Args) != 4 {
		fmt.Println("Usage: go run . --output=<fileName.txt> [STRING] [BANNER]")
		return
	}

	outputFileName, inputText, banner, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if inputText == "\\n" {
		fmt.Print("\n")
		return
	}

	if inputText == "\\a" || inputText == "\\0" || inputText == "\\f" || inputText == "\\v" || inputText == "\\r" {
		fmt.Println("Error: Non-printable character", inputText)
		return
	}

	inputText = strings.ReplaceAll(inputText, "\\t", "    ")
	inputText = strings.ReplaceAll(inputText, "\\b", "\b")

	for i := 0; i < len(inputText); i++ {
		indexb := strings.Index(inputText, "\b")
		if indexb > 0 {
			inputText = inputText[:indexb-1] + inputText[indexb+1:]
		}
	}

	words := strings.Split(inputText, "\\n")

	bannerFile := banner + ".txt"

	contents, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	contents2 := strings.Split(string(contents), "\n")

	output := ascii.AsciiArt(words, contents2)

	if outputFileName != "" {
		err := writeToFile(outputFileName, output)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	} else {
		fmt.Println(output)
	}
}

func parseArgs(args []string) (string, string, string, error) {
	var outputFileName, inputText, banner string

	for _, arg := range args[1:] {
		if strings.HasPrefix(arg, "--output=") {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 || parts[1] == "" {
				return "", "", "", fmt.Errorf("Usage: go run . --output=<fileName.txt> [STRING] [BANNER]")
			}
			outputFileName = parts[1]
		} else if inputText == "" {
			inputText = arg
		} else {
			banner = arg
		}
	}

	if inputText == "" || banner == "" {
		return "", "", "", fmt.Errorf("Usage: go run . --output=<fileName.txt> [STRING] [BANNER]")
	}

	return outputFileName, inputText, banner, nil
}

func writeToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

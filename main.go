package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		return
	}
	// Grab string to generate Ascii represatantion.
	inputText := os.Args[1]
	if len(os.Args) >= 3 && len(os.Args) <= 4 {
		inputText = os.Args[2]
	}
	switch inputText {
	case "":
		return
	case "\\a", "\\0", "\\f", "\\v", "\\r":
		fmt.Println("Error: Non printable character", inputText)
		return
	}

	inputText = strings.ReplaceAll(inputText, "\\t", "    ")
	inputText = strings.ReplaceAll(inputText, "\\b", "\b")
	inputText = strings.ReplaceAll(inputText, "\\n", "\n")
	// Logic process for handlng the backspace.
	for i := 0; i < len(inputText); i++ {
		indexB := strings.Index(inputText, "\b")
		if indexB > 0 {
			inputText = inputText[:indexB-1] + inputText[indexB+1:]
		}
	}
	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(inputText, "\n")

	// setting the bannerfile to be used according to user input.
	banner := "standard"
	if len(os.Args) == 4 {
		banner = strings.ToLower(os.Args[3])
	}
	bannerFile := banner + ".txt"

	// Read the contents of banner file.
	bannerText, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	// Confirm file information.
	fileInfo, err := os.Stat(bannerFile)
	if err != nil {
		fmt.Println("Error reading file information", err)
		return
	}
	fileSize := fileInfo.Size()

	if fileSize == 6623 || fileSize == 4702 || fileSize == 7462 {
		// Split the content to a string slice and separate with newline.
		contents := strings.Split(string(bannerText), "\n")

		outputfilename := "banner.txt"
		if len(os.Args) == 3 || len(os.Args) == 4 {
			outputfilename = os.Args[1]
			if strings.HasPrefix(outputfilename, "--output=") && strings.HasSuffix(outputfilename, ".txt") {
				outputfilename = os.Args[1]
				outputfilename = outputfilename[9:]
			} else {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
				return
			}
		}

		// Call the AsciiArt function for the returned string.
		asciiArt := ascii.AsciiArt(words, contents)
		err := os.WriteFile(outputfilename, []byte(asciiArt), 0o644)
		if err != nil {
			fmt.Println("Failed to write to file", err)
			return
		}
	} else {
		fmt.Println("Error with the file size", fileSize)
		return
	}
}

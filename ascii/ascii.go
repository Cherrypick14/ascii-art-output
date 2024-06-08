package ascii

import (
	"strings"
)

// AsciiArt processes words, printing their ASCII art
// character by character and adding new lines as needed.
func AsciiArt(words []string, contents2 []string) string {
	var output strings.Builder
	countSpace := 0
	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						return "Error: Input contains non-ASCII characters"
					}
					// Print the calculated index of 'char' Ascii Art in contents2.
					output.WriteString(contents2[int(char-' ')*9+1+i])
				}
				output.WriteString("\n")
			}
		} else {
			countSpace++
			if countSpace < len(words) {
				output.WriteString("\n")
			}
		}
	}
	return output.String()
}

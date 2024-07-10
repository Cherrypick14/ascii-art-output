package ascii

import (
	"fmt"
	"strings"
)

// AsciiArt processes words, printing their ASCII art
// character by character and adding new lines as needed.
func AsciiArt(words []string, contents []string) string {
	countSpace := 0
	// Define a strings builder
	var artBuild strings.Builder

	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						fmt.Println("Error: Input contains non-ASCII characters")
						return ""
					}
					// Add to artBuild calculated index to retrieve art representation.
					artBuild.WriteString(contents[int(char-' ')*9+1+i])
				}
				artBuild.WriteString("\n") // Add newline after each word line
			}
		} else {
			countSpace++
			if countSpace < len(words) {
				artBuild.WriteString("\n") // Add newline for empty words
			}
		}
	}
	return artBuild.String()
}

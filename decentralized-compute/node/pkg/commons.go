package pkg

import (
	"fmt"
	"regexp"
	"strings"
)

func Regex(input, pattern string) string {
	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllStringSubmatch(input, -1)

	// Print the results
	for _, match := range matches {
		data := match[1] // match[1] contains the content inside the {}
		fmt.Println(data)
	}

	return ""
}

func ExtractContent(input string) (string, error) {
	var result strings.Builder
	stack := []rune{}
	found := false

	for _, char := range input {
		if char == '{' {
			if len(stack) == 0 {
				found = true // Found the outermost opening brace
			}
			stack = append(stack, char)
		} else if char == '}' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 && found {
				break // Found the outermost closing brace
			}
		}

		if found {
			result.WriteRune(char) // Append characters inside the outermost braces
		}
	}

	if len(stack) > 0 {
		return "", fmt.Errorf("mismatched braces")
	}

	abc := result.String()
	if string(abc[len(abc)-1]) != "}" {
		abc += "}"
	}

	return abc, nil
}

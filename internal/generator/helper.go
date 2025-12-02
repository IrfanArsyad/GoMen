package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// toSnakeCase converts PascalCase or camelCase to snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// toPascalCase converts snake_case or camelCase to PascalCase
func toPascalCase(s string) string {
	words := strings.Split(s, "_")
	var result strings.Builder
	for _, word := range words {
		if len(word) > 0 {
			result.WriteString(strings.ToUpper(string(word[0])))
			result.WriteString(strings.ToLower(word[1:]))
		}
	}
	return result.String()
}

// toCamelCase converts snake_case or PascalCase to camelCase
func toCamelCase(s string) string {
	pascal := toPascalCase(s)
	if len(pascal) > 0 {
		return strings.ToLower(string(pascal[0])) + pascal[1:]
	}
	return pascal
}

// toPlural converts a word to its plural form (simple version)
func toPlural(s string) string {
	if strings.HasSuffix(s, "y") {
		return s[:len(s)-1] + "ies"
	}
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "x") ||
		strings.HasSuffix(s, "ch") || strings.HasSuffix(s, "sh") {
		return s + "es"
	}
	return s + "s"
}

// writeFile writes content to a file, creating directories if needed
func writeFile(filePath, content string) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("file already exists: %s", filePath)
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// printSuccess prints a success message
func printSuccess(fileType, filePath string) {
	fmt.Printf("\033[32m✓\033[0m %s created successfully: %s\n", fileType, filePath)
}

// printError prints an error message
func printError(err error) {
	fmt.Printf("\033[31m✗\033[0m Error: %s\n", err)
}

// getProjectRoot returns the project root directory
func getProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return dir
}

package util

import "strings"

func RemoverDuplicados(clave string) string {
	seen := make(map[rune]bool)
	builder := strings.Builder{}

	for _, char := range clave {
		if !seen[char] {
			seen[char] = true
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func LimpiarData(data string) string {

	return strings.ReplaceAll(strings.ToUpper(data), " ", "")
}

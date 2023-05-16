package monoalfabetica

import (
	"criptograms/main/data"
	"strings"
)

type TransformacionAfin struct {
}

func (r TransformacionAfin) Cypher(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.Message, " ", ""))
	return substitutionCipher(message, data.Clave)
}

func (r TransformacionAfin) Decrypt(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.Message, " ", ""))
	return inverseSubstitutionCipher(message, data.Clave)
}
func affineTransformation(a, b, x int) int {
	return (a*x + b) % 26
}

func substitutionCipher(message, keyword string) string {
	// Convertir la clave a valores numéricos
	keyword = strings.ToUpper(keyword)
	keywordValues := make([]int, len(keyword))
	for i, char := range keyword {
		keywordValues[i] = int(char - 'A')
	}

	// Calcular los coeficientes 'a' y 'b'
	a := 0
	b := 0
	for _, value := range keywordValues {
		a += value
	}
	for i, value := range keywordValues {
		b += value * (i + 1)
	}

	ciphertext := ""
	for _, char := range message {
		if char >= 'A' && char <= 'Z' {
			// Aplicar la transformación afín a cada letra del mensaje
			x := int(char - 'A')
			y := affineTransformation(a, b, x)
			ciphertext += string(y + 'A')
		} else {
			// Conservar los caracteres no alfabéticos sin cambios
			ciphertext += string(char)
		}
	}

	return ciphertext
}
func inverseAffineTransformation(a, b, y int) int {
	// Calcular el inverso multiplicativo de 'a' módulo 26
	for i := 0; i < 26; i++ {
		if (a*i)%26 == 1 {
			aInverse := i
			return (aInverse * (y - b + 26)) % 26
		}
	}
	return -1
}

func inverseSubstitutionCipher(ciphertext, keyword string) string {
	// Convertir la clave a valores numéricos
	keyword = strings.ToUpper(keyword)
	keywordValues := make([]int, len(keyword))
	for i, char := range keyword {
		keywordValues[i] = int(char - 'A')
	}

	// Calcular los coeficientes 'a' y 'b'
	a := 0
	b := 0
	for _, value := range keywordValues {
		a += value
		b += value * (len(keywordValues) + 1)
	}

	plaintext := ""
	for _, char := range ciphertext {
		if char >= 'A' && char <= 'Z' {
			// Aplicar la transformación afín inversa a cada letra del mensaje cifrado
			y := int(char - 'A')
			x := inverseAffineTransformation(a, b, y)
			plaintext += string(x + 'A')
		} else {
			// Conservar los caracteres no alfabéticos sin cambios
			plaintext += string(char)
		}
	}

	return plaintext
}

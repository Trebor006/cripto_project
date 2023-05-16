package trasposicion

import (
	"criptograms/main/data"
	"strings"
)

type Grupos struct {
}

func (r Grupos) Cypher(data data.Data) string {
	message := strings.ReplaceAll(strings.ToUpper(data.Message), " ", "")
	return transposeCipher(message, data.Clave)
}

func (r Grupos) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!
	message := strings.ReplaceAll(strings.ToUpper(data.EncryptedMessage), " ", "")

	return inverseTransposeCipher(message, data.Clave)
}
func transposeCipher(message, key string) string {
	// Obtener la longitud del grupo y ajustar el mensaje si es necesario
	groupSize := len(key)
	messageLength := len(message)
	remainder := messageLength % groupSize
	if remainder != 0 {
		padding := strings.Repeat("X", groupSize-remainder+1)
		message += padding
		messageLength = len(message)
	}

	// Dividir el mensaje en grupos
	groups := splitCadena(message, groupSize)

	textoCifrado := ""
	for _, grupo := range groups {
		textoCifrado += reorderTexto(grupo, key)
	}

	return textoCifrado
}

func splitCadena(cadena string, longitud int) []string {
	grupos := len(cadena) / longitud
	resultado := make([]string, grupos)

	for i := 0; i < grupos; i++ {
		inicio := i * longitud
		fin := inicio + longitud

		resultado[i] = cadena[inicio:fin]
	}

	return resultado
}

func reorderTexto(texto, clave string) string {
	if len(texto) != len(clave) {
		return ""
	}

	mapping := make([]rune, len(clave))
	for i, c := range clave {
		mapping[i] = rune(texto[int(c-'1')])
	}

	resultado := ""
	for _, r := range mapping {
		resultado += string(r)
	}

	return resultado
}

func inverseTransposeCipher(ciphertext, key string) string {
	// Obtener la longitud del grupo
	groupSize := len(key)

	// Calcular el número de grupos
	numGroups := len(ciphertext) / groupSize

	// Obtener la permutación inversa
	inversePermutation := make([]int, groupSize)
	for i, char := range key {
		inversePermutation[char-'1'] = i + 1
	}

	// Obtener los grupos originales
	groups := make([]string, numGroups)
	for i := 0; i < numGroups; i++ {
		groups[i] = ciphertext[i*groupSize : (i+1)*groupSize]
	}

	// Permutar los grupos a su orden original
	originalGroups := make([]string, numGroups)
	for i, index := range inversePermutation {
		originalGroups[index-1] = groups[i]
	}

	// Concatenar los grupos originales para obtener el mensaje descifrado
	plaintext := strings.Join(originalGroups, "")

	return plaintext
}

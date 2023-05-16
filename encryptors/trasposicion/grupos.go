package trasposicion

import (
	"criptograms/main/data"
	"strings"
)

type Grupos struct {
}

func (r Grupos) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return transposeCipher(data.Message, data.Clave)
}

func (r Grupos) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return inverseTransposeCipher(data.Message, data.Clave)
}
func transposeCipher(message, key string) string {
	// Obtener la longitud del grupo y ajustar el mensaje si es necesario
	groupSize := len(key)
	messageLength := len(message)
	remainder := messageLength % groupSize
	if remainder != 0 {
		padding := strings.Repeat("X", groupSize-remainder)
		message += padding
		messageLength = len(message)
	}

	// Dividir el mensaje en grupos
	groups := make([]string, messageLength/groupSize)
	for i := 0; i < messageLength; i += groupSize {
		groups[i/groupSize] = message[i : i+groupSize]
	}

	// Permutar los grupos según la clave
	permutation := make([]int, groupSize)
	for i, char := range key {
		permutation[i] = int(char - '0')
	}

	permutedGroups := make([]string, len(groups))
	for i, index := range permutation {
		permutedGroups[i] = groups[index-1]
	}

	// Concatenar los grupos permutados para obtener el mensaje cifrado
	ciphertext := strings.Join(permutedGroups, "")

	return ciphertext
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

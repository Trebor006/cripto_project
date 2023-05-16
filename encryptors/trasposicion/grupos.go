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
	groups := generarGrupos(message, groupSize)

	textoCifrado := ""
	for _, grupo := range groups {
		textoCifrado += reorderTexto(grupo, key)
	}

	return textoCifrado
}

func generarGrupos(cadena string, longitud int) []string {
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

	groups := generarGrupos(ciphertext, groupSize)

	textoCifrado := ""
	for _, grupo := range groups {
		textoCifrado += revertirReordenamiento(grupo, key)
	}

	return textoCifrado
}

func revertirReordenamiento(texto, clave string) string {
	if len(texto) != len(clave) {
		return ""
	}

	mapping := make([]rune, len(clave))
	for i, c := range clave {
		mapping[int(c-'1')] = rune(texto[i])
	}

	resultado := ""
	for _, r := range mapping {
		resultado += string(r)
	}

	return resultado
}

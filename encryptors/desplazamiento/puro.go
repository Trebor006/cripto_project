package desplazamiento

import (
	"criptograms/main/data"
	util "criptograms/main/util"
	"strings"
)

type Puro struct {
}

func (r Puro) Cypher(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.Message, " ", ""))
	alfabeto := util.ObtenerAlfabetoEspanol()

	textoEncriptado := ""
	for _, char := range message {
		pos := util.ObtenerPosicion(alfabeto, string(char))
		nuevaPos := pos + 3
		if nuevaPos >= len(alfabeto) {
			nuevaPos = nuevaPos - len(alfabeto)
		}

		textoEncriptado += alfabeto[nuevaPos]
	}

	return textoEncriptado
}

func (r Puro) Decrypt(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.EncryptedMessage, " ", ""))
	alfabeto := util.ObtenerAlfabetoEspanol()

	textoEncriptado := ""
	for _, char := range message {
		pos := util.ObtenerPosicion(alfabeto, string(char))
		nuevaPos := pos - 3
		if nuevaPos < 0 {
			nuevaPos = len(alfabeto) + nuevaPos
		}

		textoEncriptado += alfabeto[nuevaPos]
	}

	return textoEncriptado
}

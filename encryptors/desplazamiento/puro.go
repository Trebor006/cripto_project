package desplazamiento

import (
	"cripto_project/main/data"
	util "cripto_project/main/util"
)

type Puro struct {
}

func (r Puro) Cypher(data data.Data) string {
	message := util.LimpiarData(data.Message)

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
	message := util.LimpiarData(data.EncryptedMessage)

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

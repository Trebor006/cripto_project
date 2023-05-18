package homofonos

import (
	"cripto_project/main/data"
	"cripto_project/main/util"
	"strconv"
	"strings"
)

type HomofonoOrdenMayor struct {
}

func (r HomofonoOrdenMayor) Cypher(data data.Data) string {
	diccionario := data.Diccionario
	message := strings.ReplaceAll(strings.ToUpper(data.Message), " ", "")
	randomGenerator := data.RandomGenerator

	textoEncriptado := ""
	if len(message)%2 == 1 {
		message += "X"
	}

	for i := 0; i < len(message)-1; i += 2 {
		clave := message[i : i+2]

		detail := diccionario[strings.ToUpper(clave)]
		pos := util.GenerarNumeroRandom(len(detail), randomGenerator)
		textoEncriptado += strconv.Itoa(detail[pos]) + " "
	}

	return strings.Trim(textoEncriptado, " ")
}

func (r HomofonoOrdenMayor) Decrypt(data data.Data) string {
	diccionario := data.Diccionario
	valuesToDecrypt := strings.Split(strings.Trim(data.EncryptedMessage, " "), " ")

	textoClaro := ""
	for _, b := range valuesToDecrypt {
		valorEncontrado := false
		for clave, lista := range diccionario {
			for _, valor := range lista {
				if strconv.Itoa(valor) == b {
					textoClaro += clave
					valorEncontrado = true
					break
				}
			}
			if valorEncontrado {
				break
			}
		}
	}

	return textoClaro
}

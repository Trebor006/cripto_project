package homofonos

import (
	"criptograms/main/data"
	"criptograms/main/util"
	"strconv"
	"strings"
)

type HomofonoPrimerOrden struct {
}

func (h HomofonoPrimerOrden) Cypher(data data.Data) string {
	diccionario := data.Diccionario
	message := strings.ReplaceAll(data.Message, " ", "")
	textoEncriptado := ""
	var randomGenerator = data.RandomGenerator
	for _, char := range message {
		detail := diccionario[strings.ToUpper(string(char))]
		pos := util.GenerarNumeroRandom(len(detail), randomGenerator)
		textoEncriptado += strconv.Itoa(detail[pos]) + " "
	}

	return strings.Trim(textoEncriptado, " ")
}

func (h HomofonoPrimerOrden) Decrypt(data data.Data) string {
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

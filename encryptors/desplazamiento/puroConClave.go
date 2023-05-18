package desplazamiento

import (
	"cripto_project/main/data"
	"cripto_project/main/util"
	"fmt"
	"strings"
)

type PuroConClave struct {
}

func (r PuroConClave) Cypher(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.Message, " ", ""))
	clave := util.RemoverDuplicados(strings.ToUpper(strings.ReplaceAll(data.Clave, " ", "")))

	nuevoAlfabeto := generarNuevoAlfabeto(clave, data.PosCifradoCesar)

	textoEncriptado := ""
	for _, char := range message {
		pos := util.ObtenerPosicion(nuevoAlfabeto, string(char))
		nuevaPos := pos + data.PosCifradoCesar
		if nuevaPos >= len(nuevoAlfabeto) {
			nuevaPos = nuevaPos - len(nuevoAlfabeto)
		}

		textoEncriptado += nuevoAlfabeto[nuevaPos]
	}

	return textoEncriptado
}

func generarNuevoAlfabeto(clave string, pos int) []string {
	alfabeto := util.ObtenerAlfabetoEspanol()
	nuevoAlfabeto := []string{}
	for _, char := range clave {
		nuevoAlfabeto = append(nuevoAlfabeto, string(char))
	}

	for _, char := range alfabeto {
		if util.ObtenerPosicion(nuevoAlfabeto, string(char)) == -1 {
			nuevoAlfabeto = append(nuevoAlfabeto, string(char))
		}
	}

	fmt.Print("Nuevo Alfabeto  [")
	for _, char := range nuevoAlfabeto {
		fmt.Printf("%s ,", char)
	}
	fmt.Print("]\n")

	return nuevoAlfabeto
}

func (r PuroConClave) Decrypt(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.EncryptedMessage, " ", ""))
	clave := util.RemoverDuplicados(strings.ToUpper(strings.ReplaceAll(data.Clave, " ", "")))
	nuevoAlfabeto := generarNuevoAlfabeto(clave, data.PosCifradoCesar)

	textoEncriptado := ""
	for _, char := range message {
		pos := util.ObtenerPosicion(nuevoAlfabeto, string(char))
		nuevaPos := pos - data.PosCifradoCesar
		if nuevaPos < 0 {
			nuevaPos = len(nuevoAlfabeto) + nuevaPos
		}

		textoEncriptado += nuevoAlfabeto[nuevaPos]
	}

	return textoEncriptado
}

package desplazamiento

import (
	"cripto_project/main/data"
	"cripto_project/main/util"
	"fmt"
)

type PuroConClave struct {
}

func (r PuroConClave) Cypher(data data.Data) string {
	message := util.LimpiarData(data.Message)
	clave := util.RemoverDuplicados(util.LimpiarData(data.Clave))

	nuevoAlfabeto := generarNuevoAlfabeto(clave)

	textoEncriptado := ""
	for _, char := range message {
		pos := util.ObtenerPosicion(nuevoAlfabeto, string(char))
		nuevaPos := pos + 3
		if nuevaPos >= len(nuevoAlfabeto) {
			nuevaPos = nuevaPos - len(nuevoAlfabeto)
		}

		textoEncriptado += nuevoAlfabeto[nuevaPos]
	}

	return textoEncriptado
}

func generarNuevoAlfabeto(clave string) []string {
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
	message := util.LimpiarData(data.EncryptedMessage)
	clave := util.RemoverDuplicados(util.LimpiarData(data.Clave))

	nuevoAlfabeto := generarNuevoAlfabeto(clave)
	textoEncriptado := ""
	for _, char := range message {
		pos := util.ObtenerPosicion(nuevoAlfabeto, string(char))
		nuevaPos := pos - 3
		if nuevaPos < 0 {
			nuevaPos = len(nuevoAlfabeto) + nuevaPos
		}

		textoEncriptado += nuevoAlfabeto[nuevaPos]
	}

	return textoEncriptado
}

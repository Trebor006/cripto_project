package trasposicion

import (
	"cripto_project/main/data"
	"cripto_project/main/util"
	"fmt"
	"strings"
)

type Filas struct {
}

func (r Filas) Cypher(data data.Data) string {
	var message = util.LimpiarData(data.Message)
	var nroFilas = data.NroFilas

	if nroFilas <= 0 || nroFilas >= len(message) {
		return "Error no se puede cifrar!"
	}

	nroColumnas := (len(message) + nroFilas - 1) / nroFilas

	//declarando la matriz de n x m
	matriz := util.InicializarMatriz(nroFilas, nroColumnas)

	pos := 0
	for col := 0; col < nroColumnas; col++ {
		for fil := 0; fil < nroFilas; fil++ {
			if pos <= len(message)-1 {
				matriz[fil][col] = rune(message[pos])
			} else {
				matriz[fil][col] = 'X'
			}
			pos++
			util.ImprimirMatriz(matriz)
		}
	}

	//obtener texto cifrado
	var textoCifrado strings.Builder
	for fil := 0; fil < nroFilas; fil++ {
		for col := 0; col < nroColumnas; col++ {
			textoCifrado.WriteRune(matriz[fil][col])
		}

		fmt.Println(textoCifrado.String())
	}

	return textoCifrado.String()
}

func (r Filas) Decrypt(data data.Data) string {
	message := util.LimpiarData(data.EncryptedMessage)

	var nroFilas = data.NroFilas
	if nroFilas <= 0 || nroFilas >= len(message) {
		return "Error no se puede descifrar!"
	}

	nroColumnas := util.Dividir(len(message), nroFilas)

	// Crear matriz de nroColumnas x nroFilas y llenarla con caracteres de mensaje cifrado
	matriz := util.InicializarMatriz(nroFilas, nroColumnas)

	pos := 0
	for fil := 0; fil < nroFilas; fil++ {
		for col := 0; col < nroColumnas; col++ {
			matriz[fil][col] = rune(message[pos])
			pos++
		}

		util.ImprimirMatriz(matriz)
	}

	util.ImprimirMatriz(matriz)

	// Obtener mensaje original descifrando la matriz
	var mensajeOriginal strings.Builder
	for col := 0; col < nroColumnas; col++ {
		for fil := 0; fil < nroFilas; fil++ {
			mensajeOriginal.WriteRune(matriz[fil][col])
		}

		fmt.Println(mensajeOriginal.String())
	}

	return mensajeOriginal.String()
}

package sustitucion_monogramica

import (
	"cripto_project/main/data"
	"cripto_project/main/util"
	"fmt"
	"strconv"
	"strings"
)

type PolialfabeticosPeriodicos struct {
}

func (r PolialfabeticosPeriodicos) Cypher(data data.Data) string {
	message := util.LimpiarData(data.Message)
	clave := util.LimpiarData(data.Clave)

	nroFilas := 1
	nroColumnas := len(message)

	//declarando la matriz de n x m
	matriz := util.InicializarMatriz(nroFilas, nroColumnas)

	pos := 0
	for col := 0; col < nroColumnas; col++ {
		if pos < len(message) {
			matriz[0][col] = rune(message[pos])
		} else {
			matriz[0][col] = 'X'
		}
		pos++
		util.ImprimirMatriz(matriz)
	}

	alfabeto := util.ObtenerAlfabetoEspanol()
	posclave := 0
	textoCifrado := strings.Builder{}
	for col := 0; col < nroColumnas; col++ {

		posicionTexto := util.ObtenerPosicion(alfabeto, string(matriz[0][col]))
		posicionClave := util.ObtenerPosicion(alfabeto, string(clave[posclave]))
		nuevaPosicion := posicionTexto + posicionClave
		valorNuevo := string(alfabeto[nuevaPosicion%len(alfabeto)][0])

		fmt.Println(string(matriz[0][col]) + " : " + strconv.Itoa(posicionTexto) + " + " + string(clave[posclave]) + " : " + strconv.Itoa(posicionClave) + " = " + string(valorNuevo) + " : " + strconv.Itoa(nuevaPosicion))

		textoCifrado.WriteString(valorNuevo)

		if posclave < len(clave)-1 {
			posclave++
		} else {
			posclave = 0
		}
	}

	fmt.Println(textoCifrado.String())

	return textoCifrado.String()
}

func (r PolialfabeticosPeriodicos) Decrypt(data data.Data) string {
	message := util.LimpiarData(data.EncryptedMessage)
	clave := util.LimpiarData(data.Clave)

	nroFilas := 1
	nroColumnas := len(message)

	//declarando la matriz de n x m
	matriz := util.InicializarMatriz(nroFilas, nroColumnas)

	pos := 0
	for col := 0; col < nroColumnas; col++ {
		if pos < len(message) {
			matriz[0][col] = rune(message[pos])
		} else {
			matriz[0][col] = 'X'
		}
		pos++
		util.ImprimirMatriz(matriz)
	}

	alfabeto := util.ObtenerAlfabetoEspanol()
	posclave := 0
	textoCifrado := strings.Builder{}
	for col := 0; col < nroColumnas; col++ {

		posicionTexto := util.ObtenerPosicion(alfabeto, string(matriz[0][col]))
		posicionClave := util.ObtenerPosicion(alfabeto, string(clave[posclave]))
		nuevaPosicion := 0
		if posicionClave > posicionTexto {
			nuevaPosicion = posicionTexto + len(alfabeto) - posicionClave
		} else {
			nuevaPosicion = posicionTexto - posicionClave
		}
		valorNuevo := string(alfabeto[nuevaPosicion%len(alfabeto)][0])

		fmt.Println(string(matriz[0][col]) + " : " + strconv.Itoa(posicionTexto) + " - " + string(clave[posclave]) + " : " + strconv.Itoa(posicionClave) + " = " + string(valorNuevo) + " : " + strconv.Itoa(nuevaPosicion))

		textoCifrado.WriteString(valorNuevo)

		if posclave < len(clave)-1 {
			posclave++
		} else {
			posclave = 0
		}
	}

	fmt.Println(textoCifrado.String())

	return textoCifrado.String()
}

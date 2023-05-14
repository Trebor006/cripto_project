package trasposicion

import (
	"criptograms/main/data"
	"criptograms/main/util"
	"math"
	"strings"
)

type ZigZag struct {
}

func (r ZigZag) Cypher(data data.Data) string {
	nroRails := data.NroRails
	message := strings.ReplaceAll(data.Message, " ", "")

	if nroRails <= 1 || nroRails >= len(message) {
		return "No se puede cifrar!!"
	}

	tamanoOla := 2*nroRails - 2
	nroOlas := int(math.Ceil(float64(len(message)) / float64(tamanoOla)))

	nroColumnas := tamanoOla * nroOlas

	matriz := util.InicializarMatriz(nroRails, nroColumnas)

	dirInv := true
	caracter := 'X'
	i := 0
	for j := 0; j < nroColumnas; j++ {
		valor := caracter
		if j < len(message) {
			valor = rune(message[j])
		}
		matriz[i][j] = valor

		if dirInv {
			i++
		} else {
			i--
		}

		if i == nroRails-1 || i == 0 {
			dirInv = !dirInv
		}

	}

	util.ImprimirMatriz(matriz)

	//obtener texto cifrado
	var textoCifrado strings.Builder
	for rail := 0; rail < nroRails; rail++ {
		for col := 0; col < nroColumnas; col++ {
			if !(matriz[rail][col] == 0) {
				textoCifrado.WriteRune(matriz[rail][col])
			}
		}
	}

	return textoCifrado.String()
}
func (r ZigZag) Decrypt(data data.Data) string {
	nroRails := data.NroRails
	encriptedMessage := strings.ReplaceAll(data.EncryptedMessage, " ", "")

	tamanoOla := 2*nroRails - 2
	nroOlas := int(math.Ceil(float64(len(encriptedMessage)) / float64(tamanoOla)))

	nroColumnas := tamanoOla * nroOlas

	matriz := util.InicializarMatriz(nroRails, nroColumnas)

	// Rellenar la matriz con el mensaje encriptado
	idx := 0
	for i := 0; i < nroRails; i++ {
		j := 0
		for j < nroColumnas && idx < len(encriptedMessage) {
			if (j/(nroRails-1))%2 == 0 { // moviéndose hacia abajo
				if i == j%(nroRails-1) {
					matriz[i][j] = rune(encriptedMessage[idx])
					idx++
				}
			} else { // moviéndose hacia arriba
				if i == nroRails-1-j%(nroRails-1) {
					matriz[i][j] = rune(encriptedMessage[idx])
					idx++
				}
			}
			j++
		}
	}

	util.ImprimirMatriz(matriz)

	// Leer la matriz en orden para obtener el mensaje descifrado
	var mensajeDescifrado strings.Builder
	for j := 0; j < nroColumnas; j++ {
		for i := 0; i < nroRails; i++ {
			if matriz[i][j] != 0 {
				mensajeDescifrado.WriteRune(matriz[i][j])
			}
		}
	}

	return mensajeDescifrado.String()
}

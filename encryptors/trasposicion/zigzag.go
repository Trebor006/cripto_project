package trasposicion

import (
	"criptograms/main/data"
	"criptograms/main/util"
	"fmt"
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
	nroOlas := util.Dividir(len(message), tamanoOla)

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
		util.ImprimirMatriz(matriz)
	}

	//obtener texto cifrado
	var textoCifrado strings.Builder
	for rail := 0; rail < nroRails; rail++ {
		for col := 0; col < nroColumnas; col++ {
			if !(matriz[rail][col] == 0) {
				textoCifrado.WriteRune(matriz[rail][col])
			}
		}

		fmt.Println(textoCifrado.String())
	}

	return textoCifrado.String()
}
func (r ZigZag) Decrypt(data data.Data) string {
	nroRails := data.NroRails
	encriptedMessage := strings.ReplaceAll(data.EncryptedMessage, " ", "")

	tamanoOla := 2*nroRails - 2
	nroOlas := util.Dividir(len(encriptedMessage), tamanoOla)

	nroColumnas := tamanoOla * nroOlas

	matriz := util.InicializarMatriz(nroRails, nroColumnas)

	// Rellenar los espacios de los rieles superior e inferior
	for i := 0; i < nroOlas; i++ {
		matriz[0][i*tamanoOla] = rune(encriptedMessage[i])
		util.ImprimirMatriz(matriz)
	}
	fmt.Println("FIN DE RELLENO DE CRESTAS")

	for i := 0; i < nroOlas; i++ {
		matriz[nroRails-1][i*tamanoOla+nroRails-1] = rune(encriptedMessage[len(encriptedMessage)-nroOlas+i])
		//idx++
		util.ImprimirMatriz(matriz)
	}

	fmt.Println("FIN DE RELLENO DE BASES")

	// Rellenar las rieles del medio..
	idx := nroOlas
	for i := 1; i < nroRails-1; i++ {
		j := 0
		for j < nroColumnas && idx < len(encriptedMessage) {
			if (j/(nroRails-1))%2 == 0 { // moviéndose hacia abajo
				if i == j%(nroRails-1) {
					matriz[i][j] = rune(encriptedMessage[idx])
					idx++

					util.ImprimirMatriz(matriz)
				}
			} else { // moviéndose hacia arriba
				if i == nroRails-1-j%(nroRails-1) {
					matriz[i][j] = rune(encriptedMessage[idx])
					idx++

					util.ImprimirMatriz(matriz)
				}
			}
			j++
		}
	}

	// Leer la matriz en orden para obtener el mensaje descifrado
	var mensajeDescifrado strings.Builder
	for j := 0; j < nroColumnas; j++ {
		for i := 0; i < nroRails; i++ {
			if matriz[i][j] != 0 {
				mensajeDescifrado.WriteRune(matriz[i][j])
			}
		}

		fmt.Println(mensajeDescifrado.String())
	}

	return mensajeDescifrado.String()
}

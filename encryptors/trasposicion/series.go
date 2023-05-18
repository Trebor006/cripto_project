package trasposicion

import (
	"criptograms/main/data"
	"criptograms/main/util"
	"fmt"
	"strings"
)

type Series struct {
}

func (r Series) Cypher(data data.Data) string {
	message := strings.ReplaceAll(strings.ToUpper(data.Message), " ", "")
	nuevoDetalle := generateDetail(message)
	imprimirArreglo(nuevoDetalle)

	textoCifrado := ""
	for _, val := range nuevoDetalle {
		textoCifrado += string(message[val-1])
	}

	return textoCifrado
}

func (r Series) Decrypt(data data.Data) string {
	message := strings.ReplaceAll(strings.ToUpper(data.EncryptedMessage), " ", "")
	nuevoDetalle := generateDetail(message)
	imprimirArreglo(nuevoDetalle)

	arregloOriginal := make([]string, len(nuevoDetalle))

	for i := 0; i < len(message); i++ {
		pos := nuevoDetalle[i] - 1
		arregloOriginal[pos] += string(message[i])
	}

	textoCifrado := strings.Join(arregloOriginal, "")

	return textoCifrado
}

func generateDetail(message string) []int {
	seriesDetalle := []int{}
	//insertar todos los primos
	for i := 1; i <= len(message); i++ {
		if util.EsPrimo(i) {
			seriesDetalle = append(seriesDetalle, i)
		}
	}

	//insertar todos los pares
	for i := 1; i <= len(message); i++ {
		if !existeNumero(seriesDetalle, i) && i%2 == 0 {
			seriesDetalle = append(seriesDetalle, i)
		}
	}

	for i := 1; i <= len(message); i++ {
		if !existeNumero(seriesDetalle, i) {
			seriesDetalle = append(seriesDetalle, i)
		}
	}

	return seriesDetalle
}

func existeNumero(array []int, numero int) bool {
	for _, elemento := range array {
		if elemento == numero {
			return true
		}
	}
	return false
}

func imprimirArreglo(arreglo []int) {
	fmt.Print("Nuevo [")
	for _, dig := range arreglo {
		fmt.Printf("%d ,", dig)
	}
	fmt.Print("]\n")

}

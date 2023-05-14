package util

import "fmt"

func ImprimirMatriz(matriz [][]rune) {
	fmt.Println()
	for fil := 0; fil < len(matriz); fil++ {
		for col := 0; col < len(matriz[fil]); col++ {
			if !(matriz[fil][col] == 0) {
				fmt.Printf("%c ", matriz[fil][col])
			} else {
				fmt.Printf("_ ")
			}
		}
		fmt.Println()
	}
}

func InicializarMatriz(nroFilas, nroColumnas int) [][]rune {
	matriz := make([][]rune, nroFilas)
	for fil := range matriz {
		matriz[fil] = make([]rune, nroColumnas)
	}

	return matriz
}

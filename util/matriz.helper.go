package util

import "fmt"

func ImprimirMatriz(matriz [][]rune) {
	fmt.Println()
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if !(matriz[i][j] == 0) {
				fmt.Printf("%c ", matriz[i][j])
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
}

func InicializarMatriz(nroFilas, nroColumnas int) [][]rune {
	matriz := make([][]rune, nroFilas)
	for i := range matriz {
		matriz[i] = make([]rune, nroColumnas)
		//for j := range matriz[i] {
		//	matriz[i][j] = '*'
		//}
	}

	return matriz
}

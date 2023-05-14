package trasposicion

import (
	"criptograms/main/data"
	"criptograms/main/util"
	"fmt"
	"sort"
	"strings"
)

type Columnas struct {
}

func (r Columnas) Cypher(data data.Data) string {
	var encryptedMessage = ""
	if data.NroColumns > 0 {
		fmt.Println("Cypher por nro Columnas")
		encryptedMessage = cypherByColumns(data)
	} else if len(data.Clave) > 0 {
		fmt.Println("Cypher por clave")
		encryptedMessage = cypherByColumnsWithKey(data)
	}

	return encryptedMessage
}

func (r Columnas) Decrypt(data data.Data) string {
	var decryptedMessage = ""
	if data.NroColumns > 0 {
		fmt.Println("Decrypt por nro Columnas")
		decryptedMessage = decrypt(data)
	} else if len(data.Clave) > 0 {
		fmt.Println("Decrypt por clave")
		decryptedMessage = decryptWithKey(data)
	}

	return decryptedMessage
}

func cypherByColumns(data data.Data) string {
	var message = data.Message
	var nroColumns = data.NroColumns
	message = strings.ReplaceAll(message, " ", "")
	if nroColumns <= 0 || nroColumns >= len(message) {
		return "Error no se puede cifrar!"
	}

	nroFilas := (len(message) + nroColumns - 1) / nroColumns

	//declarando la matriz de n x m
	matriz := util.InicializarMatriz(nroFilas, nroColumns)

	pos := 0
	for fil := 0; fil < nroFilas; fil++ {
		for col := 0; col < nroColumns; col++ {
			if pos <= len(message)-1 {
				matriz[fil][col] = rune(message[pos])
			} else {
				matriz[fil][col] = 'X'
			}
			pos++
		}

		util.ImprimirMatriz(matriz)
	}

	//obtener texto cifrado
	var textoCifrado strings.Builder
	for fil := 0; fil < nroColumns; fil++ {
		for col := 0; col < nroFilas; col++ {
			textoCifrado.WriteRune(matriz[col][fil])
		}

		fmt.Println(textoCifrado.String())
	}

	return textoCifrado.String()
}

func cypherByColumnsWithKey(data data.Data) string {
	var message = data.Message
	var clave = data.Clave

	clave = removeDuplicates(clave)
	nroColumns := len(clave)
	message = strings.ReplaceAll(message, " ", "")

	if nroColumns <= 0 || nroColumns >= len(message) {
		return message
	}

	//fmt.Println("length :: " + strconv.Itoa(len(message)))

	nroFilas := (len(message) + nroColumns - 1) / nroColumns
	//fmt.Println("nroFilas :: " + strconv.Itoa(nroFilas))

	//declarando la matriz de n x m
	matriz := util.InicializarMatriz(nroFilas, nroColumns)

	pos := 0
	for fil := 0; fil < nroFilas; fil++ {
		for col := 0; col < nroColumns; col++ {
			if pos <= len(message)-1 {
				matriz[fil][col] = rune(message[pos])
			} else {
				matriz[fil][col] = 'X'
			}
			pos++
		}

		util.ImprimirMatriz(matriz)
	}

	//obtener texto cifrado
	claveOrdenada := ordenarCaracteres(clave)

	var textoCifrado strings.Builder
	for _, char := range claveOrdenada {
		fil := strings.Index(clave, string(char))
		for col := 0; col < nroFilas; col++ {
			textoCifrado.WriteRune(matriz[col][fil])
		}

		fmt.Println(textoCifrado.String())
	}

	return textoCifrado.String()
}

func removeDuplicates(clave string) string {
	seen := make(map[rune]bool)
	builder := strings.Builder{}

	for _, char := range clave {
		if !seen[char] {
			seen[char] = true
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func ordenarCaracteres(clave string) string {
	caracteres := []rune(clave)

	sort.Slice(caracteres, func(i, j int) bool {
		return caracteres[i] < caracteres[j]
	})

	return string(caracteres)
}

//--------------------------------------------------------------------------------------------------------------

func decrypt(data data.Data) string {
	var message = data.EncryptedMessage
	var nroColumns = data.NroColumns
	//message = strings.ReplaceAll(message, " ", "")
	if nroColumns <= 0 || nroColumns >= len(message) {
		return "Error no se puede descifrar!"
	}

	nroFilas := util.Dividir(len(message), nroColumns)

	// Crear matriz de nroFilas x nroColumns y llenarla con caracteres de mensaje cifrado
	matriz := util.InicializarMatriz(nroFilas, nroColumns)

	pos := 0
	for col := 0; col < nroColumns; col++ {
		for fil := 0; fil < nroFilas; fil++ {
			matriz[fil][col] = rune(message[pos])
			pos++
		}

		util.ImprimirMatriz(matriz)
	}

	util.ImprimirMatriz(matriz)

	// Obtener mensaje original descifrando la matriz
	var mensajeOriginal strings.Builder
	for fil := 0; fil < nroFilas; fil++ {
		for col := 0; col < nroColumns; col++ {
			mensajeOriginal.WriteRune(matriz[fil][col])
		}

		fmt.Println(mensajeOriginal.String())
	}

	return mensajeOriginal.String()
}

func decryptWithKey(data data.Data) string {
	var message = data.EncryptedMessage
	var clave = removeDuplicates(data.Clave)
	var claveOrdenada = ordenarClave(clave)
	var nroColumns = len(clave)
	message = strings.ReplaceAll(message, " ", "")
	if nroColumns <= 0 || nroColumns >= len(message) {
		return "Error no se puede descifrar!"
	}

	nroFilas := util.Dividir(len(message), nroColumns)

	matriz := util.InicializarMatriz(nroFilas, nroColumns)

	pos := 0
	for col := 0; col < nroColumns; col++ {
		for fil := 0; fil < nroFilas; fil++ {
			matriz[fil][col] = rune(message[pos])
			pos++
		}

		util.ImprimirMatriz(matriz)
	}

	// Obtener mensaje original descifrando la matriz
	var mensajeOriginal strings.Builder
	for fil := 0; fil < nroFilas; fil++ {
		for _, char := range clave {
			col := strings.Index(claveOrdenada, string(char))
			mensajeOriginal.WriteRune(matriz[fil][col])
		}

		fmt.Println(mensajeOriginal.String())
	}

	return mensajeOriginal.String()
}

func ordenarClave(clave string) string {
	claveSlice := strings.Split(clave, "")
	sort.Strings(claveSlice)
	return strings.Join(claveSlice, "")
}

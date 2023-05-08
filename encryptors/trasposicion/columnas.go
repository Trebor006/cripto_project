package trasposicion

import (
	"criptograms/main/data"
	"sort"
	"strings"
)

type Columnas struct {
}

func (r Columnas) Cypher(data data.Data) string {
	var encryptedMessage = ""
	if data.NroColumns > 0 {
		encryptedMessage = cypherByColumns(data)
	} else if len(data.Clave) > 0 {
		encryptedMessage = cypherByColumnsWithKey(data)
	}

	return encryptedMessage
}

func (r Columnas) Decrypt(data data.Data) string {
	var decryptedMessage = ""
	if data.NroColumns > 0 {
		decryptedMessage = decrypt(data)
	} else if len(data.Clave) > 0 {
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
	matriz := make([][]rune, nroFilas)
	for i := range matriz {
		matriz[i] = make([]rune, nroColumns)
	}

	pos := 0
	for i := 0; i < nroFilas; i++ {
		for j := 0; j < nroColumns; j++ {
			if pos <= len(message)-1 {
				matriz[i][j] = rune(message[pos])
			} else {
				matriz[i][j] = 'X'
			}
			pos++
		}
	}

	//obtener texto cifrado
	var textoCifrado strings.Builder
	for i := 0; i < nroColumns; i++ {
		for j := 0; j < nroFilas; j++ {
			textoCifrado.WriteRune(matriz[j][i])
		}
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
	matriz := make([][]rune, nroFilas)
	for i := range matriz {
		matriz[i] = make([]rune, nroColumns)
	}

	pos := 0
	for i := 0; i < nroFilas; i++ {
		for j := 0; j < nroColumns; j++ {
			if pos <= len(message)-1 {
				matriz[i][j] = rune(message[pos])
			} else {
				matriz[i][j] = 'X'
			}
			pos++
		}
	}

	//obtener texto cifrado
	claveOrdenada := ordenarCaracteres(clave)

	var textoCifrado strings.Builder
	for _, b := range claveOrdenada {
		i := strings.Index(clave, string(b))
		for j := 0; j < nroFilas; j++ {
			textoCifrado.WriteRune(matriz[j][i])
		}
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

func decrypt(d data.Data) string {
	var message = d.EncryptedMessage
	var nroColumns = d.NroColumns
	//message = strings.ReplaceAll(message, " ", "")
	if nroColumns <= 0 || nroColumns >= len(message) {
		return "Error no se puede descifrar!"
	}

	nroFilas := (len(message) + nroColumns - 1) / nroColumns

	// Crear matriz de nroFilas x nroColumns y llenarla con caracteres de mensaje cifrado
	matriz := make([][]rune, nroFilas)
	for i := range matriz {
		matriz[i] = make([]rune, nroColumns)
	}

	pos := 0
	for i := 0; i < nroColumns; i++ {
		for j := 0; j < nroFilas; j++ {
			matriz[j][i] = rune(message[pos])
			pos++
		}
	}

	// Obtener mensaje original descifrando la matriz
	var mensajeOriginal strings.Builder
	for i := 0; i < nroFilas; i++ {
		for j := 0; j < nroColumns; j++ {
			if matriz[i][j] != 'X' {
				mensajeOriginal.WriteRune(matriz[i][j])
			}
		}
	}

	return mensajeOriginal.String()
}

func decryptWithKey(data data.Data) string {
	//todo PENDING
	return "PENDING!!!!"
}

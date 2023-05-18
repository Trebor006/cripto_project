package monoalfabetica

import (
	"cripto_project/main/data"
	"strings"
)

type TransformacionAfin struct {
}

func (r TransformacionAfin) Cypher(data data.Data) string {
	message := strings.ToLower(strings.ReplaceAll(data.Message, " ", ""))

	return cifrar(message, data.Clave)
}

func (r TransformacionAfin) Decrypt(data data.Data) string {
	message := strings.ToLower(strings.ReplaceAll(data.EncryptedMessage, " ", ""))

	return descifrar(message, data.Clave)
}

func cifrar(texto, clave string) string {
	if clave == "" {
		return "Error: La clave no puede ser una cadena vacía"
	}
	texto = strings.ToLower(texto)
	resultado := ""
	contador := 0
	for _, c := range texto {
		if c >= 'a' && c <= 'z' {
			c -= 'a'
			k := clave[contador] - 'a'
			caracterCifrado := (c + int32(k)) % 26
			caracterCifrado += 'a'
			resultado += string(caracterCifrado)
		} else {
			resultado += string(c)
		}

		contador++
		if contador >= len(clave) {
			contador = 0
		}
	}
	return strings.ToUpper(resultado)
}

func descifrar(textoCifrado, clave string) string {
	if clave == "" {
		return "Error: La clave no puede ser una cadena vacía"
	}
	textoCifrado = strings.ToLower(textoCifrado)
	resultado := ""
	contador := 0
	for _, c := range textoCifrado {
		if c >= 'a' && c <= 'z' {
			c -= 'a'
			k := clave[contador] - 'a'
			caracterDescifrado := (c - int32(k) + 26) % 26
			caracterDescifrado += 'a'
			resultado += string(caracterDescifrado)
		} else {
			resultado += string(c)
		}

		contador++
		if contador >= len(clave) {
			contador = 0
		}
	}

	return strings.ToUpper(resultado)
}

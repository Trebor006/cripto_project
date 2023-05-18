package monoalfabetica

import (
	"cripto_project/main/data"
	"strings"
)

type TransformacionAfin struct {
}

func (r TransformacionAfin) Cypher(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.Message, " ", ""))
	return cifradoAfin(message, data.NroDecimacion, data.Desplazamiento)
}

func (r TransformacionAfin) Decrypt(data data.Data) string {
	message := strings.ToUpper(strings.ReplaceAll(data.EncryptedMessage, " ", ""))
	return descifradoAfin(message, data.NroDecimacion, data.Desplazamiento)
}
func cifradoAfin(mensaje string, decimacion, desplazamiento int) string {
	mensajeCifrado := ""
	for _, caracter := range mensaje {
		if caracter >= 'A' && caracter <= 'Z' {
			// Obtener el valor numérico del carácter
			valor := int(caracter - 'A')
			// Aplicar la transformación afín
			valorCifrado := (decimacion*valor + desplazamiento) % 26
			// Convertir el valor cifrado a carácter
			caracterCifrado := rune(valorCifrado + 'A')
			mensajeCifrado += string(caracterCifrado)
		} else if caracter >= 'a' && caracter <= 'z' {
			// Obtener el valor numérico del carácter minúscula
			valor := int(caracter - 'a')
			// Aplicar la transformación afín
			valorCifrado := (decimacion*valor + desplazamiento) % 26
			// Convertir el valor cifrado a carácter minúscula
			caracterCifrado := rune(valorCifrado + 'a')
			mensajeCifrado += string(caracterCifrado)
		} else {
			mensajeCifrado += string(caracter)
		}
	}
	return mensajeCifrado
}

func descifradoAfin(mensajeCifrado string, decimacion, desplazamiento int) string {
	mensajeDescifrado := ""
	inversoDecimacion := 0
	for i := 0; i < 26; i++ {
		if (decimacion*i)%26 == 1 {
			inversoDecimacion = i
			break
		}
	}

	for _, caracter := range mensajeCifrado {
		if caracter >= 'A' && caracter <= 'Z' {
			// Obtener el valor numérico del carácter cifrado
			valorCifrado := int(caracter - 'A')
			// Aplicar la transformación afín inversa
			valor := (inversoDecimacion * (valorCifrado - desplazamiento + 26)) % 26
			// Convertir el valor descifrado a carácter
			caracterDescifrado := rune(valor + 'A')
			mensajeDescifrado += string(caracterDescifrado)
		} else if caracter >= 'a' && caracter <= 'z' {
			// Obtener el valor numérico del carácter cifrado minúscula
			valorCifrado := int(caracter - 'a')
			// Aplicar la transformación afín inversa
			valor := (inversoDecimacion * (valorCifrado - desplazamiento + 26)) % 26
			// Convertir el valor descifrado a carácter minúscula
			caracterDescifrado := rune(valor + 'a')
			mensajeDescifrado += string(caracterDescifrado)
		} else {
			mensajeDescifrado += string(caracter)
		}
	}
	return mensajeDescifrado
}

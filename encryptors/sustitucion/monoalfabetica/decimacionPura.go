package monoalfabetica

import (
	"cripto_project/main/data"
	"cripto_project/main/util"
)

type DecimacionPura struct {
}

func (r DecimacionPura) Cypher(data data.Data) string {
	palabra := util.LimpiarData(data.Message)
	decimacion := data.NroDecimacion

	alfabeto := util.ObtenerAlfabetoEspanol()
	alfabetoCifrado := make([]rune, len(alfabeto))

	for i := 0; i < len(alfabeto); i++ {
		alfabetoCifrado[i] = rune(alfabeto[(decimacion*i)%len(alfabeto)][0])
	}

	cifrado := ""
	for _, letra := range palabra {
		for i := 0; i < len(alfabeto); i++ {
			if rune(alfabeto[i][0]) == letra {
				cifrado += string(alfabetoCifrado[i])
			}
		}
	}

	return cifrado
}

func (r DecimacionPura) Decrypt(data data.Data) string {
	palabra := util.LimpiarData(data.EncryptedMessage)
	decimacion := data.NroDecimacion

	alfabeto := util.ObtenerAlfabetoEspanol()
	n := len(alfabeto)
	x := 1
	decifrado := ""

	for (decimacion*x)%n != 1 && x <= 50 {
		x++
	}

	if x == 50 {
		return "ERROR: LA DECIMACIÓN Y EL MÓDULO DE TRABAJO NO SON PRIMOS ENTRE SÍ"
	}

	for _, letra := range palabra {
		for i := 0; i < len(alfabeto); i++ {
			if rune(alfabeto[i][0]) == letra {
				aux := (x * i) % n
				decifrado += string(alfabeto[aux])
				break
			}
		}
	}

	return decifrado
}

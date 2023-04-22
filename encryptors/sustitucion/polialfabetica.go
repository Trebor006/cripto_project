package sustitucion

import "criptograms/main/data"

type Polialfabetica struct {
}

func (r Polialfabetica) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "Polialfabetica Encripted message : " + data.Message
}

func (r Polialfabetica) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "Polialfabetica Encripted message : " + data.Message
}

package desplazamiento

import "criptograms/main/data"

type Puro struct {
}

func (r Puro) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "Puro Encripted message : " + data.Message
}

func (r Puro) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "Puro Encripted message : " + data.Message
}

package trasposicion

import "criptograms/main/data"

type Filas struct {
}

func (r Filas) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "Filas Encripted message : " + data.Message
}

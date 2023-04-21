package trasposicion

import "criptograms/main/data"

type Grupos struct {
}

func (r Grupos) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "Grupos Encripted message : " + data.Message
}

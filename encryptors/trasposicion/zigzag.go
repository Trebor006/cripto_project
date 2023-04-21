package trasposicion

import "criptograms/main/data"

type ZigZag struct {
}

func (r ZigZag) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "ZigZag Encripted message : " + data.Message
}

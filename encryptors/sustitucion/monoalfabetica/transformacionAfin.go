package monoalfabetica

import "criptograms/main/data"

type TransformacionAfin struct {
}

func (r TransformacionAfin) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "TransformacionAfin Encripted message : " + data.Message
}

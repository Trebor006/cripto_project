package homofonos

import "criptograms/main/data"

type HomofonoPrimerOrden struct {
}

func (r HomofonoPrimerOrden) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "HomofonoPrimerOrden Encripted message : " + data.Message
}

func (r HomofonoPrimerOrden) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "HomofonoPrimerOrden Decrypted message : " + data.Message
}

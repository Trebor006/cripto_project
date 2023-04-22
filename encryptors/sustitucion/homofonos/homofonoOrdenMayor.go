package homofonos

import "criptograms/main/data"

type HomofonoOrdenMayor struct {
}

func (r HomofonoOrdenMayor) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "HomofonoOrdenMayor Encripted message : " + data.Message
}

func (r HomofonoOrdenMayor) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "HomofonoOrdenMayor Decrypted message : " + data.Message
}

package monoalfabetica

import "criptograms/main/data"

type DecimacionPura struct {
}

func (r DecimacionPura) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "DecimacionPura Encripted message : " + data.Message
}

func (r DecimacionPura) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "DecimacionPura Decrypted message : " + data.Message
}

package trasposicion

import "criptograms/main/data"

type Series struct {
}

func (r Series) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "Series Encripted message : " + data.Message
}

func (r Series) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "Series Decrypted message : " + data.Message
}

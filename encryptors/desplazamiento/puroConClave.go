package desplazamiento

import "criptograms/main/data"

type PuroConClave struct {
}

func (r PuroConClave) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "PuroConClave Encripted message : " + data.Message
}

func (r PuroConClave) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "PuroConClave Decrypted message : " + data.Message
}

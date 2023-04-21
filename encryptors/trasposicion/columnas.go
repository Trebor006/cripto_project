package trasposicion

import "criptograms/main/data"

type Columnas struct {
}

func (r Columnas) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "Columnas Encripted message : " + data.Message
}

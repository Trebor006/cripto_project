package sustitucion_monogramica

import "criptograms/main/data"

type PolialfabeticosPeriodicos struct {
}

func (r PolialfabeticosPeriodicos) Cypher(data data.Data) string {
	//todo implementar aqui la logica para cifrar!!!

	return "PolialfabeticosPeriodicos Encripted message : " + data.Message
}

func (r PolialfabeticosPeriodicos) Decrypt(data data.Data) string {
	//todo implementar aqui la logica para descifrar!!!

	return "PolialfabeticosPeriodicos Encripted message : " + data.Message
}

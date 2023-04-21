package main

import (
	"criptograms/main/data"
	"criptograms/main/encryptors"
	"criptograms/main/encryptors/desplazamiento"
	"criptograms/main/encryptors/sustitucion"
	"criptograms/main/encryptors/sustitucion/homofonos"
	"criptograms/main/encryptors/sustitucion/monoalfabetica"
	"criptograms/main/encryptors/sustitucion_monogramica"
	"criptograms/main/encryptors/trasposicion"
	"fmt"
)

func main() {

	//Aqui va estar toda la información general, Mensajes, claves, etc...
	var dataToCypher = data.Data{Message: "message"}

	var cypherPuro encryptors.EncryptorInterface = desplazamiento.Puro{}
	fmt.Println(cypherPuro.Cypher(dataToCypher))

	var cypherPuroConClave encryptors.EncryptorInterface = desplazamiento.PuroConClave{}
	fmt.Println(cypherPuroConClave.Cypher(dataToCypher))

	var trasposicionGrupos encryptors.EncryptorInterface = trasposicion.Grupos{}
	fmt.Println(trasposicionGrupos.Cypher(dataToCypher))

	var trasposicionSeries encryptors.EncryptorInterface = trasposicion.Series{}
	fmt.Println(trasposicionSeries.Cypher(dataToCypher))

	var trasposicionColumnas encryptors.EncryptorInterface = trasposicion.Columnas{}
	fmt.Println(trasposicionColumnas.Cypher(dataToCypher))

	var trasposicionFilas encryptors.EncryptorInterface = trasposicion.Filas{}
	fmt.Println(trasposicionFilas.Cypher(dataToCypher))

	var trasposicionZigZag encryptors.EncryptorInterface = trasposicion.ZigZag{}
	fmt.Println(trasposicionZigZag.Cypher(dataToCypher))

	var decimacionPura encryptors.EncryptorInterface = monoalfabetica.DecimacionPura{}
	fmt.Println(decimacionPura.Cypher(dataToCypher))

	var transformacionAfin encryptors.EncryptorInterface = monoalfabetica.TransformacionAfin{}
	fmt.Println(transformacionAfin.Cypher(dataToCypher))

	var poliAlfabetica encryptors.EncryptorInterface = sustitucion.Polialfabetica{}
	fmt.Println(poliAlfabetica.Cypher(dataToCypher))

	var homofonoPrimerOrden encryptors.EncryptorInterface = homofonos.HomofonoPrimerOrden{}
	fmt.Println(homofonoPrimerOrden.Cypher(dataToCypher))

	var homofonoOrdenMayor encryptors.EncryptorInterface = homofonos.HomofonoOrdenMayor{}
	fmt.Println(homofonoOrdenMayor.Cypher(dataToCypher))

	var poliAlfabeticosPeriodicos encryptors.EncryptorInterface = sustitucion_monogramica.PolialfabeticosPeriodicos{}
	fmt.Println(poliAlfabeticosPeriodicos.Cypher(dataToCypher))

}

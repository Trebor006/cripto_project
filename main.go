package main

import (
	"criptograms/main/data"
	"criptograms/main/encryptors"
	"criptograms/main/encryptors/trasposicion"
	"criptograms/main/util"
	"fmt"
)

func main() {

	//Aqui va estar toda la información general, Mensajes, claves, etc...
	//var dataToCypher = data.Data{Message: "message"}
	//var cypherPuro encryptors.EncryptorInterface = desplazamiento.Puro{}
	//fmt.Println(cypherPuro.Cypher(dataToCypher))
	//
	//var cypherPuroConClave encryptors.EncryptorInterface = desplazamiento.PuroConClave{}
	//fmt.Println(cypherPuroConClave.Cypher(dataToCypher))
	//
	//var trasposicionGrupos encryptors.EncryptorInterface = trasposicion.Grupos{}
	//fmt.Println(trasposicionGrupos.Cypher(dataToCypher))
	//
	//var trasposicionSeries encryptors.EncryptorInterface = trasposicion.Series{}
	//fmt.Println(trasposicionSeries.Cypher(dataToCypher))

	//todo Cifrado por Columnas
	//var dataToCypherByColumns = data.Data{Message: "ANDE YO CALIENTE Y RIASE LA GENTE", NroColumns: 8}
	//var formatter = util.Formatter{}
	//var trasposicionColumnas encryptors.EncryptorInterface = trasposicion.Columnas{}
	//var textoCifrado = trasposicionColumnas.Cypher(dataToCypherByColumns)
	//fmt.Println(formatter.Format(textoCifrado))

	//dataToCypherByColumns = data.Data{EncryptedMessage: textoCifrado, NroColumns: 8}
	//var textoDescifrado = trasposicionColumnas.Decrypt(dataToCypherByColumns)
	//fmt.Println(textoDescifrado)

	var dataToCypherByColumnsAndClave = data.Data{Message: "TRANSPOSICION POR COLUMNAS", Clave: "MARIO"}
	var formatter = util.Formatter{}
	var trasposicionColumnas encryptors.EncryptorInterface = trasposicion.Columnas{}
	var textoCifrado = trasposicionColumnas.Cypher(dataToCypherByColumnsAndClave)
	fmt.Println(formatter.Format(textoCifrado))

	var dataToCypherByColumnsWithKey = data.Data{EncryptedMessage: textoCifrado, Clave: "MARIO"}
	var textoDescifrado = trasposicionColumnas.Decrypt(dataToCypherByColumnsWithKey)
	fmt.Println(textoDescifrado)

	//todo Cifrado por Columnas

	////var trasposicionFilas encryptors.EncryptorInterface = trasposicion.Filas{}
	////fmt.Println(trasposicionFilas.Cypher(dataToCypher))
	//
	//var dataToCypherByZigZag = data.Data{Message: "message"}
	//var trasposicionZigZag encryptors.EncryptorInterface = trasposicion.ZigZag{}
	//fmt.Println(trasposicionZigZag.Cypher(dataToCypherByZigZag))
	//
	////var decimacionPura encryptors.EncryptorInterface = monoalfabetica.DecimacionPura{}
	////fmt.Println(decimacionPura.Cypher(dataToCypher))
	////
	////var transformacionAfin encryptors.EncryptorInterface = monoalfabetica.TransformacionAfin{}
	////fmt.Println(transformacionAfin.Cypher(dataToCypher))
	////
	////var poliAlfabetica encryptors.EncryptorInterface = sustitucion.Polialfabetica{}
	////fmt.Println(poliAlfabetica.Cypher(dataToCypher))
	//
	//var dataToCypherByPrimerOrden = data.Data{Message: "message"}
	//var homofonoPrimerOrden encryptors.EncryptorInterface = homofonos.HomofonoPrimerOrden{}
	//fmt.Println(homofonoPrimerOrden.Cypher(dataToCypherByPrimerOrden))
	//
	////var homofonoOrdenMayor encryptors.EncryptorInterface = homofonos.HomofonoOrdenMayor{}
	////fmt.Println(homofonoOrdenMayor.Cypher(dataToCypher))
	//
	//var dataToCypherByPolialfabeticoPeriodico = data.Data{Message: "message"}
	//var poliAlfabeticosPeriodicos encryptors.EncryptorInterface = sustitucion_monogramica.PolialfabeticosPeriodicos{}
	//fmt.Println(poliAlfabeticosPeriodicos.Cypher(dataToCypherByPolialfabeticoPeriodico))

}

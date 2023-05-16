package main

import (
	"criptograms/main/data"
	"criptograms/main/encryptors"
	"criptograms/main/encryptors/sustitucion/monoalfabetica"
	"criptograms/main/encryptors/trasposicion"
	"criptograms/main/util"
	"fmt"
)

func main() {

	//Aqui va estar toda la información general, Mensajes, claves, etc...
	//var cypherPuro = GetEncryptor(PURO)
	//var dataToCypher = data.Data{Message: "ABCDEFGHIJKLMN_OPQRSTUVWXYZ"}
	////var dataToCypher = data.Data{Message: "CESAR EL EMPERADOR HA SIDO ASESINADO"}
	//textoCifrado := cypherPuro.Cypher(dataToCypher)
	//fmt.Println(util.Format(textoCifrado))
	//
	//var dataToDecrypt = data.Data{EncryptedMessage: textoCifrado}
	////var dataToDecrypt = data.Data{EncryptedMessage: "ABCD"}
	//textoDecifrado := cypherPuro.Decrypt(dataToDecrypt)
	//fmt.Println(util.Format(textoDecifrado))

	//var cypherPuroConClave = GetEncryptor(PURO_CON_CLAVE)
	//var dataToCypherConClave = data.Data{Message: "ABCDEFGHIJKLMN_OPQRSTUVWXYZ", Clave: "MAGNMA", PosCifradoCesar: 3}
	//textoCifradoConClave := cypherPuroConClave.Cypher(dataToCypherConClave)
	//fmt.Println(util.Format(textoCifradoConClave))
	//
	//var dataToDecryptConClave = data.Data{EncryptedMessage: textoCifradoConClave, Clave: "MAGNMA", PosCifradoCesar: 3}
	//textoDecifradoConClave := cypherPuroConClave.Decrypt(dataToDecryptConClave)
	//fmt.Println(util.Format(textoDecifradoConClave))

	var trasposicionGrupos encryptors.EncryptorInterface = trasposicion.Grupos{}
	var CifradoGrupos = data.Data{Message: "MENSAJEORIGINALOP", Clave: "231"}
	var DesCifradoGrupos encryptors.EncryptorInterface = trasposicion.Grupos{}
	var DescifrarGrupos = data.Data{Message: "SAJEORMENIGINALIGINAL", Clave: "312"}
	fmt.Println(trasposicionGrupos.Cypher(CifradoGrupos))
	fmt.Println(DesCifradoGrupos.Decrypt(DescifrarGrupos))

	//var trasposicionSeries encryptors.EncryptorInterface = trasposicion.Series{}
	//fmt.Println(trasposicionSeries.Cypher(dataToCypher))

	//todo Cifrado por Columnas
	//var dataToCypherByColumns = data.Data{Message: "ANDE YO CALIENTE Y RIASE LA GENTE", NroColumns: 8}
	//var dataToCypherByColumns = data.Data{Message: "ANDE YO CALIENTE Y RIASE LA GENTE", NroColumns: 8}
	//var dataToCypherByColumns = data.Data{Message: "EL MALWARE TAMBIEN CONOCIDO COMO CODIGO MALIGNO O CODIGO MALICIOSO ES UN TIPO DE SOFTWARE QUE TIENE COMO OBJETIVO INFILTRARSE O DANAR UN SISTEMA INFORMATICO", NroColumns: 9}
	//var trasposicionColumnas1 encryptors.EncryptorInterface = trasposicion.Columnas{}
	//var textoCifrado1 = trasposicionColumnas1.Cypher(dataToCypherByColumns)
	//fmt.Println(util.Format(textoCifrado1))
	//
	//dataToCypherByColumns = data.Data{EncryptedMessage: textoCifrado1, NroColumns: 9}
	//var textoDescifrado = trasposicionColumnas1.Decrypt(dataToCypherByColumns)
	//fmt.Println(util.Format(textoDescifrado))

	//var dataToCypherByColumnsAndClave = data.Data{Message: "TRANSPOSICION POR COLUMNAS", Clave: "MARIO"}
	//var dataToCypherByColumnsAndClave = data.Data{Message: "EL MALWARE TAMBIEN CONOCIDO COMO CODIGO MALIGNO O CODIGO MALICIOSO ES UN TIPO DE SOFTWARE QUE TIENE COMO OBJETIVO INFILTRARSE O DANAR UN SISTEMA INFORMATICO", Clave: "MALIGNO"}
	//var trasposicionColumnas encryptors.EncryptorInterface = trasposicion.Columnas{}
	//var textoCifradoPorColumnasConClave = trasposicionColumnas.Cypher(dataToCypherByColumnsAndClave)
	//fmt.Println(util.Format(textoCifradoPorColumnasConClave))
	//
	////var dataToCypherByColumnsWithKey = data.Data{EncryptedMessage: textoCifradoPorColumnasConClave, Clave: "MARIO"}
	//var dataToCypherByColumnsWithKey = data.Data{EncryptedMessage: textoCifradoPorColumnasConClave, Clave: "MALIGNO"}
	//var textoDescifradoPorColumnasConClave = trasposicionColumnas.Decrypt(dataToCypherByColumnsWithKey)
	//fmt.Println(textoDescifradoPorColumnasConClave)

	//todo Cifrado por Columnas

	////var trasposicionFilas encryptors.EncryptorInterface = trasposicion.Filas{}
	////fmt.Println(trasposicionFilas.Cypher(dataToCypher))
	//
	//var dataToCypherByZigZag = data.Data{Message: "VISTEME DESPACIO QUE TENGO PRISA", NroRails: 4}
	////var dataToCypherByZigZag = data.Data{Message: "EL SISTEMA RAIL FENCE SE UTILIZO EN LA GUERRA DE SECESION", NroRails: 5}
	////var dataToCypherByZigZag = data.Data{Message: "HOLA MUNDO", NroRails: 4}
	//var trasposicionZigZag encryptors.EncryptorInterface = trasposicion.ZigZag{}
	//var textoCifradoZigZag = trasposicionZigZag.Cypher(dataToCypherByZigZag)
	//fmt.Println(util.Format(textoCifradoZigZag))
	//
	//var dataToDecryptByZigZag = data.Data{EncryptedMessage: textoCifradoZigZag, NroRails: 4}
	////var dataToDecryptByZigZag = data.Data{EncryptedMessage: "EERCO TNRLU ASDVE VEONE SUAOB INUZL TBZTR IVITR YSALR CILUC AYSAZ AALOG NRRTB ATHVL REONE SUTLE EAOEA ROT", NroRails: 5}
	//var textoDescifradoZigZag = trasposicionZigZag.Decrypt(dataToDecryptByZigZag)
	//fmt.Println(util.Format(textoDescifradoZigZag))

	////var decimacionPura encryptors.EncryptorInterface = monoalfabetica.DecimacionPura{}
	////fmt.Println(decimacionPura.Cypher(dataToCypher))
	////
	var transformacionAfin encryptors.EncryptorInterface = monoalfabetica.TransformacionAfin{}
	var dataTransformacionAfin = data.Data{Message: "HOLA COMO ESTAS", Clave: "MARIO"}
	var DescifradoTranformacionAfin encryptors.EncryptorInterface = monoalfabetica.TransformacionAfin{}
	var textocifradoportransformacionafin = (transformacionAfin.Cypher(dataTransformacionAfin))
	var dataDescifradoTransformacionAfin = data.Data{Message: textocifradoportransformacionafin, Clave: "MARIO"}
	fmt.Println(util.Format(textocifradoportransformacionafin))
	fmt.Println(DescifradoTranformacionAfin.Decrypt(dataDescifradoTransformacionAfin))
	////
	////var poliAlfabetica encryptors.EncryptorInterface = sustitucion.Polialfabetica{}
	////fmt.Println(poliAlfabetica.Cypher(dataToCypher))
	//
	//
	//random := util.ObtenerGeneradorRandom()
	//diccionario := util.InicializarDiccionario(100, random)
	//var homofonoPrimerOrden encryptors.EncryptorInterface = homofonos.HomofonoPrimerOrden{}
	//
	//var dataToCypherByPrimerOrden = data.Data{Message: "HOLA MUNDO QUE TAL QUE HACIENDO", Diccionario: diccionario, RandomGenerator: random}
	//var encryptedMessageByPrimerOrden = homofonoPrimerOrden.Cypher(dataToCypherByPrimerOrden)
	//var encryptedMessageByPrimerOrden2 = homofonoPrimerOrden.Cypher(dataToCypherByPrimerOrden)
	//fmt.Println(encryptedMessageByPrimerOrden)
	//fmt.Println(encryptedMessageByPrimerOrden2)
	//
	//var dataToDecryptByPrimerOrden = data.Data{EncryptedMessage: encryptedMessageByPrimerOrden, Diccionario: diccionario, RandomGenerator: random}
	//fmt.Println(homofonoPrimerOrden.Decrypt(dataToDecryptByPrimerOrden))
	//
	//var dataToDecryptByPrimerOrden2 = data.Data{EncryptedMessage: encryptedMessageByPrimerOrden2, Diccionario: diccionario, RandomGenerator: random}
	//fmt.Println(homofonoPrimerOrden.Decrypt(dataToDecryptByPrimerOrden2))

	////var homofonoOrdenMayor encryptors.EncryptorInterface = homofonos.HomofonoOrdenMayor{}
	////fmt.Println(homofonoOrdenMayor.Cypher(dataToCypher))
	//
	//var dataToCypherByPolialfabeticoPeriodico = data.Data{Message: "HOLA AMIGOS", Clave: "CIFRA"}
	////var dataToCypherByPolialfabeticoPeriodico = data.Data{Message: "eres como una mariposa vuelas y te posas", Clave: "MANA"}
	//var poliAlfabeticosPeriodicos encryptors.EncryptorInterface = GetEncryptor(POLIALFABETICO_PERIODICO)
	//var textoCifradoVigenere = poliAlfabeticosPeriodicos.Cypher(dataToCypherByPolialfabeticoPeriodico)
	//fmt.Println(util.Format(textoCifradoVigenere))
	//
	//var dataToDecryptByPolialfabeticoPeriodico = data.Data{Message: textoCifradoVigenere, Clave: "CIFRA"}
	//var textoDesencriptadoVigenere = poliAlfabeticosPeriodicos.Decrypt(dataToDecryptByPolialfabeticoPeriodico)
	//fmt.Println(util.Format(textoDesencriptadoVigenere))

}

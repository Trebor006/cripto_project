package main

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
	//var dataToCypherByZigZag = data.Data{Message: "EL SISTEMA RAIL FENCE SE UTILIZO EN LA GUERRA DE SECESION", NroRails: 5}
	//////var dataToCypherByZigZag = data.Data{Message: "HOLA MUNDO", NroRails: 4}
	//var trasposicionZigZag encryptors.EncryptorInterface = trasposicion.ZigZag{}
	//var textoCifradoZigZag = trasposicionZigZag.Cypher(dataToCypherByZigZag)
	//fmt.Println(util.Format(textoCifradoZigZag))
	//
	//var dataToDecryptByZigZag = data.Data{EncryptedMessage: textoCifradoZigZag, NroRails: 5}
	////var dataToDecryptByZigZag = data.Data{EncryptedMessage: "EERCO TNRLU ASDVE VEONE SUAOB INUZL TBZTR IVITR YSALR CILUC AYSAZ AALOG NRRTB ATHVL REONE SUTLE EAOEA ROT", NroRails: 5}
	//var textoDescifradoZigZag = trasposicionZigZag.Decrypt(dataToDecryptByZigZag)
	//fmt.Println(util.Format(textoDescifradoZigZag))

	////var decimacionPura encryptors.EncryptorInterface = monoalfabetica.DecimacionPura{}
	////fmt.Println(decimacionPura.Cypher(dataToCypher))
	////
	////var transformacionAfin encryptors.EncryptorInterface = monoalfabetica.TransformacionAfin{}
	////fmt.Println(transformacionAfin.Cypher(dataToCypher))
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

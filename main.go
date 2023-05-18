package main

func main() {

	//Aqui va estar toda la información general, Mensajes, claves, etc...
	//var cypherPuro = encryptors.GetEncryptor(encryptors.PURO)
	////var dataToCypherSeries = data.Data{Message: "ABCDEFGHIJKLMN_OPQRSTUVWXYZ"}
	//var dataToCypherSeries = data.Data{Message: "CESAR EL EMPERADOR HA SIDO ASESINADO"}
	//textoCifrado := cypherPuro.Cypher(dataToCypherSeries)
	//fmt.Println(util.Format(textoCifrado))
	//
	//var dataToDecrypt = data.Data{EncryptedMessage: textoCifrado}
	////var dataToDecrypt = data.Data{EncryptedMessage: "ABCD"}
	//textoDecifrado := cypherPuro.Decrypt(dataToDecrypt)
	//fmt.Println(util.Format(textoDecifrado))

	//var cypherPuroConClave = encryptors.GetEncryptor(encryptors.PURO_CON_CLAVE)
	//var dataToCypherConClave = data.Data{Message: "ABCDEFGHIJKLMN_OPQRSTUVWXYZ", Clave: "MAGNMA"}
	//textoCifradoConClave := cypherPuroConClave.Cypher(dataToCypherConClave)
	//fmt.Println(util.Format(textoCifradoConClave))
	//
	//var dataToDecryptConClave = data.Data{EncryptedMessage: textoCifradoConClave, Clave: "MAGNMA"}
	//textoDecifradoConClave := cypherPuroConClave.Decrypt(dataToDecryptConClave)
	//fmt.Println(util.Format(textoDecifradoConClave))
	//
	//cifradoGrupos := encryptors.GetEncryptor(encryptors.GRUPOS)
	//var CifradoGrupos = data.Data{Message: "UN SECRETO ALGO EXTRANO", Clave: "13572468"}
	//textoEncriptado := cifradoGrupos.Cypher(CifradoGrupos)
	//fmt.Println(util.Format(textoEncriptado))
	//
	//var DescifrarGrupos = data.Data{EncryptedMessage: textoEncriptado, Clave: "13572468"}
	//textoDesencriptado := cifradoGrupos.Decrypt(DescifrarGrupos)
	//fmt.Println(util.Format(textoDesencriptado))
	//
	//
	//var trasposicionSeries = encryptors.GetEncryptor(encryptors.SERIES)
	//var dataToCypherSeries = data.Data{Message: "AHORA CIFRAMOS POR SERIES"}
	//textoCifradoPorSeries := trasposicionSeries.Cypher(dataToCypherSeries)
	//fmt.Println(util.Format(textoCifradoPorSeries))
	//
	//var dataToDecriptSeries = data.Data{EncryptedMessage: textoCifradoPorSeries}
	//textoDescifradoPorSeries := trasposicionSeries.Decrypt(dataToDecriptSeries)
	//fmt.Println(util.Format(textoDescifradoPorSeries))

	//todo Cifrado por Columnas
	//var dataToCypherByColumns = data.Data{Message: "ANDE YO CALIENTE Y RIASE LA GENTE", NroColumns: 8}
	////var dataToCypherByColumns = data.Data{Message: "ANDE YO CALIENTE Y RIASE LA GENTE", NroColumns: 8}
	////var dataToCypherByColumns = data.Data{Message: "EL MALWARE TAMBIEN CONOCIDO COMO CODIGO MALIGNO O CODIGO MALICIOSO ES UN TIPO DE SOFTWARE QUE TIENE COMO OBJETIVO INFILTRARSE O DANAR UN SISTEMA INFORMATICO", NroColumns: 9}
	//var trasposicionColumnas1 encryptors.EncryptorInterface = trasposicion.Columnas{}
	//var textoCifrado1 = trasposicionColumnas1.Cypher(dataToCypherByColumns)
	//fmt.Println(util.Format(textoCifrado1))
	////
	//dataToCypherByColumns = data.Data{EncryptedMessage: textoCifrado1, NroColumns: 8}
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

	//dataFilas := data.Data{Message: "TRANSPOSICION POR FILAS", NroFilas: 4}
	//trasposicionFilas := encryptors.GetEncryptor(encryptors.FILAS)
	//resultado := trasposicionFilas.Cypher(dataFilas)
	//fmt.Println(util.Format(resultado))
	//
	//dataFilasDecifrado := data.Data{EncryptedMessage: resultado, NroFilas: 4}
	//resultadoDescifrado := trasposicionFilas.Decrypt(dataFilasDecifrado)
	//fmt.Println(util.Format(resultadoDescifrado))

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

	//decimacionPura := encryptors.GetEncryptor(encryptors.DECIMACION_PURA)
	//var dataToCypherDecimacionPura = data.Data{Message: "DILE A LAURA QUE LA QUIERO", NroDecimacion: 20}
	//textoCifradoPorDecimacionPura := decimacionPura.Cypher(dataToCypherDecimacionPura)
	//fmt.Println(util.Format(textoCifradoPorDecimacionPura))
	//
	////var dataToDecryptDecimacionPura = data.Data{EncryptedMessage: textoCifradoPorDecimacionPura, NroDecimacion: 20}
	//var dataToDecryptDecimacionPura = data.Data{EncryptedMessage: "M_ZHW DHBGR HZZAU DHRG", NroDecimacion: 22}
	//textoDesencriptadoPorDecimacionPura := decimacionPura.Decrypt(dataToDecryptDecimacionPura)
	//fmt.Println(util.Format(textoDesencriptadoPorDecimacionPura))
	////
	//var transformacionAfin encryptors.EncryptorInterface = monoalfabetica.TransformacionAfin{}
	//var dataTransformacionAfin = data.Data{Message: "HOLA MUNDO", NroDecimacion: 5, Desplazamiento: 8}
	//var DescifradoTranformacionAfin encryptors.EncryptorInterface = monoalfabetica.TransformacionAfin{}
	//var textocifradoportransformacionafin = (transformacionAfin.Cypher(dataTransformacionAfin))
	//var dataDescifradoTransformacionAfin = data.Data{Message: textocifradoportransformacionafin, NroDecimacion: 5, Desplazamiento: 8}
	//fmt.Println(util.Format(textocifradoportransformacionafin))
	//fmt.Println(DescifradoTranformacionAfin.Decrypt(dataDescifradoTransformacionAfin))
	//
	//var poliAlfabetica encryptors.EncryptorInterface = sustitucion.Polialfabetica{}
	//fmt.Println(poliAlfabetica.Cypher(dataToCypherSeries))
	//
	//
	//random := util.ObtenerGeneradorRandom()
	//diccionario := util.InicializarDiccionarioPrimerOrden(100, random)
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

	//var homofonoOrdenMayor = encryptors.GetEncryptor(encryptors.ORDEN_MAYOR)
	//
	//random := util.ObtenerGeneradorRandom()
	//diccionario := util.InicializarDiccionarioOrdenMayor(5000, random)
	//var dataToCypherOrdenMayor = data.Data{Message: "DILE A LAURA QUE LA QUIERO", Diccionario: diccionario, RandomGenerator: random}
	//textoCifradoPorOrdenMayor := homofonoOrdenMayor.Cypher(dataToCypherOrdenMayor)
	//fmt.Println(util.Format(textoCifradoPorOrdenMayor))
	//
	//var dataToDecryptOrdenMayor = data.Data{EncryptedMessage: textoCifradoPorOrdenMayor, Diccionario: diccionario, RandomGenerator: random}
	//textoDescifradoPorOrdenMayor := homofonoOrdenMayor.Decrypt(dataToDecryptOrdenMayor)
	//fmt.Println(util.Format(textoDescifradoPorOrdenMayor))

	//var dataToCypherByPolialfabeticoPeriodico = data.Data{Message: "HOLA AMIGOS", Clave: "CIFRA"}
	////var dataToCypherByPolialfabeticoPeriodico = data.Data{Message: "eres como una mariposa vuelas y te posas", Clave: "MANA"}
	//var poliAlfabeticosPeriodicos encryptors.EncryptorInterface = encryptors.GetEncryptor(encryptors.POLIALFABETICO_PERIODICO)
	//var textoCifradoVigenere = poliAlfabeticosPeriodicos.Cypher(dataToCypherByPolialfabeticoPeriodico)
	//fmt.Println(util.Format(textoCifradoVigenere))
	//
	//var dataToDecryptByPolialfabeticoPeriodico = data.Data{EncryptedMessage: textoCifradoVigenere, Clave: "CIFRA"}
	//var textoDesencriptadoVigenere = poliAlfabeticosPeriodicos.Decrypt(dataToDecryptByPolialfabeticoPeriodico)
	//fmt.Println(util.Format(textoDesencriptadoVigenere))

}

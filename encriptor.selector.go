package main

import (
	"criptograms/main/encryptors"
	"criptograms/main/encryptors/desplazamiento"
	"criptograms/main/encryptors/sustitucion/homofonos"
	"criptograms/main/encryptors/sustitucion/monoalfabetica"
	"criptograms/main/encryptors/sustitucion_monogramica"
	"criptograms/main/encryptors/trasposicion"
	"fmt"
)

// Definir un tipo base para el enum
type MetodosEncriptado int

// Definir las constantes que representan los valores del enum
const (
	PURO                     MetodosEncriptado = iota
	PURO_CON_CLAVE           MetodosEncriptado = iota
	ORDEN_MAYOR              MetodosEncriptado = iota
	PRIMER_ORDEN             MetodosEncriptado = iota
	DECIMACION_PURA          MetodosEncriptado = iota
	TRANSFORMACION_AFIN      MetodosEncriptado = iota
	POLIALFABETICO_PERIODICO MetodosEncriptado = iota
	COLUMNAS                 MetodosEncriptado = iota
	FILAS                    MetodosEncriptado = iota
	GRUPOS                   MetodosEncriptado = iota
	SERIES                   MetodosEncriptado = iota
	ZIG_ZAG                  MetodosEncriptado = iota
)

func GetEncryptor(metodoEncriptado MetodosEncriptado) encryptors.EncryptorInterface {
	switch metodoEncriptado {
	case PURO:
		return desplazamiento.Puro{}
	case PURO_CON_CLAVE:
		return desplazamiento.PuroConClave{}
	case ORDEN_MAYOR:
		return homofonos.HomofonoOrdenMayor{}
	case PRIMER_ORDEN:
		return homofonos.HomofonoPrimerOrden{}
	case DECIMACION_PURA:
		return monoalfabetica.DecimacionPura{}
	case TRANSFORMACION_AFIN:
		return monoalfabetica.TransformacionAfin{}
	case POLIALFABETICO_PERIODICO:
		return sustitucion_monogramica.PolialfabeticosPeriodicos{}
	case COLUMNAS:
		return trasposicion.Columnas{}
	case FILAS:
		return trasposicion.Filas{}
	case GRUPOS:
		return trasposicion.Grupos{}
	case SERIES:
		return trasposicion.Series{}
	case ZIG_ZAG:
		return trasposicion.ZigZag{}

	default:
		fmt.Println("Método de encriptado desconocido")
		return nil
	}
}

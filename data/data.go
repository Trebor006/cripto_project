package data

import "math/rand"

type Data struct {
	Message          string // En caso de que quieran encriptar un mensaje
	Clave            string // Se utilizaria en caso de que tengan metodos con clave de cifrado
	EncryptedMessage string //En caso de que quieran desencriptar un mensaje tendrian que setearlo aqui
	NroColumns       int    // Para Cifrado por Columnas
	NroFilas         int    // Para Cifrado por Fila
	NroRails         int    // Para Cifrado por Rails o ZigZag
	NroDecimacion    int    //Transformacion Afin
	//Desplazamiento   int    //Transformacion Afin

	Diccionario     map[string][]int //Para Cifrado de Primer Orden
	RandomGenerator *rand.Rand
}

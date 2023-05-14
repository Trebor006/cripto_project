package data

type Data struct {
	Message          string // En caso de que quieran encriptar un mensaje
	Clave            string // Se utilizaria en caso de que tengan metodos con clave de cifrado
	EncryptedMessage string //En caso de que quieran desencriptar un mensaje tendrian que setearlo aqui
	NroColumns       int    // Para Cifrado por Columnas
	NroRails         int    // Para Cifrado por Rails o ZigZag
}

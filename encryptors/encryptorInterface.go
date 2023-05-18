package encryptors

import "criptograms/main/data"

type EncryptorInterface interface {
	Cypher(data data.Data) string
	Decrypt(data data.Data) string
}

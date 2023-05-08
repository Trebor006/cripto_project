package util

type Formatter struct {
}

func (f Formatter) Format(text string) string {
	textoFormateado := ""

	for i, char := range text {
		if i%5 == 0 && i != 0 {
			textoFormateado += " "
		}
		textoFormateado += string(char)
	}

	return textoFormateado
}

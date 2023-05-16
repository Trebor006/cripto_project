package util

func ObtenerAlfabetoEspanol() []string {
	return []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "_", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
}

func ObtenerPosicion(alfabeto []string, r string) int {
	for pos, char := range alfabeto {
		if char == r {
			return pos
		}
	}

	return -1
}

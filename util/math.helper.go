package util

import "math"

func Dividir(dividendo, divisor int) int {
	return int(math.Ceil(float64(dividendo) / float64(divisor)))
}

func EsPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

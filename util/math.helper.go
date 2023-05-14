package util

import "math"

func Dividir(dividendo, divisor int) int {
	return int(math.Ceil(float64(dividendo) / float64(divisor)))
}

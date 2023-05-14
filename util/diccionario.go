package util

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func InicializarDiccionario(longitudDiccionario int, r *rand.Rand) map[string][]int {
	diccionario := map[string][]int{}

	// Agregar elementos al slice
	alfabeto := ObtenerAlfabetoEspanol()

	if longitudDiccionario < len(alfabeto) {
		fmt.Println("Error al cargar el diccionario, longitud minima no establecida")
		return nil
	}

	generados := map[int]bool{}

	for _, char := range alfabeto {
		valorAsignado := false
		for !valorAsignado {
			numero := generarNumeroRandom(10*longitudDiccionario/4, r)
			if !generados[numero] {
				diccionario[char] = append(diccionario[char], numero)
				generados[numero] = true
				valorAsignado = true
			}
		}
	}

	for i := len(alfabeto); i < longitudDiccionario-1; i++ {
		valorAsignado := false
		for !valorAsignado {
			numero := generarNumeroRandom(10*longitudDiccionario/4, r)
			pos := generarNumeroRandom(len(alfabeto), r)
			if !generados[numero] {
				diccionario[alfabeto[pos-1]] = append(diccionario[alfabeto[pos-1]], numero)
				generados[numero] = true
				valorAsignado = true
			}
		}
	}

	ImprimirDiccionario(diccionario)

	return diccionario
}

func ImprimirDiccionario(diccionario map[string][]int) {
	// Obtener las claves en orden alfabético
	claves := make([]string, 0, len(diccionario))
	for clave := range diccionario {
		claves = append(claves, clave)
	}
	sort.Strings(claves)

	for _, clave := range claves {
		fmt.Printf("%s => ", clave)
		for _, v := range diccionario[clave] {
			fmt.Printf("%d, ", v)
		}
		fmt.Println("")
	}
}

func generarNumeroRandom(longitudDiccionario int, r *rand.Rand) int {
	return r.Intn(longitudDiccionario) + 1
}

func ObtenerGeneradorRandom() *rand.Rand {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r
}

func GenerarNumeroRandom(length int, r *rand.Rand) int {
	return r.Intn(length)
}

package util

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func InicializarDiccionarioPrimerOrden(longitudDiccionario int, r *rand.Rand) map[string][]int {
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

func InicializarDiccionarioOrdenMayor(longitudDiccionario int, r *rand.Rand) map[string][]int {
	claves := generarClaves()
	diccionario := map[string][]int{}

	numerosAsignados := map[int]bool{}

	for _, valor := range claves {
		posibilidadesPorValor := generarNumeroRandom(10, r) + 2
		contadorPoibilidadesAsignadas := 0

		diccionario[valor] = []int{}

		for contadorPoibilidadesAsignadas < posibilidadesPorValor {

			valorAsignado := false
			for !valorAsignado {
				numero := generarNumeroRandom(5*longitudDiccionario/4, r)

				if !numerosAsignados[numero] {
					diccionario[valor] = append(diccionario[valor], numero)
					numerosAsignados[numero] = true
					valorAsignado = true
					contadorPoibilidadesAsignadas++
				}
			}
		}
	}
	ImprimirDiccionario(diccionario)

	return diccionario
}

func generarClaves() []string {
	alfabeto := ObtenerAlfabetoEspanol()
	claves := make([]string, len(alfabeto)*len(alfabeto))
	index := 0

	for _, letra1 := range alfabeto {
		for _, letra2 := range alfabeto {
			clave := letra1 + letra2
			claves[index] = clave
			index++
		}
	}

	return claves
}

func existeNumero(arr []int, valor int) bool {
	for _, num := range arr {
		if num == valor {
			return true
		}
	}
	return false
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

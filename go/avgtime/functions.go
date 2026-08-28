package avgtime

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Crea un array con numeros aleatorios (pensados para ser insetados o buscados en el arbol)
func RandomNumbersArray(lenght int) []int {
	array := make([]int, 0, lenght)
	for range lenght {
		array = append(array, rand.IntN(10000000))
	}
	return array
}

// Ejecuta n test y calcula su promedio (PARA TEMINAR)
func AverageTime(threads int, test func() time.Duration) {
	var testTime time.Duration
	var totalTime time.Duration
	var minTime time.Duration
	var maxTime time.Duration
	runs := 15

	for i := range runs {
		testTime = test()
		if i == 0 {
			minTime = testTime
		}
		if testTime > maxTime {
			maxTime = testTime
		}
		if testTime < minTime {
			minTime = testTime
		}
		totalTime += testTime
	}
	avTime := totalTime / time.Duration(runs)
	times := []time.Duration{avTime, minTime, maxTime}
	fmt.Println("Promedio con", threads, "hilos:", times)
}

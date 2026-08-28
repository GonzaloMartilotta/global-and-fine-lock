package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Inserta una lista de numeros dentro el arbol
func (t *tree) fillTree(lenght int) {
	elements := randomNumbersArray(lenght)
	for i := range len(elements) {
		t.Add(elements[i])
	}
}

// Inserta una lista de numeros dentro el arbol simultaneamente ocn varios hilos
func (t *tree) fillTreeGo(workers int, lenght int) {
	elements := randomNumbersArray(lenght)
	cycles := len(elements) / workers
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func(cycles int) {
			defer wg.Done()
			for i := range cycles {
				t.Add(elements[i])
			}
		}(cycles)
	}
	wg.Wait()
}

// Crea un array con numeros aleatorios (pensados para ser insetados o buscados en el arbol)
func randomNumbersArray(lenght int) []int {
	array := make([]int, 0, lenght)
	for range lenght {
		array = append(array, rand.IntN(10000000))
	}
	return array
}

// Test sin hilos
func (t *tree) searchNum(searchArray []int) {
	for i := range searchArray {
		t.Search(searchArray[i])
	}
}

// PRE: El arbol tiene que tener elementos
func (t *tree) searchNumGo(workers int, searchArray []int) {
	searchBlockLenght := len(searchArray) / workers
	var wg sync.WaitGroup

	for i := range workers {
		wg.Add(1)
		indexStart := i * searchBlockLenght
		indexEnd := indexStart + searchBlockLenght
		if i == workers-1 {
			indexEnd = len(searchArray)
		}
		searchBlock := searchArray[indexStart:indexEnd]

		//fmt.Println("Hilo", i, "busca este rango", "[", indexStart, indexEnd, "]")

		go func(block []int) {
			defer wg.Done()
			for j := range block {
				t.Search(block[j])
			}
		}(searchBlock)
	}
	wg.Wait()
}

// Ejecuta n test y calcula su promedio (PARA TEMINAR)
func averageTime(threads int, test func() time.Duration) {
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

func (t *tree) addTest(lenght int, workers int) {
	testTime := time.Now()
	t.fillTreeGo(workers, lenght)
	fmt.Println(lenght, "elementos insertados en el arbol en", time.Since(testTime))
	fmt.Println()

	t.root = nil

	testTime2 := time.Now()
	t.fillTree(lenght)
	fmt.Println(lenght, "elementos insertados en el arbol en", time.Since(testTime2))

	fmt.Println("Arbol y array creados, empezando el test...")
}

func main() {
	t := new(tree)
	t.fillTreeGo(4, 500)
	menuTesting(t)
}

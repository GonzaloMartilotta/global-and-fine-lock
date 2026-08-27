package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Inserta una lista de numeros dentro el arbol
func (t *tree) fillTree(elements []int) {
	for i := range len(elements) {
		t.Add(elements[i])
	}
}

// Inserta una lista de numeros dentro el arbol simultaneamente ocn varios hilos
func (t *tree) fillTreeGo(workers int, elements []int) {
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

/*
// Test sin hilos
func (t *tree) noGoTestSearch(searchArray []int) time.Duration {
	start := time.Now()

	for i := range searchArray {
		t.Search(searchArray[i])
	}

	return time.Since(start)
}

// Test con n hilos
func (t *tree) goTestSearch(workers int, searchArray []int) time.Duration {
	searchBlockLenght := len(searchArray) / workers
	var wg sync.WaitGroup

	start := time.Now()
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
	return time.Since(start)
}

// Ejecuta n test y calcula su promedio
func averageTime(runs int, threads int, test func() time.Duration) {
	var testTime time.Duration
	var totalTime time.Duration
	var minTime time.Duration
	var maxTime time.Duration

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
*/

func main() {
	t := new(tree)

	listNumbers := randomNumbersArray(5000000)
	searchArray := randomNumbersArray(500000)

	testTime := time.Now()
	t.fillTreeGo(4, listNumbers)
	fmt.Println(len(listNumbers), "elementos insertados en el arbol en", time.Since(testTime))
	fmt.Println()

	t.root = nil

	testTime2 := time.Now()
	t.fillTree(listNumbers)
	fmt.Println(len(listNumbers), "elementos insertados en el arbol en", time.Since(testTime2))

	fmt.Println("Arbol y array creados, empezando el test...")

	fmt.Println(len(searchArray))

	//averageTime(10, 0, func() time.Duration { return t.noGoTestSearch(searchArray) })
	//averageTime(10, 1, func() time.Duration { return t.goTestSearch(1, searchArray) })
	//averageTime(10, 4, func() time.Duration { return t.goTestSearch(4, searchArray) })
	//averageTime(10, 8, func() time.Duration { return t.goTestSearch(8, searchArray) })
	//averageTime(10, 16, func() time.Duration { return t.goTestSearch(16, searchArray) })

}

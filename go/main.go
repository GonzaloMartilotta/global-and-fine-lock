package main

import (
	"sync"
	"time"

	"bst/avgtime"
	"bst/fine"
	//"bst/global"
	//"bst/seq"
)

type BST interface {
	Add(value int)
	Search(value int) bool
	Delete(value int)
	Clear()
}

// Inserta una lista de numeros dentro el arbol
func fillTree(t BST, lenght int) {
	elements := avgtime.RandomNumbersArray(lenght)
	for i := range len(elements) {
		t.Add(elements[i])
	}
}

// Inserta una lista de numeros dentro el arbol simultaneamente ocn varios hilos
func fillTreeGo(t BST, workers int, lenght int) {
	elements := avgtime.RandomNumbersArray(lenght)
	blocks := len(elements) / workers
	var wg sync.WaitGroup
	for i := range workers {
		wg.Add(1)
		indexStart := i * blocks
		indexEnd := indexStart + blocks
		if i == workers-1 {
			indexEnd = len(elements)
		}
		block := elements[indexStart:indexEnd]

		go func(block []int) {
			defer wg.Done()
			for j := range block {
				t.Add(block[j])
			}
		}(block)
	}
	wg.Wait()
}

// Busca numeros sin gorutines

func searchNum(t BST, searchArray []int) {
	for i := range searchArray {
		t.Search(searchArray[i])
	}
}

// Lo mismo pero con gorutines
// PRE: El arbol tiene que tener elementos
func searchNumGo(t BST, workers int, searchArray []int) {
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

// Test de tiempo en add
// PRE: si se usa seq, NO poner mas de 1 worker
func addTest(t BST, lenght int, workers int) {
	avgtime.AverageTime(workers, func() time.Duration {
		t.Clear()
		start := time.Now()
		if workers == 1 {
			fillTree(t, lenght)
		} else {
			fillTreeGo(t, workers, lenght)
		}
		return time.Since(start)
	})
}

func main() {
	arbol := new(fine.Tree)
	workers := 0
	lenght := 500000

	addTest(arbol, lenght, workers)
}

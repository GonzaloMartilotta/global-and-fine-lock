package main

import (
	"fmt"
	"sync"
	"time"

	"bst/avgtime"
	"bst/fine"
	"bst/global"
	"bst/seq"
)

type BST interface {
	Add(value int)
	Search(value int) bool
	Delete(value int)
	Clear()
}

// Inserta una lista de numeros dentro el arbol
func fillTree(t BST, elements []int) {
	for i := range len(elements) {
		t.Add(elements[i])
	}
}

// Inserta una lista de numeros dentro el arbol simultaneamente ocn varios hilos
func fillTreeGo(t BST, workers int, elements []int) {
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
	fmt.Println("----------Test de Insertar----------")
	avgtime.AverageTime(workers, func() time.Duration {
		t.Clear()
		elements := avgtime.RandomNumbersArray(lenght)
		start := time.Now()
		if workers == 1 {
			fillTree(t, elements)
		} else {
			fillTreeGo(t, workers, elements)
		}
		return time.Since(start)
	})
	fmt.Println("------------------------------------")
	fmt.Println()
}

// PRE: El arbol tiene que tener elementos
func searchTest(t BST, lenght int, workers int) {
	fmt.Println("----------Test de busqueda----------")
	avgtime.AverageTime(workers, func() time.Duration {
		searchNumbers := avgtime.RandomNumbersArray(lenght)
		start := time.Now()
		if workers == 1 {
			searchNum(t, searchNumbers)
		} else {
			searchNumGo(t, workers, searchNumbers)
		}
		return time.Since(start)
	})
	fmt.Println("------------------------------------")
	fmt.Println()
}

func main() {
	var option int
	var test int
	running := true
	var runningTwo bool

	seqTree := new(seq.Tree)
	globalTree := new(global.Tree)
	fineTree := new(fine.Tree)

	workers := 1
	treeLenght := 2500000
	searchLength := 500000

	for running {
		fmt.Println("----------Demo menu----------")
		fmt.Println("1) Arbol binario de busqueda")
		fmt.Println("2) Arbol con global locking")
		fmt.Println("3) Arbol con fine locking")
		fmt.Println("4) Salir")
		fmt.Println("------------------------------------")
		fmt.Scan(&option)

		if option == 4 {
			running = false
		} else {
			runningTwo = true
			for runningTwo {
				fmt.Println("------------------------------------")
				fmt.Println("1) Add Test")
				fmt.Println("2) Search Test")
				fmt.Println("3) Salir")
				fmt.Println("------------------------------------")
				fmt.Scan(&test)
				switch test {
				case 1:
					switch option {
					case 1:
						addTest(seqTree, treeLenght, workers)
					case 2:
						addTest(globalTree, treeLenght, workers)
					case 3:
						addTest(fineTree, treeLenght, workers)
					}
				case 2:
					switch option {
					case 1:
						searchTest(seqTree, searchLength, workers)
					case 2:
						searchTest(globalTree, searchLength, workers)
					case 3:
						searchTest(fineTree, searchLength, workers)
					}
				case 3:
					runningTwo = false
				}
			}
		}
	}
}

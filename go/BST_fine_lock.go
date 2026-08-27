// Arbol con fine lock, la idea es que se puedan ejecutar varias operaciones en el arbol al mismo timepo
// la idea es mejorar el tiempo y testear la diferencia entre la cantidad de hilos proporcionados

package main

import (
	"fmt"
	"sync"
)

type tree struct {
	mu   sync.Mutex
	root *node
}

type node struct {
	mu    sync.Mutex
	Value int
	Left  *node
	Right *node
}

func add(value int, root *node) {
	if value < root.Value {
		if root.Left == nil {
			new := new(node)
			new.Value = value
			root.Left = new
			root.mu.Unlock()
		} else {
			root.Left.mu.Lock()
			root.mu.Unlock()
			add(value, root.Left)
		}
	} else if value > root.Value {
		if root.Right == nil {
			new := new(node)
			new.Value = value
			root.Right = new
			root.mu.Unlock()
		} else {
			root.Right.mu.Lock()
			root.mu.Unlock()
			add(value, root.Right)
		}
	} else {
		root.mu.Unlock()
	}
}
func (t *tree) Add(value int) {
	t.mu.Lock()
	if t.root == nil {
		newNode := new(node)
		newNode.mu.Lock()
		newNode.Value = value
		t.root = newNode
		t.root.mu.Unlock()
		t.mu.Unlock()
		return
	}
	t.root.mu.Lock()
	t.mu.Unlock()

	add(value, t.root)
}

func search(value int)

func navegate(root *node) *node {
	if root == nil {
		return nil
	}

	var option int
	fmt.Println("1 = Left, 2 = Right")
	fmt.Println("Node:", root)
	fmt.Print("Option: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		return navegate(root.Left)
	case 2:
		return navegate(root.Right)
	}
	return root
}

func menuTesting(t *tree) {
	var option int
	var value int
	var stay bool = true

	for stay {
		fmt.Println("---Menu---")
		fmt.Println("1-Add")
		fmt.Println("2-Delete")
		fmt.Println("3-Search")
		fmt.Println("4-Navegate")
		fmt.Println("5-Quit")
		fmt.Print("Option: ")
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Scan(&value)
			t.Add(value)
			fmt.Println("Valor agregado")
		case 4:
			navegate(t.root)
		case 5:
			stay = false
		}
	}
}

// Este arbol binario de busqueda esta protegido de que dos hilos accedan al mismo espacio de memoria
// NO ES MAS EFICIENTE solo evita que se corrompa la memoria al mandar muchos goroutines

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
	Value int
	Left  *node
	Right *node
}

func add(value int, root *node) *node {
	if root == nil {
		new := new(node)
		new.Value = value
		return new
	}

	if value < root.Value {
		root.Left = add(value, root.Left)
	} else if value > root.Value {
		root.Right = add(value, root.Right)
	}
	return root
}

func (t *tree) Add(value int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = add(value, t.root)
}

func search(value int, root *node) *node {
	if root == nil {
		return nil
	}

	if value < root.Value {
		return search(value, root.Left)
	} else if value > root.Value {
		return search(value, root.Right)
	} else {
		return root
	}
}

func (t *tree) Search(value int) *node {
	t.mu.Lock()
	defer t.mu.Unlock()
	return search(value, t.root)
}

func delete(value int, root *node) *node {
	if root == nil {
		return nil
	}

	if value < root.Value {
		root.Left = delete(value, root.Left)
	} else if value > root.Value {
		root.Right = delete(value, root.Right)
	} else {
		current := root

		if current.Left == nil {
			return current.Right
		} else if current.Right == nil {
			return current.Left
		}

		min := findMin(root.Right)
		root.Value = min.Value
		root.Right = delete(min.Value, root.Right)
	}
	return root
}

func (t *tree) Delete(value int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = delete(value, t.root)
}

func findMin(root *node) *node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

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
		case 2:
			fmt.Scan(&value)
			t.Delete(value)
		case 3:
			fmt.Scan(&value)
			node := t.Search(value)
			if node != nil {
				fmt.Println("Node found:", node)
			} else {
				fmt.Println("Node not found")
			}
		case 4:
			navegate(t.root)
		case 5:
			stay = false
		}
	}
}

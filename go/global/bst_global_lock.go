// Este arbol binario de busqueda esta protegido de que dos hilos accedan al mismo espacio de memoria
// NO ES MAS EFICIENTE solo evita que se corrompa la memoria al mandar muchos goroutines

package global

import (
	"fmt"
	"sync"
)

type Tree struct {
	mu   sync.Mutex
	root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func add(value int, root *Node) *Node {
	if root == nil {
		new := new(Node)
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

func (t *Tree) Add(value int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = add(value, t.root)
}

func search(value int, root *Node) *Node {
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

func (t *Tree) Search(value int) *Node {
	t.mu.Lock()
	defer t.mu.Unlock()
	return search(value, t.root)
}

func delete(value int, root *Node) *Node {
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

func (t *Tree) Delete(value int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = delete(value, t.root)
}

func (t *Tree) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = nil
}

func findMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func navegate(root *Node) *Node {
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

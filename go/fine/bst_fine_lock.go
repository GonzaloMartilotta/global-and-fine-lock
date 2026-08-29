// Arbol con fine lock, la idea es que se puedan ejecutar varias operaciones en el arbol al mismo timepo
// la idea es mejorar el tiempo y testear la diferencia entre la cantidad de hilos proporcionados

package fine

import (
	"fmt"
	"sync"
)

type Tree struct {
	mu   sync.Mutex
	root *Node
}

type Node struct {
	mu    sync.Mutex
	Value int
	Left  *Node
	Right *Node
}

func add(value int, root *Node) {
	if value < root.Value {
		if root.Left == nil {
			new := new(Node)
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
			new := new(Node)
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
func (t *Tree) Add(value int) {
	t.mu.Lock()
	if t.root == nil {
		newNode := new(Node)
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

func search(value int, root *Node) bool {
	if root.Value > value {
		if root.Left != nil {
			root.Left.mu.Lock()
			root.mu.Unlock()
			return search(value, root.Left)
		}
		root.mu.Unlock()
		return false
	} else if root.Value < value {
		if root.Right != nil {
			root.Right.mu.Lock()
			root.mu.Unlock()
			return search(value, root.Right)
		}
		root.mu.Unlock()
		return false
	} else {
		root.mu.Unlock()
		return true
	}
}
func (t *Tree) Search(value int) bool {
	t.mu.Lock()
	if t.root == nil {
		t.mu.Unlock()
		return false
	}
	t.root.mu.Lock()
	t.mu.Unlock()
	return search(value, t.root)
}

func (t *Tree) Delete(value int) {
	fmt.Println("No implementado")
}

func (t *Tree) Clear() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.root = nil
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

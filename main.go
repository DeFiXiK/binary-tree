package main

import "fmt"

// Tree - структура узла дерева.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Create - инициализирует пустое дерево
func Create() *Tree {
	return nil
}

// Found - функция которая осуществляет поиск заданного значения
// если значение не найдено, возвращает 0.
func (t *Tree) Found(v int) int {
	if t == nil {
		return 0
	}
	if v != t.Value {
		if v < t.Value {
			return t.Left.Found(v)
		}
		return t.Right.Found(v)
	}
	return t.Value
}

// InsertValue - добавляет лист к дереву,
// либо создает дерево с заданным элементом в корне если оно пустое
func (t *Tree) InsertValue(v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = t.Left.InsertValue(v)
		return t
	}
	t.Right = t.Right.InsertValue(v)
	return t
}

// InOrdered - префиксный обход дерева:
// левое поддерево, корень, правое поддерево
func (t *Tree) InOrdered() {
	if t == nil {
		return
	}
	t.Left.InOrdered()
	fmt.Printf("%v ", t.Value)
	t.Right.InOrdered()
}

func main() {
	tree := &Tree{nil, 8, nil}
	tree.InsertValue(4)
	tree.InsertValue(2)
	tree.InsertValue(3)
	tree.InsertValue(10)
	tree.InsertValue(6)
	tree.InsertValue(7)
	tree.InOrdered()
}

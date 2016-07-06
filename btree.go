package btree

import "fmt"

// Tree - структура узла дерева.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Create - инициализирует пустое дерево.
func Create() *Tree {
	return nil
}

// Find - функция которая проверяет наличие заданного значения
// если значение найдено то возвращает true, иначе false.
func (t *Tree) Find(v int) bool {
	if t == nil {
		return false
	}
	if v != t.Value {
		if v < t.Value {
			return t.Left.Find(v)
		}
		return t.Right.Find(v)
	}
	return true
}

// InsertValue - добавляет лист к дереву,
// либо создает дерево с заданным элементом в корне если оно пустое.
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
// левое поддерево, корень, правое поддерево.
func (t *Tree) InOrdered() {
	if t == nil {
		return
	}
	t.Left.InOrdered()
	fmt.Printf("%v ", t.Value)
	t.Right.InOrdered()
}

func main() {
	v := []int{8, 4, 2, 3, 10, 6, 7}
	tr := Create()
	for i := range v {
		tr = tr.InsertValue(v[i])
	}
	tr.InOrdered()
}

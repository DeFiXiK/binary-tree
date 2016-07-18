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

// InOrdered - инфиксный обход:
// левое поддерево, корень, правое поддерево.
func (t *Tree) InOrdered() {
	if t == nil {
		return
	}
	t.Left.InOrdered()
	fmt.Printf("%v ", t.Value)
	t.Right.InOrdered()
}

// PreOrder - префиксный обход дерева:
// левое поддерево, корень, правое поддерево.
func (t *Tree) PreOrder() {
	if t == nil {
		return
	}
	fmt.Printf("%v ", t.Value)
	t.Left.PreOrder()
	t.Right.PreOrder()
}

// Min - находит минимальный элемент в дереве
func (t *Tree) Min() int {
	if t.Left == nil {
		return t.Value
	}
	return t.Left.Min()
}

// Max - находит максимальный элемент в дереве
func (t *Tree) Max() int {
	if t.Right == nil {
		return t.Value
	}
	return t.Right.Max()
}

// InsertValues - позволяет ввести массив числе в дерево
func (t *Tree) InsertValues(values []int) *Tree {
	for v := range values {
		t = t.InsertValue(values[v])
	}
	return t
}

// Sorted - возвращает значения в узлах дерева отсортированным массивом.
func (t *Tree) Sorted() (array []int) {
	if t == nil {
		return nil
	}
	array = append(array, t.Left.Sorted()...)
	array = append(array, t.Value)
	array = append(array, t.Right.Sorted()...)
	return
}

// Remove - удаляет значение из дерева.
func (t *Tree) Remove(v int) *Tree {
	if t == nil || !t.Find(v) {
		return t
	}

	if v < t.Value {
		t.Left = t.Left.Remove(v)
		return t
	} else if v > t.Value {
		t.Right = t.Right.Remove(v)
		return t
	}

	// Then v == t.Value
	if t.Left != nil && t.Right != nil {
		t.Value = t.Right.Min()
		t.Right = t.Right.Remove(t.Value)
	} else if t.Left != nil {
		t = t.Left
	} else {
		t = t.Right
	}
	return t
}

func main() {
	v := []int{3, 2, 5, 4}
	tr := Create()
	for i := range v {
		tr = tr.InsertValue(v[i])
	}
	tr.InOrdered()
	fmt.Println()
	tr.PreOrder()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%v %v\n\n", tr.Min(), tr.Max())
	// tr = tr.Remove(2)
	tr.InOrdered()
}

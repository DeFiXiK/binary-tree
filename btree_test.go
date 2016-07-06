package btree

import "testing"

func TestFound(t *testing.T) {
	v := []int{8, 4, 2, 3, 10, 6, 7}
	tr := &Tree{nil, v[0], nil}
	if !tr.Find(v[0]) {
		t.Error("Значение в корне не найдено")
	}
	tr.Left = &Tree{nil, v[1], nil}
	if !tr.Find(v[1]) {
		t.Error("Значение левого потомка не найдено")
	}
	tr.Right = &Tree{nil, v[4], nil}
	if !tr.Find(v[4]) {
		t.Error("Значение правого потомка не найдено")
	}
	tr = Create()
	for i := range v {
		tr = tr.InsertValue(v[i])
	}
	for i := range v {
		if !tr.Find(v[i]) {
			t.Errorf("Значение %v - не найдено", v[i])
		}
	}
}

func TestInsertValue(t *testing.T) {
	tree := Create()
	tree = tree.InsertValue(2)
	if tree.Value != 2 {
		t.Errorf("Значение в правом потомке не совпадает ожидалось: %v, получено: %v ", 2, tree.Value)
	}
	tree = tree.InsertValue(3)
	if tree.Right.Value != 3 {
		t.Errorf("Значение в правом потомке не совпадает ожидалось: %v, получено: %v ", 3, tree.Value)
	}
	tree = tree.InsertValue(1)
	if tree.Left.Value != 1 {
		t.Errorf("Значение в правом потомке не совпадает ожидалось: %v, получено: %v ", 1, tree.Value)
	}
}

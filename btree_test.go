package btree

import (
	"reflect"
	"testing"
)

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
			t.Errorf("Значение %v - не найдено\n", v[i])
		}
	}
}

func TestInsertValue(t *testing.T) {
	tree := Create()
	tree = tree.InsertValue(2)
	if tree.Value != 2 {
		t.Errorf("Значение в правом потомке не совпадает ожидалось: %v, получено: %v\n", 2, tree.Value)
	}
	tree = tree.InsertValue(3)
	if tree.Right.Value != 3 {
		t.Errorf("Значение в правом потомке не совпадает ожидалось: %v, получено: %v\n", 3, tree.Value)
	}
	tree = tree.InsertValue(1)
	if tree.Left.Value != 1 {
		t.Errorf("Значение в правом потомке не совпадает ожидалось: %v, получено: %v\n", 1, tree.Value)
	}
}

func TestMax(t *testing.T) {
	tree := Create()
	v := []int{8, 4, 2, 3, 10, 6, 7}
	tree = tree.InsertValues(v)
	if tree.Max() != 10 {
		t.Errorf("Максимальный элемент найден не верно ожидалось: %v, получено: %v\n", 10, tree.Max())
	}
}

func TestMin(t *testing.T) {
	tree := Create()
	v := []int{8, 4, 2, 3, 10, 6, 7}
	tree = tree.InsertValues(v)
	if tree.Min() != 2 {
		t.Errorf("Минимальный элемент найден не верно ожидалось: %v, получено: %v\n", 2, tree.Min())
	}
}

func TestSorted(t *testing.T) {
	cases := []struct {
		in  []int
		out []int
	}{
		{nil, nil},
		{
			[]int{8, 4, 2, 3, 10, 6, 7},
			[]int{2, 3, 4, 6, 7, 8, 10},
		},
	}

	for _, c := range cases {
		t.Logf("Кейс %v...", c.in)
		tree := Create()
		tree = tree.InsertValues(c.in)
		ret := tree.Sorted()
		if !reflect.DeepEqual(ret, c.out) {
			t.Errorf("\tПолучено %v", ret)
			t.Errorf("\tОжидалось %v", c.out)
		}
	}
}

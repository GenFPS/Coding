package main

import (
	"fmt"
)

/*
Задача 3 (бонусная).

Есть односвязный список символов. Вывести его элементы задом наперед за один проход (За O(n)).
*/

// Создадим структуру односвязного списка
type ListNode struct {
	val  string
	next *ListNode
}

// создадим интерфейс для структуры ListNode
type ListNodeFunc interface {
	Print()   // выводит односвязный список
	Reverse() // переворачивает заданный список
}

// Реализация метода Print() для структуры LinkNode
func (s *ListNode) Print() {
	cur := s
	splitter := ""
	for cur != nil {
		if cur != s {
			splitter = " -> "
		}
		fmt.Printf("%s%s", splitter, cur.val)
		cur = cur.next
	}
	fmt.Println()
}

func (s *ListNode) Reverse() *ListNode {

	// используем два указателя - на предыдущий и текущий символ
	cur := s
	var prev *ListNode

	for cur != nil {
		// сохраняем ссылку на следующий символ чтобы заменить его на предыдущий
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	symb1 := &ListNode{val: "a"}

	symb2 := &ListNode{val: "b"}
	symb1.next = symb2

	symb3 := &ListNode{val: "c"}
	symb2.next = symb3

	symb1.Print()           // a -> b -> c
	symb1.Reverse().Print() // c -> b -> a
}

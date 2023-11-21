package main

import ( //импортирование двух пакетов: errors для создания и возврата ошибок и fmt для вывода информации на консоль.
	"errors"
)

// SNode - это структура, представляющая элемент (узел) стека. Узел содержит два поля: next, которое указывает на следующий узел в стеке, и val, которое хранит целочисленное (строковые) значение элемента стека.
type SNode struct { //стековая структура
	next *SNode //указатель на следующий узел в стеке
	val  string //значения типа string , хранящиеся в элементе стека
}

// Стек
type Stack struct {
	head *SNode //head - одно поле, которое указывает на верхний элемент стека
}

// Push добавляет новый элемент в стек
func (s *Stack) SPush(value string) {
	newNode := &SNode{
		val:  value,
		next: s.head,
	}
	s.head = newNode //обновление вершины стека на новый узел
}

// Pop удаляет и возвращает верхний элемент стека
func (s *Stack) SPop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("Stack is empty") //вызов проверки на пустоту стека
	}
	value := s.head.val  //сохранение значение верхнего элемента стека
	s.head = s.head.next //обновление вершины стека
	return value, nil    //возврат значения
}

// IsEmpty возвращает true, если стек пуст, иначе false
func (s *Stack) IsEmpty() bool {
	return s.head == nil //если равны, то стек считается пустым => true
}

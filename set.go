package main



//Определение структуры Set
type Set struct {
	elements map[string]bool //Используем map для хранения элементов множества
}

//Операция для создания нового множества 
func NewSet() *Set {
	s := &Set{}
	s.elements = make(map[string]bool)
	return s
	}   

//Операция для добавления элемента множества 
func (s *Set) SetAdd(element string) {
	s.elements[element] = true
}

//Операция для удаления элемента из множества
func (s *Set) SetRemove(element string) {
	delete(s.elements, element)
}

//Операция для проверки наличия элемента в множестве
func (s *Set) SetContains(element string) bool {
	return s.elements[element]
}

//Операция для получения размера множества 
func (s *Set) SetSize() int {
	return len(s.elements)
}

//Операция для вывода элементов множества
// Операция для вывода элементов множества
func (s *Set) SetPrint() string {
    var result string
    for element := range s.elements {
        result +=element
    }
    return result
}












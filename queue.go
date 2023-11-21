package main

import( //импортирование двух пакетов: errors для создания и возврата ошибок и fmt для вывода информации на консоль.
	"errors" 
)
type QNode struct{ //стековая структура (представляет элемент/узел в очереди)
	next *QNode //указатель на след. узел в очереди
	val string //значения типа string , хранящая в элементе очереди
}
//Очередь 
type Queue struct{
	front *QNode //указатель на первый элемент в очереди (голову)
	rear *QNode //указатель на последний элемент в очереди (хвост)
}
//Enqueue добавляет новый элемент в конец очереди
func (q *Queue) Qpush(value string){
	newNode:=&QNode{
		val: value, 
		next: nil,
	}
	//проверка на пустоту очереди
	if q.rear==nil{
		q.front = newNode
		q.rear = newNode
	} else{
		q.rear.next = newNode
		q.rear = newNode
	}
	}
	//Dequeue удаляет и возвращает элемент из начала очереди
	func(q *Queue) Qpop() (string, error){
		if q.IsEmpty(){
			return "", errors.New("Queue is empty")
	}

	value:=q.front.val	//извлекаем значение из переднего элемента очереди
	q.front=q.front.next //обновление указателя front

		//Проверка на пустоту очереди после извлечения
	if q.front==nil{
		q.rear=nil
	}
	return value, nil
}

//IsEmpty возвращает true, если очередь пуста, иначе false
func (q *Queue) IsEmpty() bool{
	return q.front==nil
}


package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type DataBases struct {
	Stack *Stack
	Queue *Queue
	Set *Set
	HashMap *HashMap
}

//Функция для создания базы данных
func CreationDatabase() *DataBases {
	Set := NewSet() 
	return &DataBases{
		Stack: &Stack{},
		Queue: &Queue{},
		Set: Set,
		HashMap: &HashMap{},
	}
}

//Функция Connection обрабатывает связь с подключённым клиентом (пользователем) и его запросом(команды)
func Connection(conn net.Conn, dataBas *DataBases) {
	defer conn.Close()
	bufer := make([]byte, 512) //создание bufer для чтения запросов пользователя(клиента)

	for {
		x, err := conn.Read(bufer) //чтение отправленных клиентом данных
		if err != nil {
			fmt.Printf("Соединение закрыто \n")
			return
		}
		//разбиение строки на элементы
		//Разбиение (разделение по запятой)
		input := string(bufer[:x])
		fmt.Printf("Полученный ввод: %s\n", input)
		strokaVvoda := strings.Split(input, " ") //инциализация строки ввода
		var singleCommand string //длина команды, длинной в единицу
		if len(strokaVvoda) == 1 {
			singleCommand = strokaVvoda[0] //инциализация коман6д без передаваемых значений
		}
		singleCommand = strings.TrimSpace(singleCommand) //убираем пробел в конце
		//fmt.Println(len(strokaVvoda))
		methodChoice := strokaVvoda[0]
		switch methodChoice { //для команд длинной больше единицы
		case "SPUSH":
			value := strokaVvoda[1]
			dataBas.Stack.SPush(value)
		case "QPUSH":
			value := strokaVvoda[1]
			dataBas.Queue.Qpush(value)
		case "SETADD":
			value := strokaVvoda[1]
			dataBas.Set.SetAdd(value)
		case "SETREMOVE":
			value := strokaVvoda[1]
			dataBas.Set.SetRemove(value)
		case "SETCONTAINS":
			value := strokaVvoda[1]
			res := dataBas.Set.SetContains(value)
			conn.Write([]byte(fmt.Sprintf("%t", res)))
		case "HPUSH":
			//key := strokaVvoda[1]
			if len(strokaVvoda)==3{
								
			value := strokaVvoda[2]
			
			//dataBas.HashMap.Insert(key, value)
			key:=strings.TrimSpace(strokaVvoda[1])
			dataBas.HashMap.Insert(key,value)

			}else{
				conn.Write([]byte("Command len < 3 "))
			}
		case "HGET": 
			key:=strings.TrimSpace(strokaVvoda[1])
			res,err:= dataBas.HashMap.HGet(key)
			if err!=nil{
				conn.Write([]byte(err.Error()))
			}
			conn.Write([]byte(res))
		case "HDEL":
			key := strings.TrimSpace(strokaVvoda[1])
			dataBas.HashMap.HDel(key)
			//conn.Write([]byte(fmt.Sprintf("%t", deleted)))
		}


		switch singleCommand { //для команд длиной единица 
		case "SPOP":
			poped, err := dataBas.Stack.SPop()
			if err != nil {
				conn.Write([]byte("Stack is empty"))
			} else {
				conn.Write([]byte(poped))
			}
		case "QPOP":  
			poped, err := dataBas.Queue.Qpop()
			if err != nil {
				conn.Write([]byte("Queue is empty"))
				} else {
					conn.Write([]byte(poped))
					}
		case "SETPRINT":
					setRespone := dataBas.Set.SetPrint()
					conn.Write([]byte(setRespone))
				
		}
	}



}
func main() {
	//Создание подключения
	listen, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Error")
		return
	}

	defer listen.Close() //закрытие пользователя, когда он выйдет (чтоб не было нагрузки на сервер и оперативную память)
	dataBas := CreationDatabase() //инциализация базы данных
	for { //цикл для постоянной обработки всех подключений
		conn, err := listen.Accept()
		if err != nil {

			fmt.Println(err)
			return
		}

		var group sync.WaitGroup //обработка распределения пользователей (для обработки множества запроов) очередь запросов
		group.Add(1)
		go func() {
			defer group.Done()
			go Connection(conn, dataBas)
		}()
		group.Wait()
	}
}

package main

import (
	"fmt"

	list "github.com/ernanej/ed-i/lists/doubleLinkedList"
)

func main() {
	doubleLinkedList := list.DoubleLinkedList{}
	doubleLinkedList.Init()

	fmt.Println("-> Lista inicializada.")
	doubleLinkedList.ShowData()

	fmt.Println("\n-> Inserindo 15 elementos na lista.")
	for i := 0; i < 15; i++ {
		doubleLinkedList.Add(i)
		doubleLinkedList.ShowData()
	}

	fmt.Println("\n-> Adicionando o numero -5 no index 5 da lista.")
	doubleLinkedList.AddOnIndex(5, -5)
	doubleLinkedList.ShowData()

	fmt.Println("\n-> Removendo o ultimo elemento da lista.")
	doubleLinkedList.Remove()
	doubleLinkedList.ShowData()

	fmt.Println("\n-> Aterando o número -5 no index 5 da lista para 5.")
	doubleLinkedList.Set(5, 5)
	doubleLinkedList.ShowData()

	fmt.Println("\n-> Removendo o número da posição 10 da lista.")
	doubleLinkedList.RemoveOnIndex(10)
	doubleLinkedList.ShowData()

	value, _ := doubleLinkedList.Get(10)
	fmt.Println("\n-> O elemento na posição 10 é: ", value)

	fmt.Println("\n-> A quantidade de elementos presentes na lista é: ", doubleLinkedList.Size())
}

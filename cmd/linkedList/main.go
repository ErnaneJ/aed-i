package main

import (
	"fmt"

	list "github.com/ernanej/ed-i/lists/linkedList"
)

func main() {
	linkedList := list.LinkedList{}
	linkedList.Init()

	fmt.Println("-> Lista inicializada.")
	linkedList.ShowData()

	fmt.Println("\n-> Inserindo 15 elementos na lista.")
	for i := 0; i < 15; i++ {
		linkedList.Add(i)
		linkedList.ShowData()
	}

	fmt.Println("\n-> Adicionando o numero -5 no index 5 da lista.")
	linkedList.AddOnIndex(5, -5)
	linkedList.ShowData()

	fmt.Println("\n-> Removendo o ultimo elemento da lista.")
	linkedList.Remove()
	linkedList.ShowData()

	fmt.Println("\n-> Aterando o número -5 no index 5 da lista para 5.")
	linkedList.Set(5, 5)
	linkedList.ShowData()

	fmt.Println("\n-> Removendo o número da posição 10 da lista.")
	linkedList.RemoveOnIndex(10)
	linkedList.ShowData()

	value, _ := linkedList.Get(10)
	fmt.Println("\n-> O elemento na posição 10 é: ", value)

	fmt.Println("\n-> A quantidade de elementos presentes na lista é: ", linkedList.Size())
}

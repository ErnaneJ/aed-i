package main

import (
	"fmt"

	list "github.com/ernanej/ed-i/lists/arrayList"
)

func main() {
	arraylist := list.ArrayList{}
	arraylist.Init()

	fmt.Println("-> Lista inicializada com capacidade de 10 elementos.")
	arraylist.ShowData()

	fmt.Println("\n-> Inserindo 15 elementos na lista. Deverá duplicar sua capcidade quando necessário.")
	for i := 0; i < 15; i++ {
		arraylist.Add(i)
		arraylist.ShowData()
	}

	fmt.Println("\n-> Adicionando o numero -5 no index 5 da lista.")
	arraylist.AddOnIndex(5, -5)
	arraylist.ShowData()

	fmt.Println("\n-> Removendo o ultimo elemento da lista.")
	arraylist.Remove()
	arraylist.ShowData()

	fmt.Println("\n-> Aterando o número -5 no index 5 da lista para 5.")
	arraylist.Set(5, 5)
	arraylist.ShowData()

	fmt.Println("\n-> Removendo o número da posição 10 da lista.")
	arraylist.RemoveOnIndex(10)
	arraylist.ShowData()

	value, _ := arraylist.Get(10)
	fmt.Println("\n-> O elemento na posição 10 é: ", value)

	fmt.Println("\n-> A quantidade de elementos presentes na lista é: ", arraylist.Size())
}

# ArrayList

> ArrayLists são listas implementadas com arrays (por "debaixo do panos").

![ArrayList representacao](../../imgs/arraylist-apresentacao.png)

## Métodos implementados:

- `initialize`: Inicializa a lista de acordo com uma capacidade especificada
- `getElementInPosition`: Captuda o valor correspondente à posição especificada
- `doubleCapacity`: Nos métodos em que acabamos excedendo a capacidade atual da lista utilizamos esse método para duplicar essa capacidade e armazenar a informação desejada.
- `insertElementAtEnd`: Insere elemento desejado no final da lista.
- `insertElementInPosition`: Insere um elemento na lista em uma determinada posiçao especificada.
- `updateElementInposition`: Atualiza o valor em uma posição já ocupada na lista. 
- `removeElementAtEnd`: Remove o ultimo elemento da lista.
- `removeElementInposition`: Remove um valor presente na lista em uma posição informada.
- `printListData`: Imprime no console as informações presentes na lista; seus elementos, quantidade de elementos e capacidade atual de armazenamento.
- `freeArrayList`: Destroi a lista e todas as suas infomações liberando espaço alocado previamente.

## Observações

- **Vantagens**:
  - Possuimos uma estrutura simples de se trabalhar;
  - Inserir e remover elementos no final da lista é extremamente barato `O(1)`.
  - Atualizar valores, sabendo sua posição é extremamente barato.
  - Temos um menor consumo de memória em comparação com outras EDs.
- **Desvantagens**:
  - Inserir e remover elementos em qualquer posição que não seja o final da lista é extremamente caro `O(n)` (pior caso) pois é necessário realizar toda uma operação de movimentação dos valores para suas posições correspondentes após inserção.
  - Eventualmente teremos espaço ocioso na memória (espaço alocado mas não utilizado).


## O método DoubleCapacity

Uma seguna, e possível, implementação mais manual do método para duplicar a capacidade da lista seria o modelo abaixo. Nele não precisamos do método `realoc` como utilizado [aqui](https://github.com/ErnaneJ/ED-I/blob/f04c0703847442b11a6eebf528f9c7f163adef3d/lists/arrayList/lib.h#L30).

```c
void doubleCapacity(struct ArrayList* list) {
  int new_capacity = list -> capacity * 2;
  int new_list[new_capacity];

  for(int i = 0; i < list -> length; i++){
    new_list[i] = list -> vector[i];
  }

  // DÚVIDA: a instrução abaixo atualiza a capacidade da lista, 
  // no material isso não é feito. Pq?
  list -> capacity = new_capacity;

  free(list -> vector);
  list -> vector = new_list;
}
```

## Execução: main.c

- Para executar utilize:

```bash
gcc ./lists/arrayList/main.c -o ./lists/arrayList/main && ./lists/arrayList/main
```

- Saída:

```bash
-> Lista inicializada com capacidade de 10 elementos.
.:: capacity: 10
.:: Quantidade de elementos: 0
.:: Elementos: []

-> insertElementAtEnd: Inserindo 20 elementos no fim da lista. A mesma deverá duplicar sua caapcidade.
.:: capacity: 20
.:: Quantidade de elementos: 20
.:: Elementos: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]

-> insertElementInPosition: Inserindo o número 999 na posição 10 da lista. A mesma deverá aumentar a quantidade de elementos em 1 e ter a capacidade dobrada.
.:: capacity: 40
.:: Quantidade de elementos: 21
.:: Elementos: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 999, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]

-> removeElementInposition: Removendo o número 999 da posição 10 da lista. A mesma deverá reduzir a quantidade de elementos em 1 e manter a capacidade.
.:: capacity: 40
.:: Quantidade de elementos: 20
.:: Elementos: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]

-> updateElementInposition: Atualizando o número 10 para 999 da posição 10 da lista. A mesma deverá permanecer com a mesma quantidade de elementos.
.:: capacity: 40
.:: Quantidade de elementos: 20
.:: Elementos: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 999, 11, 12, 13, 14, 15, 16, 17, 18, 19]

-> updateElementInposition: Voltando com o número 10 na posição 10 da lista. A mesma deverá permanecer com a mesma quantidade de elementos.
.:: capacity: 40
.:: Quantidade de elementos: 20
.:: Elementos: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19]

-> getElementInPosition: O elemento na posição 15 é: 15 
```
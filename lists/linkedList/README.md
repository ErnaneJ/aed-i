# LinkedList

> LinkedLists (ou listas ligadas, em português) são listas implementadas com nós, e cada nó possui um espaço de memória para armazenar o elemento e outro espaço de memória para armazenar o ponteiro (endereço de memória) para o nó seguinte.

![LinkedList representacao](../../imgs/linkedlist-apresentacao.png)

## Observações

- **Vantagens**:
  - Lista ligadas nunca terão espaço ocioso, ao contrário das ArrayLists por exemplo.
  - Temos a possibilidade de utilização da memória mesmo quando fragmentada.
- **Desvantagens**:
  - Para cada elemento armazenado na lista ligada, os `nós`, teremos um espaço adicional (`4 bytes`) utilizado para armazenar o ponteiro do próximo `nó` além do valor normal do tipo de dado armazenado na posição de valor de cada `nó`.
  - O modo de acesso às posições dos elementos é realizado de forma mais complicada já que os nós são alocados de forma não-contígua o que atrapalha o desempenho para uma série de operações.


## Execução: main.c

- Para executar utilize:

```bash
gcc ./lists/linkedList/main.c -o ./lists/linkedList/main && ./lists/linkedList/main
```

- Saída:

```bash
-> Inserindo elementos no final, a ordem deve ser 10, 20, 30 e o size deve ser 3
.:: Size: 3
.:: Elementos: [10, 20, 30]

-> Inserindo elemento (50) no inicio a ordem deve ser 50, 10, 20, 30 e o size deve ser 4.
.:: Size: 4
.:: Elementos: [50, 10, 20, 30]

-> Inserindo elemento (100) na posição 2, a ordem deve ser 50, 10, 100, 20, 30 e o size deve ser 5.
.:: Size: 5
.:: Elementos: [50, 10, 100, 20, 30]

-> Capturando elemento (20) da posição 3: 20

-> Removendo elemento (100) da posição 2, a ordem deve ser 50, 10, 20, 30 e o size deve ser 4.
.:: Size: 4
.:: Elementos: [50, 10, 20, 30]
```
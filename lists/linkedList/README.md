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
# LinkedList

> LinkedLists (ou listas ligadas, em português) são listas implementadas com nós, e cada nó possui um espaço de memória para armazenar o elemento e outro espaço de memória para armazenar o ponteiro (endereço de memória) para o nó seguinte.

![LinkedList representacao](../../imgs/linkedlist-apresentacao.png)

## Estrutura

```go
type LinkedList struct {
	head   *Node
	length int
}

type Node struct {
	value int
	next  *Node
}
```
## Observações

- **Vantagens**:
  - Matriz Dinâmica.
  - Facilidade de Inserção/Exclusão.
  - A inserção no início é uma operação de tempo constante e leva tempo O(1), em comparação com arrays onde inserir um elemento no início leva tempo O(n), onde n é o número de elementos no array.
  - Lista ligadas nunca terão espaço ocioso, ao contrário das ArrayLists por exemplo.
  - Temos a possibilidade de utilização da memória mesmo quando fragmentada.
- **Desvantagens**:
  - O acesso aleatório não é permitido. Temos que acessar os elementos sequencialmente a partir do primeiro nó (nó principal).
  - Espaço de memória extra (`4 bytes`), se compararmos com as ArrayLists, para um ponteiro é necessário com cada elemento da lista. 
  - Leva muito tempo para percorrer, se necessário, e alterar os ponteiros.
  - O acesso direto a um elemento não é possível
  - A busca por um elemento é cara e requer complexidade de tempo O(n).
  - Anexar um elemento a uma lista encadeada é uma operação cara e leva tempo O(n), onde n é o número de elementos na lista encadeada, em comparação com arrays que levam tempo O(1).

## Aplicações

- Listas vinculadas podem ser usadas para implementar estruturas de dados úteis, como pilhas e filas. 
- Listas vinculadas podem ser usadas para implementar tabelas de hash, cada bloco da tabela de hash pode ser uma lista vinculada. 
- As listas encadeadas podem ser usadas para implementar gráficos (representação da lista de adjacência do gráfico). 
- As listas encadeadas podem ser usadas de maneira refinada na implementação de diferentes sistemas de arquivos de uma forma ou de outra.
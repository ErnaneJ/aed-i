# Estrutura de Dados: Listas

| Operação                                           | ArrayLists                                                                       | LinkedLists                                                    | DoubleLinkedList                                                 |
|----------------------------------------------------|----------------------------------------------------------------------------------|----------------------------------------------------------------|------------------------------------------------------------------|
| Inserir elemento no fim (ainda há espaço no array) | **O(1)** *_performance não tem relação com o tamanho da lista_                   | **O(n)** *_performance **tem relação** com o tamanho da lista_ | performance **não tem relação** com o tamanho da lista: **O(1)** |
| Inserir elemento no fim (não há espaço no array)   | **O(n)** *_quando a realocação não pode ser feita reaproveitando o mesmo espaço_ | **O(n)** *_performance **tem relação** com o tamanho da lista_ | performance **não tem relação** com o tamanho da lista: **O(1)** |
| Inserir elemento no início                         | **O(n)** *_quando se insere na primeira posição de uma lista não vazia_          | **O(1)**                                                       | **O(1)**                                                         |
| Inserir elemento em posição                        | **O(n)** *_inserção no início_                                                   | **O(1)** *_inserção no fim_                                    | **O(1)**                                                         |
| Obter elemento em posição                          | **O(1)** *_conseguimos calcular o endereço de memória_                           | **O(n)**                                                       | **O(n)**                                                         |
| Atualizar elemento em posição                      | **O(1)** *_conseguimos calcular o endereço de memória_                           | **O(n)**                                                       | **O(n)**                                                         |
| Obter tamanho da lista                             | **O(1)**                                                                         | **O(1)**                                                       | **O(1)**                                                         |
| Remover elemento no fim                            | **O(1)**                                                                         | **O(n)**                                                       | **O(1)**                                                         |
| Remover elemento no início                         | **O(n)**                                                                         | **O(1)**                                                       | **O(1)**                                                         |
| Remover elemento em posição                        | **O(n)**                                                                         | **O(n)**                                                       | **O(n)**                                                         |

## ArrayList

> ArrayLists são listas implementadas com arrays (por "debaixo do panos").

![ArrayList representacao](../imgs/lists/arraylist-apresentacao.png)

### Estrutura

```go
type ArrayList struct {
	values []int
	length int
}
```

### Observações

- **Vantagens**:
  - Tamanho dinâmico: ArrayList pode aumentar e diminuir de tamanho dinamicamente, facilitando a adição ou remoção de elementos conforme necessário.
  - Fácil de usar: ArrayList é simples de usar, tornando-o uma escolha popular para muitos desenvolvedores.
  - Acesso rápido: ArrayList fornece acesso rápido aos elementos, pois é implementado como um array por baixo dos panos.
  - Coleção ordenada: ArrayList preserva a ordem dos elementos, permitindo que você acesse os elementos na ordem em que foram adicionados.
  - Suporta valores nulos: ArrayList pode armazenar valores nulos, tornando-o útil nos casos em que a ausência de um valor precisa ser representada.
  - Inserir e remover elementos no final da lista é extremamente barato `O(1)`.
  - Atualizar valores, sabendo sua posição é extremamente barato.
  - Temos um menor consumo de memória em comparação com outras EDs.
- **Desvantagens**:
  - Mais lento que arrays: ArrayList é mais lento que arrays para certas operações, como inserir elementos no meio da lista.
  - Maior uso de memória: ArrayList requer mais memória do que arrays, pois precisa manter seu tamanho dinâmico e manipular o redimensionamento.
    - Eventualmente teremos espaço ocioso na memória (espaço alocado mas não utilizado).
  - Não thread-safe: ArrayList não é thread-safe, o que significa que vários threads podem acessar e modificar a lista simultaneamente, levando a possíveis condições de corrida e corrupção de dados.
  - Degradação de desempenho: o desempenho de ArrayList pode diminuir à medida que o número de elementos na lista aumenta, especialmente para operações como procurar elementos ou inserir elementos no meio da lista.
    - Inserir e remover elementos em qualquer posição que não seja o final da lista é extremamente caro `O(n)` (pior caso) pois é necessário realizar toda uma operação de movimentação dos valores para suas posições correspondentes após inserção.


## LinkedList

> LinkedLists (ou listas ligadas, em português) são listas implementadas com nós, e cada nó possui um espaço de memória para armazenar o elemento e outro espaço de memória para armazenar o ponteiro (endereço de memória) para o nó seguinte.

![LinkedList representacao](../imgs/lists/linkedlist-apresentacao.png)

### Estrutura

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
### Observações

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

### Aplicações

- Listas vinculadas podem ser usadas para implementar estruturas de dados úteis, como pilhas e filas. 
- Listas vinculadas podem ser usadas para implementar tabelas de hash, cada bloco da tabela de hash pode ser uma lista vinculada. 
- As listas encadeadas podem ser usadas para implementar gráficos (representação da lista de adjacência do gráfico). 
- As listas encadeadas podem ser usadas de maneira refinada na implementação de diferentes sistemas de arquivos de uma forma ou de outra.


## DoubleLinkedList

> Uma lista duplamente ligada (DLL) contém um ponteiro extra, geralmente chamado de ponteiro anterior, junto com o próximo ponteiro e os dados que estão na lista vinculada individualmente.

![DoubleLinkedList representacao](../imgs/lists/doubly-linked-list-node.png)

### Estrutura

```go
type DoubleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}
```

### Observações

- **Vantagens**:
  - Pode ser percorrida nas direções para frente e para trás. 
  - A operação de exclusão na DLL é mais eficiente se for fornecido um ponteiro para o nó a ser excluído.
  - Podemos inserir rapidamente um novo nó antes de um determinado nó.
  - Na DLL, podemos obter o nó anterior facilmente utilizando o ponteiro anterior. 
- **Desvantagens**:
  - Cada nó da DLL requer espaço extra para um ponteiro anterior. 
  - Todas as operações requerem um ajuste de ponteiro extra para serem mantidas.
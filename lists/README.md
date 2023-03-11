# Estrutura de Dados: Listas

- [ArrayList](./arrayList/)
- [LinkedList](./linkedList/)
- [DoubleLinkedList](./doubleLinkedList/)

## Noções de desempenho

|   | Operação                                           | ArrayLists                                                                     | LinkedLists                                                  |   |
|---|----------------------------------------------------|--------------------------------------------------------------------------------|--------------------------------------------------------------|---|
|   | Inserir elemento no fim (ainda há espaço no array) | performance **não tem relação** com o tamanho da lista: **O(1)**               | performance **tem relação** com o tamanho da lista: **O(n)** |   |
|   | Inserir elemento no fim (não há espaço no array)   | quando a realocação não pode ser feita reaproveitando o mesmo espaço: **O(n)** | performance **tem relação** com o tamanho da lista: **O(n)** |   |
|   | Inserir elemento no início                         | quando se insere na primeira posição de uma lista não vazia: **O(n)**          | **O(1)**                                                     |   |
|   | Inserir elemento em posição                        | inserção no início: **O(n)**                                                   | inserção no fim: **O(n)**                                    |   |
|   | Obter elemento em posição                          | conseguimos calcular o endereço de memória: **O(1)**                           | **O(n)**                                                     |   |
|   | Atualizar elemento em posição                      | **O(1)**                                                                       | **O(n)**                                                     |   |
|   | Obter tamanho da lista                             | **O(1)**                                                                       | **O(1)**                                                     |   |
|   | Remover elemento no fim                            | **O(1)**                                                                       | **O(n)**                                                     |   |
|   | Remover elemento no início                         | **O(n)**                                                                       | **O(1)**                                                     |   |
|   | Remover elemento em posição                        | **O(n)**                                                                       | **O(n)**                                                     |   |
# Estrutura de Dados: Listas

- [ArrayList](./arrayList/)
- [LinkedList](./linkedList/)
- [DoubleLinkedList](./doubleLinkedList/)

## Noções de desempenho

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
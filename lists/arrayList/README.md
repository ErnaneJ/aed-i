# ArrayList

> ArrayLists são listas implementadas com arrays (por "debaixo do panos").

![ArrayList representacao](../../imgs/arraylist-apresentacao.png)

## Observações

- **Vantagens**:
  - Possuimos uma estrutura simples de se trabalhar;
  - Inserir e remover elementos no final da lista é extremamente barato `O(1)`.
  - Atualizar valores, sabendo sua posição é extremamente barato.
  - Temos um menor consumo de memória em comparação com outras EDs.
- **Desvantagens**:
  - Inserir e remover elementos em qualquer posição que não seja o final da lista é extremamente caro `O(n)` (pior caso) pois é necessário realizar toda uma operação de movimentação dos valores para suas posições correspondentes após inserção.
  - Eventualmente teremos espaço ocioso na memória (espaço alocado mas não utilizado).
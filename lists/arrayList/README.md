# ArrayList

> ArrayLists são listas implementadas com arrays (por "debaixo do panos").

![ArrayList representacao](../../imgs/arraylist-apresentacao.png)

## Estrutura

```go
type ArrayList struct {
	values []int
	length int
}
```

## Observações

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
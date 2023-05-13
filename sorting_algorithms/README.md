# Exercícios sobre Algoritmos de Ordenação

<p align="center">
  <b>Universidade Federal do Rio Grande do Norte (UFRN)</b> <br/>
  <b>Centro de Tecnologia - Departamento de Engenharia de Computação e Automação (DCA)</b>
</p>

- **Disciplina:** Algoritmos e Estruturas de Dados I (EDI)
- **Professor:** Eduardo de Lucena Falcão

---

1. **Programe na linguagem `C/Go` cada um dos seguintes algoritmos de ordenação:**
    - **a)** [SelectionSort](./selection_sort.go) (_in-place*_);
    - **b)** [BubbleSort](./bubble_sort.go) (_in-place*_);
    - **c)** [InsertionSort](./insertion_sort.go) (_in-place*_);
    - **d)** [MergeSort](./merge_sort.go);
    - **e)** [QuickSort](./merge_sort.go) (_com randomização de pivô_);
    - **f)** [CountingSort](./counting_sort.go).

    <br>

    _* `in-place` refere-se a um algoritmo ou função que altera os dados de entrada diretamente, sem a necessidade de criar uma cópia dos dados em uma nova área de memória._

2. **Explique a complexidade de tempo de pior caso dos algoritmos acima utilizando o código construído. Por fim, preencha a seguinte tabela, como forma de resumo.**

    | **Algoritmo**     |  **Pior Caso**                                                | **Melhor Caso** |                                                              **Meta**                                                               |
    |:------------------|:--------------------------------------------------------------|:----------------|:------------------------------------------------------------------------------------------------------------------------------------|
    | `BubbleSort`      | O(n²)                                                         |      Ω(n)       | Mover o maior elemento para a direita                                                                                               |
    | `CountingSort`    | O(n+k), k >> n                                                |      Ω(n)       | Classificar matrizes inteiras por contagem                                                                                          |
    | `InsertionSort`   | O(n²)                                                         |      Ω(n)       | Inserir o i-ésimo elemento em sua posição correta à esquerda                                                                        |
    | `MergeSort`       | O(nlog(n))                                                    |   Ω(nlog(n))    | Divide a matriz em metades com chamadas recursivas. Em seguida, dadas duas metades classificadas, mescla em uma matriz classificada |
    | `QuickSort`       | O(n²) if pivot isn't randomly taken, otherwise, O(nlog(n))    |   Ω(nlog(n))    | Em cada iteração, escolhe um pivô e coloqua os elementos maiores que ele à sua direita e os menores à sua esquerda                  |
    | `SelectionSort`   | O(n²)                                                         |      Ω(n²)      | Move o menor elemento para a esquerda                                                                                               |

3. **Ilustre, em detalhes, o funcionamento dos seguintes algoritmos com os seguintes vetores.**

    - **aleatório** = `[3, 6, 2, 5, 4, 3, 7, 1, 10⁹]`
    - **decrescente** = `[7, 6, 5, 4, 3, 3, 2, 1]`
    - **crescente** = `[1, 2, 3, 3, 4, 5, 6, 7]`

    - **a.** `SelectionSort` (in-place)
      - **i.** vetor para ilustrar: **aleatório**

      > **R:** É realizado um laço para iterar em cada elemento do array, e em cada iteração é realizado um novo laço que percorre desde o próximo elemento (i+1) até o final do array. Em cada iteração é comparado o menor elemento com o elemento atual e se o elemento atual for menor, guardamos o seu index. Ao final do loop mais interno é realizada a troca desse menor elemento com o indice de iteraçao externa atual. Ordenando sempre da esquerda para a direita, menor valor para o menor valor, considerando sempre a seleção dos menores valores. Abaixo está uma representação de como é realizada essas operações para ordenar o array de modelo aleatório.

      ```go
      list := []int{3, 6, 2, 5, 4, 3, 7, 1, 10⁹}
	    SelectionSort(list)

      // .:: Iteração: 1
      // .:: Comparando valores de 1 até 9...
      // .:: O menor valor encontrado foi 1 na posição 7.
      // .:: Realizamos a troca do elemento na posição 0 com o da posição 7.
      // .:: Estado atual da lista: [1 6 2 5 4 3 7 3 10⁹]

      // .:: Iteração: 2
      // .:: Comparando valores de 2 até 9...
      // .:: O menor valor encontrado foi 2 na posição 2.
      // .:: Realizamos a troca do elemento na posição 1 com o da posição 2.
      // .:: Estado atual da lista: [1 2 6 5 4 3 7 3 10⁹]

      // .:: Iteração: 3
      // .:: Comparando valores de 3 até 9...
      // .:: O menor valor encontrado foi 3 na posição 5.
      // .:: Realizamos a troca do elemento na posição 2 com o da posição 5.
      // .:: Estado atual da lista: [1 2 3 5 4 6 7 3 10⁹]

      //...

      // .:: Iteração: 9
      // .:: Comparando valores de 9 até 9...
      // .:: O menor valor encontrado foi 10⁹ na posição 8.
      // .:: Realizamos a troca do elemento na posição 8 com o da posição 8.
      // .:: Estado atual da lista: [1 2 3 3 4 5 6 7 10⁹]
      ```

    - **b.** `BubbleSort`
      - **i.** vetor para ilustrar: **aleatório**

      > **R:** Nesse caso será realizada uma iteração aninhada de dois loops. O loop mais externos corresponde ao indice atual da ordenação e o loop mais interno serve de varredura para a operação de comparação. Esse algoritmo realiza uma comparação entre os elementos e faz uma troca sempre que identifica que o elemento atual é maior que o próximo elemento. Ou seja, os elementos maiores 'flutuam' para o final do array. Abaixo está uma representação de como é realizada essas operações para ordenar o array de modelo aleatório.

      ```go
      list := []int{3, 6, 2, 5, 4, 3, 7, 1, 10⁹}
	    BubbleSort(list)

      // .:: Iteração:  1
      //   .:: [3 2 6 5 4 3 7 1 10⁹] 2 <-> 6
      //   .:: [3 2 5 6 4 3 7 1 10⁹] 5 <-> 6
      //   .:: [3 2 5 4 6 3 7 1 10⁹] 4 <-> 6
      //   .:: [3 2 5 4 3 6 7 1 10⁹] 3 <-> 6
      //   .:: [3 2 5 4 3 6 1 7 10⁹] 1 <-> 7

      // .:: Iteração:  2
      //   .:: [2 3 5 4 3 6 1 7 10⁹] 2 <-> 3
      //   .:: [2 3 4 5 3 6 1 7 10⁹] 4 <-> 5
      //   .:: [2 3 4 3 5 6 1 7 10⁹] 3 <-> 5
      //   .:: [2 3 4 3 5 1 6 7 10⁹] 1 <-> 6

      // .:: Iteração:  3
      //   .:: [2 3 3 4 5 1 6 7 10⁹] 3 <-> 4
      //   .:: [2 3 3 4 1 5 6 7 10⁹] 1 <-> 5

      // .:: Iteração:  4
      //   .:: [2 3 3 1 4 5 6 7 10⁹] 1 <-> 4

      // .:: Iteração:  5
      //   .:: [2 3 1 3 4 5 6 7 10⁹] 1 <-> 3

      // .:: Iteração:  6
      //   .:: [2 1 3 3 4 5 6 7 10⁹] 1 <-> 3

      // .:: Iteração:  7
      //   .:: [1 2 3 3 4 5 6 7 10⁹] 1 <-> 2

      // .:: Iteração:  8
      // .:: Não houve troca na iretação  8. Supomos que o array já está ordenado e encerramos o algorítmo.
      ```

    - **c.** `InsertionSort` (in-place)
      - **i.** vetor para ilustrar: **aleatório**

      > **R:** Como nos algoritmos anteriores, teremos dois laços. Um para controlar a iteração atual varrendo o vetor da esquerda para a direita e um outro laço mais interno que varre o vetor de forma contraria: da direita para a esqueda. Mas, essa varredura interna irá desde o final do vetor até a iteração atual mais externa (i). Ou seja, para o vetor atual no caso de i ser igual a 2 teramos o vetor interno j varrendo o vetor de n-1 até i-1. Essa varredura interna verifica se o elemento atual é menor que o elemento anterior, se for, realiza a troca. Sendo assim, essa abordagem se assemelha com o buble sorte porém, no lugar de fluturamos os elementos maiores para o fim, nós inserimos os elementos menores no inicio do array. Ou, flutuamos os menores para o início. Abaixo está uma representação de como é realizada essas operações para ordenar o array de modelo aleatório.
      
      ```go
      list := []int{3, 6, 2, 5, 4, 3, 7, 1, 10⁹}
      InsertionSort(list)

      //.:: Iteração:  1
      //  .:: [3 6 2 5 4 3 1 7 10⁹] 1 <-> 7
      //  .:: [3 6 2 5 4 1 3 7 10⁹] 1 <-> 3
      //  .:: [3 6 2 5 1 4 3 7 10⁹] 1 <-> 4
      //  .:: [3 6 2 1 5 4 3 7 10⁹] 1 <-> 5
      //  .:: [3 6 1 2 5 4 3 7 10⁹] 1 <-> 2
      //  .:: [3 1 6 2 5 4 3 7 10⁹] 1 <-> 6
      //  .:: [1 3 6 2 5 4 3 7 10⁹] 1 <-> 3

      //.:: Iteração:  2
      //  .:: [1 3 6 2 5 3 4 7 10⁹] 3 <-> 4
      //  .:: [1 3 6 2 3 5 4 7 10⁹] 3 <-> 5
      //  .:: [1 3 2 6 3 5 4 7 10⁹] 2 <-> 6
      //  .:: [1 2 3 6 3 5 4 7 10⁹] 2 <-> 3

      // .:: Iteração:  3
      //   .:: [1 2 3 6 3 4 5 7 10⁹] 4 <-> 5
      //   .:: [1 2 3 3 6 4 5 7 10⁹] 3 <-> 6

      // .:: Iteração:  4
      //   .:: [1 2 3 3 4 6 5 7 10⁹] 4 <-> 6

      // .:: Iteração:  5
      //  .:: [1 2 3 3 4 5 6 7 10⁹] 5 <-> 6

      // .:: Iteração:  6
      // .:: Não houve troca na iretação  6  supomos que o array já está ordenado e encerramos o algorítmo.
      ```

    - **d.** `MergeSort`
      - **i.** vetor para ilustrar: **decrescente**

      > **R:** O algoritmo de ordenação MergeSorte segue a premissa de subdividir o array desejado à ordenação em sua menor parte possivel (n arrays unitários) e, somente assim, realizar a mescla deles. Pois, dessa forma, é possível considerar que um array unitário sempre estará ordenado. Sendo assim, para o exemplo, o que teremos é a subdivisão  do array de exemplo em partes iguais recursivamente. Por exemplo:

      ```yml
      início: [7, 6, 5, 4, 3, 3, 2, 1]
              [7, 6, 5, 4][3, 3, 2, 1]
              [7, 6][5, 4][3, 3][2, 1]
              [7][6][5][4][3][3][2][1]
      ```

      Porém, é necessário ressaltar que esse processo não acontece de forma direta como na exemplificação acima. Ele é totalmente executado pela esquerda. Como assim? Bom, abaixo está uma representação de como é realizada essas operações para ordenar o array com o MergeSort.

      ```go
      list := []int{7, 6, 5, 4, 3, 3, 2, 1}
      MergeSort(list)
      
      // Divisão 0  -> [7 6 5 4 3 3 2 1]
      // Divisão 1  -> [7 6 5 4]
      // Divisão 2  -> [7 6]
      // Divisão 3  -> [7]
      // Divisão 4  -> [6]

      // Merge 0    -> [6] + [7] => [6 7]

      // Divisão 5  -> [5 4]
      // Divisão 6  -> [5]
      // Divisão 7  -> [4]

      // Merge 1    -> [4] + [5] => [4 5]
      // Merge 2    -> [4 5] + [6 7] => [4 5 6 7]

      // Divisão 8  -> [3 3 2 1]
      // Divisão 9  -> [3 3]
      // Divisão 10 -> [3]
      // Divisão 11 -> [3]

      // Merge 3    -> [3] + [3] => [3 3]

      // Divisão 12 -> [2 1]
      // Divisão 13 -> [2]
      // Divisão 14 -> [1]
      
      // Merge 4    -> [1] + [2] => [1 2]
      // Merge 5    -> [1 2] + [3 3] => [1 2 3 3]
      // Merge 6    -> [1 2 3 3] + [4 5 6 7] => [1 2 3 3 4 5 6 7]
      ```

    - **e.** `QuickSort` (**sem** randomização de pivô)
      - **i.** vetor para ilustrar: **crescente**

      > **R:** ...

      ```go
      rand.Seed(time.Now().UnixNano())
      list := []int{1, 2, 3, 3, 4, 5, 6, 7}
      QuickSort(false, list) // false para pivo padrão (ultimo elemento da lista)

      // .:: Pivot: 7 - Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 0 - smallersIterator: 0
      //     .::Comparação 1 < 7
      //     .::Troca 1 com 1
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 1 - smallersIterator: 1
      //     .::Comparação 2 < 7
      //     .::Troca 2 com 2
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 2 - smallersIterator: 2
      //     .::Comparação 3 < 7
      //     .::Troca 3 com 3
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 3 - smallersIterator: 3
      //     .::Comparação 3 < 7
      //     .::Troca 3 com 3
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 4 - smallersIterator: 4
      //     .::Comparação 4 < 7
      //     .::Troca 4 com 4
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 5 - smallersIterator: 5
      //     .::Comparação 5 < 7
      //     .::Troca 5 com 5
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 6 - smallersIterator: 6
      //     .::Comparação 6 < 7
      //     .::Troca 6 com 6
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //   .::Troca 7 com 7 <- ajusta posição do pivot

      // .:: Pivot: 6 - Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 0 - smallersIterator: 0
      //     .::Comparação 1 < 6
      //     .::Troca 1 com 1
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 1 - smallersIterator: 1
      //     .::Comparação 2 < 6
      //     .::Troca 2 com 2
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 2 - smallersIterator: 2
      //     .::Comparação 3 < 6
      //     .::Troca 3 com 3
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 3 - smallersIterator: 3
      //     .::Comparação 3 < 6
      //     .::Troca 3 com 3
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 4 - smallersIterator: 4
      //     .::Comparação 4 < 6
      //     .::Troca 4 com 4
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 5 - smallersIterator: 5
      //     .::Comparação 5 < 6
      //     .::Troca 5 com 5
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //   .::Troca 6 com 6 <- ajusta posição do pivot

      // .....

      // .:: Pivot: 2 - Lista: [1 2 3 3 4 5 6 7]
      //     .::largestIterator 0 - smallersIterator: 0
      //     .::Comparação 1 < 2
      //     .::Troca 1 com 1
      //   .:: Lista: [1 2 3 3 4 5 6 7]
      //   .::Troca 2 com 2 <- ajusta posição do pivot
      ```

    - **f.** `QuickSort` (**com** randomização de pivô)
      - **i.** vetor para ilustrar: **crescente**

      > **R:** ...

      ```go
      rand.Seed(time.Now().UnixNano())
      
      list := []int{1, 2, 3, 3, 4, 5, 6, 7}
      QuickSort(false, list) // false para pivo aleatório
      // ...
      ```

    - **g.** `CountingSort`
      - **i.** vetor para ilustrar: **aleatório**

      > **R:** ...

      ```go
      // ...
      ```

4. **A seguir são apresentados 5 fatos sobre algoritmos de ordenação. Planeje, execute experimentos, e apresente resultados que evidenciem cada afirmação.**

    - **A.** Para vetores de tamanho pequeno, a performance da maioria dos algoritmos de ordenação não vai influenciar, independente da disposição dos elementos.
      - **a.** Sugestão: um algoritmo `O(n²)`, um algoritmo `O(nlogn)`, e como exceção um algoritmo `O(k+n)`

    - **B.** Vetor de tamanho grande, a performance do algoritmo influencia de forma significativa. Além disso, dependendo da disposição (e valores) dos elementos no vetor, podemos experimentar performances bem diferentes (melhor e pior caso).
      - **a.** Sugestão: um algoritmo `O(n²)`, um algoritmo `O(nlogn)`, um algoritmo `O(k+n)`

    - **C.** `MergeSort` tem sempre um desempenho muito bom, independente da disposição dos elementos no vetor.
    
    - **D.** O pior caso do `Quicksort` é com o vetor ordenado de forma crescente/decrescente. O `Quicksort` com randomização de pivô resolve esse mau desempenho.
    
    - **E.** Explique quando o `CountingSort` tem bom desempenho e quando tem mau desempenho mostrando os resultados através dos experimentos.

5. **Perguntas com respostas rápidas:**
    - **a)** Por que `SelectionSort` não consegue melhorar o desempenho para cenários nos quais o vetor já está ordenado?
        > **R:** O `SelectionSort` é um algoritmo de ordenação que percorre o vetor diversas vezes, buscando o menor elemento e o colocando na posição correta. Quando o vetor já está ordenado, o seu desempenho não melhora pois o algoritmo ainda precisará percorrer todo o vetor para verificar se o próximo elemento é menor do que o atual. Isso ocorre porque o ele não aproveita a informação de que o vetor já está ordenado para reduzir a quantidade de comparações.

    - **b)** Por que BubbleSort consegue melhorar o desempenho para cenários nos quais o vetor já está ordenado? 
        > **R:** O `BubbleSort` é um algoritmo de ordenação que percorre o vetor diversas vezes, comparando pares de elementos adjacentes e trocando suas posições caso estejam na ordem errada. Quando o vetor já está ordenado, o seu desempenho melhora pois o algoritmo detecta que o vetor já está ordenado durante a primeira iteração e, portanto, não continua percorrendo o vetor desnecessariamente. Isso ocorre porque o ele se utiliza de uma estratégia chamada de "*adiantamento de verificação*", que compara os pares de elementos adjacentes e os troca somente se estiverem fora de ordem. Como o vetor já está ordenado, para esse caso, não haverá nenhum par de elementos fora de ordem e, portanto, o algoritmo não fará nenhuma troca de posições.

    - **c)** Por que `InsertionSort` consegue melhorar o desempenho para cenários nos quais o vetor já está ordenado?
        > **R:** O `InsertionSort` é um algoritmo de ordenação que percorre o vetor e, em cada iteração, insere o elemento atual na posição correta em relação aos elementos anteriores já ordenados. Quando o vetor já está ordenado, o seu desempenho melhora pois o algoritmo pode detectar rapidamente que um elemento é maior do que o elemento anterior e, portanto, não precisará percorrer o restante do vetor para encontrar a posição correta do elemento. Como o vetor já está ordenado, o algoritmo encontrará rapidamente a posição correta para cada elemento e, portanto, o número de comparações e movimentações de elementos será reduzido. Sendo assim, esse algorítimo consegue melhorar o seu desempenho para cenários nos quais o vetor já está ordenado porque utiliza-se da informação de que o vetor já está ordenado para reduzir a quantidade de comparações e movimentações de elementos.

    - **d)** Por que o `MergeSort` sempre tem o mesmo desempenho para qualquer cenário (vetor organizado de diferentes formas)?
        > **R:** O `MergeSort` é um algoritmo de ordenação que utiliza a estratégia de 'dividir e conquistar' para ordenar um vetor de elementos. Ele divide o vetor em duas metades, ordena cada metade separadamente e depois combina as duas metades em um vetor ordenado final. Dessa forma, ele sempre possui o mesmo desempenho para qualquer cenário pois a sua complexidade dependerá apenas do tamanho do vetor e não da ordem dos elementos presentes nele.

    - **e)** Por que o pior caso do `QuickSort` é `O(n²)`?
        > **R:** O `QuickSort`, semelhante ao `MergeSort` comentado anteriormente, é um algoritmo de ordenação baseado na estratégia de 'dividir e conquistar'. Ele seleciona um elemento, chamado de *pivô*, do vetor a ser ordenado e rearranja os outros elementos em torno dele, de forma que os elementos menores do que o *pivô* fiquem à sua esquerda e os elementos maiores fiquem à sua direita. Em seguida, o algoritmo aplica, de forma recursiva, o mesmo processo nas sub-listas à esquerda e à direita do *pivô*. No entanto, o pior caso do `QuickSort` ocorre quando o pivô escolhido sempre é o **maior ou o menor** elemento do vetor em cada recursão. Nesse caso, o `QuickSort` terá um desempenho de `O(n²)`, onde `n` é o tamanho do vetor. Isso acontece pois, no pior caso, o `QuickSort` terá que realizar um número quadrático de comparações e movimentações de elementos para ordenar o vetor. 

    - **f)** Como mitigar a probabilidade do pior caso acontecer no `QuickSort`?
        > **R:** Para mitigar o pior caso do `QuickSort`, é comum realizar uma escolha aleatória do pivô, na qual o pivô é escolhido *randomicamente* a partir do vetor em cada recursão. Essa abordagem reduz a probabilidade de escolher um pivô que gere o pior caso, melhorando assim o desempenho geral do algoritmo. Além disso, existem outras variantes do `QuickSort`, como o [*Dual Pivot QuickSort*](https://www.geeksforgeeks.org/dual-pivot-quicksort/), que são projetados exatamente para reduzir a probabilidade de ocorrer o pior caso.

    - **g)** O `CountingSort` é melhor ou pior do que o `MergeSort`? E em relação ao `QuickSort`?
        > **R:** O `CountingSort` é um algoritmo de ordenação que tem um desempenho *assintótico* melhor do que o `MergeSort` e o `QuickSort` em alguns cenários específicos. No entanto, o `CountingSort` possui algumas limitações que o tornam impraticável para uso geral. Em relação ao `MergeSort`, o `CountingSort` pode ter um desempenho melhor em cenários em que o intervalo dos valores do vetor é pequeno em comparação com o tamanho do vetor. O desempenho do `CountingSort` é `O(n + k)`, onde `n` é o tamanho do vetor e `k` é o intervalo dos valores do vetor. Em relação ao `QuickSort`, o `CountingSort` tem um desempenho assintótico melhor em **todos os casos**. Porém é válido ressaltar que, o `CountingSort` só pode ser usado para ordenar vetores com valores **inteiros** e com um **intervalo limitado**. Além disso, ele tem uma limitação de espaço, já que é necessário criar um vetor de contagem com tamanho igual ao intervalo dos valores do vetor. Sendo assim, podemos dizer que o `CountingSort` é uma boa opção quando se tem um vetor de valores inteiros com um intervalo limitado. Se o intervalo dos valores for grande, o `CountingSort` pode consumir muita memória, tornando-o uma má escolha. Para uso geral, o `MergeSort` e o `QuickSort` são opções melhores, sendo o `QuickSort` geralmente preferido em termos de desempenho e eficiência de memória.

---

<p align="center">
  <b>Universidade Federal do Rio Grande do Norte (UFRN)</b> <br/>
  <b>Centro de Tecnologia - Departamento de Engenharia de Computação e Automação (DCA)</b>
</p>
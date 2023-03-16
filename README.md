# Estruturas de Dados e Algoritmos (EDI - DCA0208)

Exercícios e anotações referentes ao estudo das notas de aula apresentadas na disciplina de Algoritmos e Estruturas de Dados I, ofertada pelo Departamento de Engenharia de Computação e Automação (UFRN).

## 🐛 Testes unitários

- Para executar todos os testes no diretório atual e todos os seus subdiretórios:
  ```bash
  go test ./...

  # Para ignorar os arquivos .go que não são testes
  sudo chmod +c ./scripts/run.sh # => permissão de execução
  ./scripts/run.sh
  ```

- Para executar todos os testes em determinados diretórios:
  ```bash
  go test ./tests/... ./unit-tests/... ./my-packages/...
  ```

- Para executar todos os testes com o caminho de importação prefixado com foo/:

  ```bash
  go test foo/...
  ```

- Para executar todos os caminhos de importação de testes prefixados com foo:

  ```bash
  go test foo...
  ```
- Para executar todos os testes em seu $GOPATH:

  ```bash
  go test ...
  ```

## ⚙️ Exemplos de utilização

São basicamente `scripts` que demonstram a maneira que o pacote seria utilizado em casos reais, demonstrando o seu funcionamento.

- Para executa-los urilize:
  ```bash
  go run ./cmd/package-folder

  # exemplo
  # go run ./cmd/arrayList
  ```


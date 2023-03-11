#include "lib.h"

int main(){
  printf("-> Lista inicializada com capacidade de 10 elementos.\n");
  struct ArrayList* list = initialize(10);
  printListData(list);

  printf("-> insertElementAtEnd: Inserindo 20 elementos no fim da lista. A mesma deverá duplicar sua caapcidade.\n");
  for(int i = 0; i < 20; i++){
    insertElementAtEnd(list, i);
  }
  printListData(list);

  printf("-> insertElementInPosition: Inserindo o número 999 na posição 10 da lista. A mesma deverá aumentar a quantidade de elementos em 1 e ter a capacidade dobrada.\n");
  insertElementInPosition(list, 999, 10);
  printListData(list);

  printf("-> removeElementInposition: Removendo o número 999 da posição 10 da lista. A mesma deverá reduzir a quantidade de elementos em 1 e manter a capacidade.\n");
  removeElementInposition(list, 10);
  printListData(list);

  printf("-> updateElementInposition: Atualizando o número 10 para 999 da posição 10 da lista. A mesma deverá permanecer com a mesma quantidade de elementos.\n");
  updateElementInposition(list, 999, 10);
  printListData(list);

  printf("-> updateElementInposition: Voltando com o número 10 na posição 10 da lista. A mesma deverá permanecer com a mesma quantidade de elementos.\n");
  updateElementInposition(list, 10, 10);
  printListData(list);

  printf("-> getElementInPosition: O elemento na posição 15 é: %d \n", getElementInPosition(list, 15));
  
  freeArrayList(list); // Libera a memória alocada para a lista

  return 0;
}
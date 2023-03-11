#include "lib.h"

int main(){
  struct Linkedlist* list = initialize();

  printf("-> Inserindo elementos no final, a ordem deve ser 10, 20, 30 e o size deve ser 3\n");
  insertElementAtEnd(list, 10);
  insertElementAtEnd(list, 20);
  insertElementAtEnd(list, 30);
  printListData(list);

  printf("-> Inserindo elemento (50) no inicio a ordem deve ser 50, 10, 20, 30 e o size deve ser 4.\n");
  insertElementAtStart(list, 50);
  printListData(list);

  printf("-> Inserindo elemento (100) na posição 2, a ordem deve ser 50, 10, 100, 20, 30 e o size deve ser 5.\n");
  insertElementInPosition(list, 100, 2);
  printListData(list);

  printf("-> Capturando elemento (20) da posição 3: %d\n\n", getElementInPosition(list, 3));

  printf("-> Removendo elemento (100) da posição 2, a ordem deve ser 50, 10, 20, 30 e o size deve ser 4.\n");
  removeElementInPosition(list, 2);
  printListData(list);

  return 0;
}
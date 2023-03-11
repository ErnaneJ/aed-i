#ifndef ARRAYLIST_H
#define ARRAYLIST_H

#include <stdio.h>
#include <stdlib.h>

struct ArrayList {
  int* vector;
  int length;
  int capacity;
};

struct ArrayList* initialize(int capacity) {
  struct ArrayList* list = (struct ArrayList*)malloc(sizeof(struct ArrayList));
  list -> vector = (int*)calloc(capacity, sizeof(int));
  list -> length = 0;
  list -> capacity = capacity;
  return list;
}

int getElementInPosition(struct ArrayList* list, int position) {
  if(position >= 0 && position < list -> length){
    return list -> vector[position];
  }
}

// Utilizando o realloc
void doubleCapacity(struct ArrayList* list) {
  int new_capacity = list -> capacity * 2;
  list -> vector = (int*)realloc(list -> vector, 2 * list -> capacity * sizeof(int));
  list -> capacity = new_capacity;
}

void insertElementAtEnd(struct ArrayList* list, int value) {
  if(list -> length == list -> capacity){
    doubleCapacity(list);
  }

  list -> vector[list -> length] = value;
  list -> length++;
}

void insertElementInPosition(struct ArrayList* list, int value, int position) {
  if(position >= 0 && position < list -> length){ // DÚVIDA: no material está <=, pq?
    if(list -> length == list -> capacity){
      doubleCapacity(list);
    }

    for(int i = list -> length; i > position; i--){
      list -> vector[i] = list -> vector[i - 1];
    }

    list -> vector[position] = value;
    list -> length++;
  }
}

void updateElementInposition(struct ArrayList* list, int value, int position) {
  if(position >= 0 && position <= list -> length){
    list -> vector[position] = value;
  }
}

void removeElementAtEnd(struct ArrayList* list) {
  list -> length--;
}

void removeElementInposition(struct ArrayList* list, int position) { // DÚVIDA: Essa implementação está correta. Mas, é eficiente? Como a do material?
  if(position >= 0 && position < list -> length){
    for(int i = position; i < list -> length - 1; i++){
      list -> vector[i] = list -> vector[i + 1];
    }

    list -> length--;
  }
}

void printListData(struct ArrayList* list) {
  printf(".:: capacity: %d", list -> capacity);
  printf("\n.:: Quantidade de elementos: %d", list -> length);
  printf("\n.:: Elementos: [");
  for (int i = 0; i < list -> length; i++) {
    printf("%d", list -> vector[i]);
    if (i < list -> length - 1) {
      printf(", ");
    }
  }
  printf("]\n\n");
}

void freeArrayList(struct ArrayList* list){
  free(list -> vector);
  free(list);
}

#endif //ARRAYLIST_H
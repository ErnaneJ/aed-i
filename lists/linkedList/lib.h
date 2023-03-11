#ifndef LINKEDLIST_H
#define LINKEDLIST_H

#include <stdio.h>
#include <stdlib.h>

struct node{
  int value;
  struct node* next;
};

struct Linkedlist {
  struct node* head;
  int size;
};

struct Linkedlist* initialize(){
  struct Linkedlist* list = (struct Linkedlist*)malloc(sizeof(struct Linkedlist));
  list -> head = NULL;
  list -> size = 0;

  return list;
}

void insertElementAtEnd(struct Linkedlist* list, int value){
  if(list->head == NULL){
    list -> head = (struct node*)malloc(sizeof(struct node));
    list -> head -> value = value;
    list -> head -> next = NULL;
  }else{
    struct node* aux = list -> head;
    while(aux -> next != NULL){
      aux = aux -> next;
    }

    aux -> next = (struct node*)malloc(sizeof(struct node));
    aux -> next -> value = value;
  }

  list -> size++;
}

void insertElementAtStart(struct Linkedlist* list, int value){
  if(list->head == NULL){
    list -> head = (struct node*)malloc(sizeof(struct node));
    list -> head -> value = value;
    list -> head -> next = NULL;
  }else{
    struct node* newNode = (struct node*)malloc(sizeof(struct node));
    newNode -> value = value;
    newNode -> next = list -> head;

    list -> head = newNode;
  }

  list -> size++;
}

void insertElementInPosition(struct Linkedlist* list, int value, int position){
  if(position == 0){
    insertElementAtStart(list, value);

  // DUVIDA: faz sentido proteger assim? Se eu colocar na posicao 1000 e a lista só tiver 3 elementos, adiciono no 4 ou lanço exceção?
  }else if(position >= list -> size){ 
    insertElementAtEnd(list, value);
  }else{
    struct node* positionElement = list -> head;
    for(int i = 0; i < position - 1; i++){
      positionElement = positionElement -> next;
    }

    struct node* newNode = (struct node*)malloc(sizeof(struct node));
    newNode -> value = value;
    newNode -> next = positionElement -> next;
    positionElement -> next = newNode;

    list -> size++;
  }
}

int getElementInPosition(struct Linkedlist* list, int position){
  if(position < 0 || position >= list -> size){ // DUVIDA: faz sentido essa proteção para esse escopo?
    printf("GetElementInPosition Error: invalid position.\n");
    return -1;
  }

  struct node* aux = list -> head;
  for(int i = 0; i < position; i++){
    aux = aux -> next;
  }

  return aux -> value;
}

void removeElementInPosition(struct Linkedlist* list, int position){
  if(position < 0 || position >= list -> size){ // DUVIDA: faz sentido essa proteção para esse escopo?
    printf("RemoveElementInPosition Error: invalid position.\n");
  }

  if(position == 0){
    struct node* aux = list -> head;
    list -> head = list -> head -> next;
    free(aux);
  }else{
    struct node* aux = list -> head;
    for(int i = 0; i < position - 1; i++){
      aux = aux -> next;
    }

    struct node* aux2 = aux -> next;
    aux -> next = aux2 -> next;
    
    free(aux2);
    list -> size--;
  }
}

void printListData(struct Linkedlist* list) {
  if (list -> head != NULL) {
    struct node* aux = list -> head;
    printf(".:: Size: %d", list -> size);
    printf("\n.:: Elementos: [");
    for(int i = 0; i < list -> size; i++){
      printf("%d", aux -> value);
      aux = aux -> next;
      if (aux != NULL) {
        printf(", ");
      }
    }
    printf("]\n\n");
  }
  else {
    printf("A lista está vazia!\n");
  }
}

#endif // LINKEDLIST_H
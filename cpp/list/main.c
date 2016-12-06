#include <stdio.h>
#include <stdlib.h>

#include "list.h"

typedef struct test
{
	int i;
	char *str;	
}T;

void showT(T *t){
	printf("t->i = %d\nt->str = %s\n",t->i,(t->str));
}

int main(void) {
	/*
	int i, length, is_next, is_prev;
	
	LinkedList *list = linked_create();

	i = -1;

	linked_add_last(&list, &i);

	for (i = 0; i < 6; i++) {
		linked_add_index(&list, (t + i), 1);
	}
	is_prev = 666;
	linked_set(list, 0, &is_prev);

	for (i = 0; i < 6; i++) {
		printf("list[%d] = %d\n", i, *((int *) linked_get(list, i)));
	}
	linked_show(list);
	linked_delete_first(&list);

	length = linked_length(list);

	printf("length = %d\n", length);
	 linked_show(list);
	 */
	 /*
	LinkedList *list = linked_create();

	 T *t = malloc(sizeof(T));
	 t->i = 100;
	 t->str = "200";
	showT(t);
	linked_add_first(&list, t);
	linked_show(list);
	free(t);
	showT(linked_get(list, 0));

	linked_show(list);
	linked_delete_first(&list);
	*/
	int i;
	int t[6] = { 0, 1, 2, 3, 4, 5 };
	RETURN(1, 2, 3);

	LinkedList *ll = linked_create(sizeof(int*));
	int data = 100;
	linked_add_first(&ll,&data);
	for(i = 0; i < 6; i++){
		linked_add_last(&ll, t+i);
	}
	printf("ll show\n");
	linked_show(ll);

	LinkedList *ll2 = linked_clone(ll);
	
	linked_add_all_2void(&ll, linked_toArray(ll2), linked_size(ll2));
	printf("linked_add_all_linked show\n");
	linked_show(ll);

	linked_free(&ll);
	// CHECK(ll, 5);
	printf("-------------------------------------------------\n");
	ArrayList *al = linked_parse_array(ll2);
	linked_free(&ll2);
	array_show(al);
	printf("-------------------------------------------------\n");
	array_delete_last(&al);
	for(i = 0; i < 6; i++){
		array_add_last(&al, t+i);
	}
	
	for(i = 0; i < 4; i++){
		printf("%d\n", i);
		array_add_last(&al, &i);
	}

	array_show(al);

	array_free(&al);
	array_show(al);
	exit(0);
}

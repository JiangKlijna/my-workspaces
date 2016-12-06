#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "list.h"

#define NUM 32
#if 0
#define NUM_2 64
#define NUM_3 128
#define NUM_4 256
#define NUM_5 512
#define NUM_6 1024
#define array_get_length(l) ((l) < NUM ? NUM : ((l) < NUM_2 ? NUM_2 : ((l) < NUM_3 ? NUM_3 : ((l) < NUM_4 ? NUM_4 : ((l) < NUM_5 ? NUM_5 : NUM_6)))))
#endif

#define array_set_node(list, p) ((list)->values[(p)])
#define array_get_node(list, p) (*((list)->values + (p)))
#define array_node(list, p) array_set_node((list), (p)) = malloc((list)->item_size)
#define array_foreach(list, data, i) for(i = 0,data = array_get_node(list, i); i < (list)->length; data = array_get_node(list, ++i))
#define linked_foreach(list, node) for((node) = (list)->first; (node) != NULL; (node) = (node)->next)
#define void2_foreach(values, data, i, n) for(i = 0,data = *(values + i); i < n; data = *(values + ++i))
#define linked_free_node(node) if(node){free(node->value);free(node);}
#define list_size(list) (list == NULL ? LIST_NULL : list->length)


typedef struct array_list {
	void **values;
	int length,actual,item_size;
} ArrayList;

typedef struct linked_node {
	struct linked_node *prev;
	struct linked_node *next;
	void *value;
} Node;

typedef struct linked_list {
	Node *first;
	Node *last;
	int length,item_size;
} LinkedList;

static Node * linked_node(int item_size) {
	Node *node = malloc(sizeof(*node));
	CHECK(node, NULL);
	node->value = malloc(item_size);
	node->next = NULL;
	node->prev = NULL;
	return node;
}

static Node * linked_get_node(const LinkedList *list, const int index) {
	DETERMINE(index < 0 || index >= (list->length), NULL);//-3

	RETURN(index, 0, list->first);
	RETURN(index, list->length - 1, list->last);

	int i;
	Node *node = NULL;
	if (index < list->length - 1 - index) {//F - L
		node = list->first;
		for (i = 0; i < index; i++) {
			node = node->next;
		}
	} else {//L - F
		node = list->last;
		for (i = 0; i < (list->length - 1 - index); i++) {
			node = node->prev;
		}
	}
	return node;
}


static void array_check(ArrayList *list, int num){
	if(list->length + num > list->actual){

	}
}

static void array_add(ArrayList *list, const void *data){
	array_node(list, list->length);
	memcpy(array_get_node(list, list->length++), data, list->item_size);
}

//LinkedList
LinkedList * linked_create(const int item_size) {
	LinkedList *list = malloc(sizeof(*list));
	CHECK(list, NULL);
	list->length = 0;
	list->first = NULL;
	list->last = NULL;
	list->item_size = item_size;
	return list;
}

int linked_size(const LinkedList *list) {
	return list_size(list);
}

void * linked_get_index(const LinkedList *list, const int index) {
	CHECK(list, NULL);
	Node *node = linked_get_node(list, index);
	return node == NULL ? node : node->value;
}

void * linked_get_first(const LinkedList *list) {
	return list == NULL ? NULL : linked_get_index(list, 0);
}

void * linked_get_last(const LinkedList *list) {
	return list == NULL ? NULL : linked_get_index(list, list->length-1);
}

int linked_set_index(const LinkedList *list, const void *data, const int index) {
	CHECK(list, LIST_NULL);
	Node *node = linked_get_node(list, index);
	CHECK(node, MALLOC_FAIL);

	memcpy(node->value, data, list->item_size);
	return LIST_OK;
}

int linked_set_first(const LinkedList *list, const void *data) {
	return list == NULL ? LIST_NULL : linked_set_index(list, data, 0);
}

int linked_set_last(const LinkedList *list, const void *data) {
	return list == NULL ? LIST_NULL : linked_set_index(list, data, list->length-1);
}

int linked_add_last(LinkedList **list, const void *data) {
	CHECK(*list, LIST_NULL);
	Node *node = linked_node((*list)->item_size);
	CHECK(node, MALLOC_FAIL);
	memcpy(node->value, data, (*list)->item_size);
	//node->value = data;
	node->prev = (*list)->last;
	//node->next = NULL;
	(*list)->last = node;
	(*list)->length++;
	if (node->prev == NULL) {
		(*list)->first = node;
	} else {
		node->prev->next = node;
	}
	return LIST_OK;
}

int linked_add_first(LinkedList **list, const void *data) {
	CHECK(*list, LIST_NULL);
	Node *node = linked_node((*list)->item_size);
	CHECK(node, MALLOC_FAIL);
	memcpy(node->value, data, (*list)->item_size);
	//node->value = data;
	node->next = (*list)->first;
	//node->prev = NULL;
	(*list)->first = node;
	(*list)->length++;
	if (node->next == NULL) {
		(*list)->last = node;
	} else {
		node->next->prev = node;
	}
	return LIST_OK;
}

int linked_add_index(LinkedList **list, const void *data, const int index) {
	CHECK(*list, LIST_NULL);
	DETERMINE(index < 0 || index > ((*list)->length), INVALID_PAR);

	RETURN(index, 0, linked_add_first(list, data));
	RETURN(index, (*list)->length, linked_add_last(list, data));

	Node *new_node = linked_node((*list)->item_size);
	CHECK(new_node, -2);

	Node *old_node = linked_get_node(*list, index);

	memcpy(new_node->value, data, (*list)->item_size);
	//new_node->value = data;
	new_node->prev = old_node->prev;
	new_node->next = old_node;
	old_node->prev->next = new_node;
	old_node->prev = new_node;
	(*list)->length++;
	return LIST_OK;
}

int linked_delete_last(LinkedList **list) {
	CHECK(*list, LIST_NULL);
	DETERMINE((*list)->length == 0, INVALID_PAR);

	Node *old_node = (*list)->last;
	(*list)->last = (*list)->last->prev;
	if ((*list)->last == NULL) {
		(*list)->first = NULL;
	} else {
		(*list)->last->next = NULL;
	}
	(*list)->length--;
	linked_free_node(old_node);
	return LIST_OK;
}

int linked_delete_first(LinkedList **list) {
	CHECK(*list, LIST_NULL);
	DETERMINE((*list)->length == 0, INVALID_PAR);

	Node *old_node = (*list)->first;
	(*list)->first = (*list)->first->next;
	if ((*list)->first == NULL) {
		(*list)->last = NULL;
	} else {
		(*list)->first->prev = NULL;
	}
	(*list)->length--;
	linked_free_node(old_node);
	return LIST_OK;
}

int linked_delete_index(LinkedList **list, const int index) {
	CHECK(*list, LIST_NULL);
	DETERMINE(index < 0 || index >= ((*list)->length), INVALID_PAR);

	RETURN(index, 0, linked_delete_first(list));
	RETURN(index, (*list)->length - 1, linked_delete_last(list));

	Node *old_node = linked_get_node(*list, index);
	CHECK(old_node, MALLOC_FAIL);

	old_node->prev->next = old_node->next;
	old_node->next->prev = old_node->prev;
	(*list)->length--;
	linked_free_node(old_node);
	return LIST_OK;
}

int linked_add_all_linked(LinkedList **dest, const LinkedList *src){
	CHECK(*dest, LIST_NULL);
	CHECK(src, INVALID_PAR);
	Node *node = NULL;
	linked_foreach(src, node){
		linked_add_last(dest, node->value);
	}
	return LIST_OK;
}

int linked_add_all_array(LinkedList **dest, const ArrayList *src){
	CHECK(*dest, LIST_NULL);
	CHECK(src, INVALID_PAR);
	int i;
	void *data = NULL;
	array_foreach(src, data, i){
		linked_add_last(dest, data);
	}
	return LIST_OK;
}

int linked_add_all_2void(LinkedList **dest, void **src, const int src_length){
	CHECK(*dest, LIST_NULL);
	CHECK(src, INVALID_PAR);
	int i;
	void *data = NULL;
	void2_foreach(src, data, i, src_length){
		linked_add_last(dest, data);
	}
	return -1;
}

LinkedList * linked_clone(const LinkedList *list){
	CHECK(list, NULL);
	LinkedList *dest = linked_create(list->item_size);
	CHECK(dest, NULL);
	Node *node = NULL;
	linked_foreach(list, node){
		linked_add_last(&dest, node->value);
	}
	return dest;
}

LinkedList * linked_sublist(const LinkedList *list){
	//TODO
	return NULL;
}

void ** linked_toArray(const LinkedList *list){
	CHECK(list, NULL);

	void **data = malloc(sizeof(void *) * list->length);
	static int i = 0;
	Node *node = NULL;
	linked_foreach(list, node){
		data[i++] = node->value;
	}
	return data;
}

ArrayList * linked_parse_array(const LinkedList *list){
	CHECK(list, NULL);
	ArrayList *dest = array_create(list->item_size);
	CHECK(dest, NULL);

	array_check(dest, list->length);
	Node *node = NULL;
	linked_foreach(list, node){
		array_add(dest, node->value);
	}
	return dest;
}

int linked_clear(LinkedList *list){
	CHECK(list, LIST_NULL);
	Node *node = NULL;
	linked_foreach(list, node){
		linked_free_node(node->prev);
	}
	linked_free_node(list->last);
	list->length = 0;
	return LIST_OK;
}

void linked_free(LinkedList **list){
	linked_clear(*list);
	free(*list);
	*list = NULL;
}

void linked_show(const LinkedList *list) {
	if (list == NULL) {
		return;
	}
	Node *node = NULL;
	printf("LinkedList.length = %d\n", list->length);
	linked_foreach(list, node){
		printf("%d\n", *((int *)(node->value)));
	}
}

//ArrayList
ArrayList * array_create(const int item_size) {
	ArrayList * list = NULL;
	list = malloc(sizeof(*list));
	if (list) {
		static void *values[NUM] ;
		list->values = values;
		list->actual = NUM;
		list->length = 0;
		list->item_size = item_size;
	}
	return list;
}

int array_size(const ArrayList *list){
	return list_size(list);
}

void * array_get_index(const ArrayList *list, const int index) {
	CHECK(list, NULL);
	DETERMINE(index < 0 || index >= list->length, NULL);
	return array_get_node(list , index);
}

void * array_get_first(const ArrayList *list){
	CHECK(list, NULL);
	return array_get_index(list, 0);
}

void * array_get_last(const ArrayList *list) {
	CHECK(list, NULL);
	return array_get_index(list, list->length - 1);
}

int array_set_index(const ArrayList *list, const void *data, const int index) {
	CHECK(list, LIST_NULL);
	DETERMINE(index < 0 || index >= list->length, INVALID_PAR);
	memcpy(array_get_node(list, index), data, list->item_size);
	return LIST_OK;
}

int array_set_first(const ArrayList *list, const void *data){
	CHECK(list, LIST_NULL);
	return array_set_index(list, data, 0);
}

int array_set_last(const ArrayList *list, const void *data) {
	CHECK(list, LIST_NULL);
	return array_set_index(list, data, list->length - 1);
}

int array_add_last(ArrayList **list, const void *data){
	CHECK(*list, LIST_NULL);
	array_check(*list, 1);
	array_add(*list, data);
	return LIST_OK;
}

int array_add_first(ArrayList **list, const void *data){
	return array_add_index(list, data, 0);
}

int array_add_index(ArrayList **list, const void *data, const int index){
	CHECK(*list, LIST_NULL);
	//TODO
	return LIST_OK;
}

int array_delete_last(ArrayList **list){
	CHECK(*list, LIST_NULL);
	free(array_get_node(*list, --(*list)->length));
	return LIST_OK;
}

int array_delete_first(ArrayList **list){
	//TODO
	return -1;
}

int array_delete_index(ArrayList **list, const int index){
	//TODO
	return -1;
}

int array_add_all_linked(ArrayList **dest, const LinkedList *src){
	return -1;
}

int array_add_all_array(ArrayList **dest, const ArrayList *src){
	return -1;
}

int array_add_all_2void(ArrayList **dest, const void **src){
	return -1;
}

ArrayList * array_clone(const ArrayList *list){
	return NULL;
}

ArrayList * array_sublist(const ArrayList *list){
	return NULL;
}

void ** array_toArray(const ArrayList *list){
	CHECK(list, NULL);
	return list->values;
}

LinkedList * array_parse_linked(const ArrayList *list){
	CHECK(list, NULL);
	LinkedList *dest = linked_create(list->item_size);
	CHECK(dest, NULL);
	int i;
	void *data = NULL;
	array_foreach(list, data, i){
		linked_add_last(&dest, data);
	}
	return dest;
}

int array_clear(ArrayList *list){
	CHECK(list , LIST_NULL);
	int i;
	for(i = 0; i < list->length; i++){
		free(array_get_node(list, i));
	}
	list->length = 0;
	return LIST_OK;
}

void array_free(ArrayList **list){
	array_clear(*list);
	free(*list);
	*list = NULL;
}

void array_show(const ArrayList *list){
	if(list == NULL){
		return;
	}
	printf("ArrayList.length = %d\n", list->length);

	void *data = NULL;
	int i;
	array_foreach(list, data, i){
		printf("%d\n", *((int *)data));
	}
}

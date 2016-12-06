#ifndef JIANG_LIST_H__
#define JIANG_LIST_H__

#define MALLOC_FAIL -3
#define INVALID_PAR -2
#define LIST_NULL -1
#define LIST_OK 0

// != 0
#define DETERMINE(b, v) if(b){return (v);}

#define RETURN(a, b, v) DETERMINE((a) == (b), (v))

#define CHECK(a, v) RETURN((a), NULL, (v))

#define PL printf(">>>%d\n", __LINE__)

//Implementation of Java.util.LinkedList
typedef struct linked_list LinkedList;

//Implementation of Java.util.ArrayList
typedef struct array_list ArrayList;

//LinkedList
LinkedList * linked_create(const int item_size);

int linked_size(const LinkedList *);

void * linked_get_index(const LinkedList *, const int index) ;

void * linked_get_first(const LinkedList *);

void * linked_get_last(const LinkedList *) ;

int linked_set_index(const LinkedList *, const void *, const int index) ;

int linked_set_first(const LinkedList *, const void *);

int linked_set_last(const LinkedList *, const void *) ;

int linked_add_last(LinkedList **, const void *);

int linked_add_first(LinkedList **, const void *);

int linked_add_index(LinkedList **, const void *, const int index);

int linked_delete_last(LinkedList **);

int linked_delete_first(LinkedList **);

int linked_delete_index(LinkedList **, const int index);

int linked_add_all_linked(LinkedList **dest, const LinkedList *src);

int linked_add_all_array(LinkedList **dest, const ArrayList *src);

int linked_add_all_2void(LinkedList **dest, void **src, const int src_length);

LinkedList * linked_clone(const LinkedList *);

LinkedList * linked_sublist(const LinkedList *);

void ** linked_toArray(const LinkedList *);

ArrayList * linked_parse_array(const LinkedList *);

int linked_clear(LinkedList *);

void linked_free(LinkedList **);

void linked_show(const LinkedList *);

//ArrayList
ArrayList * array_create(const int);

int array_size(const ArrayList *);

void * array_get_index(const ArrayList *, const int index) ;

void * array_get_first(const ArrayList *);

void * array_get_last(const ArrayList *) ;

int array_set_index(const ArrayList *, const void *, const int index) ;

int array_set_first(const ArrayList *, const void *);

int array_set_last(const ArrayList *, const void *) ;

int array_add_last(ArrayList **, const void *);

int array_add_first(ArrayList **, const void *);

int array_add_index(ArrayList **, const void *, const int index);

int array_delete_last(ArrayList **);

int array_delete_first(ArrayList **);

int array_delete_index(ArrayList **, const int index);

int array_add_all_linked(ArrayList **dest, const LinkedList *src);

int array_add_all_array(ArrayList **dest, const ArrayList *src);

int array_add_all_2void(ArrayList **dest, const void **src);

ArrayList * array_clone(const ArrayList *);

ArrayList * array_sublist(const ArrayList *);

void ** array_toArray(const ArrayList *);

LinkedList * array_parse_linked(const ArrayList *list);

int array_clear(ArrayList *);

void array_free(ArrayList **);

void array_show(const ArrayList *);
#endif

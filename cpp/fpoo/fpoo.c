#include <stdio.h>
#include <stdlib.h>

#define c_new_obj(type, ...) ({type* __t = (type*) malloc(sizeof(type));*__t = (type){__VA_ARGS__};__t; })

#define c_new_arr(type, size, ...) ({type** __t = (type**) malloc(sizeof(type*)*size);__t; })

typedef struct {
    int id;
    const char *name;

    void (*say)();
} Student;

void say() {
    printf("Student {id, name}\n");
}

void show(Student *s) {
    printf("Student {id=%d, name=%s}\n", s->id, s->name);
}

Student *newStudent() {
    return c_new_obj(Student, 0, "", say);
}

void map(void **arr, int length, void (*func)(void *)) {
    int i = 0;
    while (i < length) func(arr[i++]);
}

void fill(void **arr, int length, void *(*func)()) {
    int i = 0;
    while (i < length) arr[i++] = func();
}

int main() {
    Student *s = newStudent();
    s->say();
    free(s);

    void **arr = (void **) c_new_arr(Student, 10);
    fill(arr, 10, (void *(*)()) newStudent);
    map(arr, 10, (void (*)(void *)) show);
    map(arr, 10, free);
    return 0;
}
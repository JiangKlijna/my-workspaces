#include "sort.h"

int main(void) {
	test(new BubbleSort());
	exit(EXIT_SUCCESS);
	return 0;
}

void test(Sort* s) {
	IntArray* arr = new IntArray(20, 1000);
	s->sort(*arr);
	printf("%s\t", typeid(*s).name()+2);
	arr->print();
	delete arr;
}

// implement IntArray
IntArray::IntArray(int size, int digit) : length(size) {
	arr = (int*) malloc(sizeof(int) * size);
	srand((unsigned)time(0));
	for(int i = 0; i < size; i++) {
		arr[i] = RANDOM(digit);
	}
}
IntArray::~IntArray() {
	free(arr);
}
int& IntArray::operator[](int i) {
    return arr[i];
}
int IntArray::size() {
	return length;
}
void IntArray::print() {
	for(int i = 0; i < length; i++) {
		printf((i == 0 ? "[" : ""));
		printf("%d", arr[i]);
		printf((i == length-1 ? "]\n" : ", "));
	}
}


// 冒泡排序
void BubbleSort::sort(IntArray arr) {
}

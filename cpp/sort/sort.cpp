#include "sort.h"

int main(void) {
	test(new BubbleSort());
	test(new SelectSort());
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
	srand(time(0)-RANDOM(digit));
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
	for (int i = arr.size() - 1; i >= 0; i--) {
		for (int j = 0; j < i; j++) {
			if (arr[j] > arr[j + 1]) {
				int temp = arr[j + 1];
				arr[j + 1] = arr[j];
				arr[j] = temp;
			}
		}
	}
}
// 选择排序
void SelectSort::sort(IntArray arr) {
	for (int i = 0, n = arr.size(); i < n; i++) {
		int min = i;
		for (int j = i + 1; j < n; j++) {
			if (arr[j] < arr[min])
				min = j;
		}
		if (min != i) {
			int temp = arr[i];
			arr[i] = arr[min];
			arr[min] = temp;
		}
	}
}
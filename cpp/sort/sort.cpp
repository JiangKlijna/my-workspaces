#include "sort.h"

int main(void) {
	test(new BubbleSort());
	test(new SelectSort());
	test(new InsertSort());
	test(new ShellSort());
	test(new MergeSort());
	exit(EXIT_SUCCESS);
	return 0;
}

void test(Sort* s) {
	IntArray* arr = new IntArray(20, 1000);
	s->sort(*arr);
	printf("%s\t", s->getClassName());
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

const char* Sort::getClassName(){
	const char* name = typeid(*this).name();
	while(1) {
		if((*name) > '9') return name;
		name++;
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
// 插入排序
void InsertSort::sort(IntArray arr) {
	for (int i = 1, n = arr.size(); i < n; i++) {
		int j = i - 1;
		int tmp = arr[i];
		while (j >= 0 && tmp < arr[j]) {
			arr[j + 1] = arr[j];
			j--;
		}
		arr[j + 1] = tmp;
	}
}
// 希尔排序
void ShellSort::sort(IntArray arr) {
	int n = arr.size();
	int gap = n / 2;
	while (1 <= gap) {
		for (int i = gap; i < n; i++) {
			int j = i - gap;
			int tmp = arr[i];
			while (j >= 0 && tmp < arr[j]) {
				arr[j + gap] = arr[j];
				j -= gap;
			}
			arr[j + gap] = tmp;
		}
		gap /= 2;
	}
}
// 归并排序
void MergeSort::sort(IntArray arr) {
	sort(arr, 0, arr.size() - 1);
}
void MergeSort::sort(IntArray arr, int l, int r) {
	if (l >= r) return;
	int m = (l + r) / 2;
	sort(arr, l, m);
	sort(arr, m + 1, r);
	merge(arr, l, m, r);
}
void MergeSort::merge(IntArray arr, int l, int m, int r) {
	IntArray tmp(r - l + 1);
	int i = l, j = m + 1, k = 0;
	while (i <= m && j <= r) tmp[k++] = (arr[i] < arr[j]) ? arr[i++] : arr[j++];
	while (i <= m) tmp[k++] = arr[i++];
	while (j <= r) tmp[k++] = arr[j++];
	for(int i = l; i <= r; i++) arr[i] = tmp[i-l];
}
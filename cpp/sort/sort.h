#ifndef JIANGKLIJNA_SORT_H__
#define JIANGKLIJNA_SORT_H__
#include <cstdio>
#include <cstdlib>
#include <ctime>
#include <typeinfo>

#define RANDOM(digit) (rand() % digit)

class IntArray {
public:
    IntArray(int size=10, int digit=10);
	~IntArray();
    int& operator[](int index);
	int size();
	void print();
private:
	int* arr;
	int length;
};

class Sort {
public:
	virtual void sort(IntArray arr) = 0;
	const char* getClassName();
};

void test(Sort* s);

class BubbleSort : public Sort {
public: void sort(IntArray arr);
};
class SelectSort : public Sort {
public: void sort(IntArray arr);
};
class InsertSort : public Sort {
public: void sort(IntArray arr);
};
class ShellSort : public Sort {
public: void sort(IntArray arr);
};
class MergeSort : public Sort {
public: void sort(IntArray arr);
private:
	void sort(IntArray arr, int l, int r);
	void merge(IntArray arr, int l, int m, int r);
};
class QuickSort : public Sort {
public: void sort(IntArray arr);
private:
	void sort(IntArray arr, int l, int r);
	int partition(IntArray arr, int l, int r);
};
class HeapSort : public Sort {
public: void sort(IntArray arr);
private: void sift(IntArray arr, int l, int r);
};
#endif
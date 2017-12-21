#ifndef JIANGKLIJNA_SORT_H__
#define JIANGKLIJNA_SORT_H__
#include <cstdio>
#include <cstdlib>
#include <ctime>
#include <vector>
#include <typeinfo>

#define RANDOM(digit) (rand() % digit)

std::vector<int>* array(int size=10, int digit=10);

class Sort {
public:
	virtual void sort(std::vector<int>* arr) = 0;
};

void test(Sort* s);

class BubbleSort : public Sort {
public:
	void sort(std::vector<int>* arr);
};

#endif
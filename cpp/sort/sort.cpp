#include "sort.h"

using namespace std;

vector<int>* array(int size, int digit) {
	//int* arr = (int*)malloc(sizeof(int)*size);
	vector<int>* arr = new vector<int>();
	srand((unsigned)time(0));
	for(int i = 0; i < size; i++) {
		arr->push_back(RANDOM(digit));
	}
	return arr;
}

void test(Sort* s) {
	vector<int>* arr = array(20, 1000);
	s->sort(arr);
	printf("%s\t", typeid(*s).name());
	for(int i = 0, n = arr->size(); i < n; i++) {
		printf((i == 0 ? "[" : ""));
		printf("%d", arr->at(i));
		printf((i == n-1 ? "]\n" : ", "));
	}
	delete arr;
}

int main(void) {
	test(new BubbleSort());
	exit(EXIT_SUCCESS);
}

void BubbleSort::sort(vector<int>* arr) {

}
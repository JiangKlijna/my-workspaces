
from random import random as rdm

random = lambda digit : int(rdm() * digit)

array = lambda size=10, digit=10 : [random(digit) for e in range(size)]

def test(Sort):
	arr = array(20, 1000)
	Sort().sort(arr)
	print(Sort.__name__, arr, sep='\t')

class Sorter(object):
	def sort(arr): pass

# 冒泡排序
class BubbleSort(Sorter):
	def sort(self, arr):
		for i in range(len(arr)-1, 0, -1):
			for j in range(0, i):
				if arr[j] > arr[j+1]:
					tmp = arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = tmp
# 选择排序
class SelectSort(Sorter):
	def sort(self, arr):
		count = len(arr)
		for i in range(0, count):
			min = i
			for j in range(i+1, count):
				if arr[j] < arr[min]:
					min = j
			if i == min: continue
			tmp = arr[i]
			arr[i] = arr[min]
			arr[min] = tmp
# 插入排序
class InsertSort(Sorter):
	def sort(self, arr):
		count = len(arr)
		for i in range(1, count):
			j = i - 1
			tmp = arr[i]
			while j >= 0 and arr[j] > tmp:
				arr[j+1] = arr[j]
				j -= 1
			arr[j+1] = tmp
# 希尔排序
class ShellSort(Sorter):
	def sort(self, arr):
		count = len(arr)
		gap = int(count/2)
		while 1 <= gap:
			for i in range(gap, count):
				j = i - gap
				tmp = arr[i]
				while j >= 0 and arr[j] > tmp:
					arr[j+gap] = arr[j]
					j -= gap
				arr[j+gap] = tmp
			gap = int(gap/2)
# 归并排序
class MergeSort(Sorter):
	def __init__(self):
		self.sort = lambda arr : MergeSort.sort(arr, 0, len(arr)-1)
	@staticmethod
	def sort(arr, l, r):
		if l >= r: return
		m = int((l + r) / 2)
		MergeSort.sort(arr, l, m)
		MergeSort.sort(arr, m+1, r)
		MergeSort.merge(arr, l, m, r)
	@staticmethod
	def merge(arr, l, m, r):
		tmp = []
		i, j = l, m+1
		while i <= m and j <= r:
			if arr[i] < arr[j]:
				tmp.append(arr[i])
				i += 1
			else:
				tmp.append(arr[j])
				j += 1
		while i <= m:
			tmp.append(arr[i])
			i += 1
		while j <= r:
			tmp.append(arr[j])
			j += 1
		for i in range(r-l+1):
			arr[i+l] = tmp[i]

# 快速排序
class QuickSort(Sorter):
	def __init__(self):
		self.sort = lambda arr : QuickSort.sort(arr, 0, len(arr)-1)
	@staticmethod
	def sort(arr, l, r):
		#return
		if l >= r: return
		m = QuickSort.partition(arr, l, r)
		QuickSort.sort(arr, l, m-1)
		QuickSort.sort(arr, m+1, r)
	@staticmethod
	def partition(arr, l, r):
		k = l
		for i in range(l, r):
			if arr[i] < arr[r]:
				tmp = arr[i]
				arr[i] = arr[k]
				arr[k] = tmp
				k += 1
		tmp = arr[r]
		arr[r] = arr[k]
		arr[k] = tmp
		return k
# 堆排序
class DeapSort(Sorter):
	def sort(self, arr):
		pass
	def sift(self, arr, l, r):
		pass
# 基数排序
class RadixSort(Sorter):
	def sort(self, arr):
		pass
	def maxbit(self, arr):
		pass
# main
if __name__ == '__main__':
	sorters = [BubbleSort, SelectSort, InsertSort, ShellSort, MergeSort, QuickSort, DeapSort, RadixSort]
	[test(Sort) for Sort in sorters]

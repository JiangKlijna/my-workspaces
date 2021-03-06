package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
	"math/rand"
	"strconv"
)

var rdm *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func array(size int, digit int) []int {
	arr := make([]int, size, size)
	for i := range arr {
		arr[i] = rdm.Intn(digit)
	}
	return arr
}
func getSimpleName(i interface{}) string {
	name := reflect.ValueOf(i).Type().String()
	a := strings.Split(name, ".")
	return a[len(a) - 1]
}
func getArrayString(arr []int) string {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for i := 0; i < len(arr) - 1; i++ {
		buf.WriteString(strconv.Itoa(arr[i]))
		buf.WriteString(", ")
	}
	buf.WriteString(strconv.Itoa(arr[len(arr) - 1]))
	buf.WriteString("]")
	return buf.String()
}
func test(sort Sorter) {
	arr := array(20, 1000)
	sort.sort(arr)
	fmt.Printf("%s\t%s\n", getSimpleName(sort), getArrayString(arr))
}

type Sorter interface {
	sort(arr []int)
}

type BubbleSort struct{}

func (self BubbleSort) sort(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j + 1] {
				tmp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = tmp
			}
		}
	}
}

type SelectSort struct{}

func (self SelectSort) sort(arr []int) {
	count := len(arr)
	for i := 0; i < count; i++ {
		min := i
		for j := i + 1; j < count; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			tmp := arr[i]
			arr[i] = arr[min]
			arr[min] = tmp
		}
	}
}

type InsertSort struct{}

func (self InsertSort) sort(arr []int) {
	count := len(arr)
	for i := 1; i < count; i++ {
		j := i - 1
		tmp := arr[i]
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = tmp
	}
}

type ShellSort struct{}

func (self ShellSort) sort(arr []int) {
	count := len(arr)
	gap := count / 2
	for 1 <= gap {
		for i := gap; i < count; i++ {
			j := i - gap
			tmp := arr[i]
			for j >= 0 && arr[j] > tmp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = tmp
		}
		gap /= 2
	}
}

type MergeSort struct{}

func (self MergeSort) sort(arr []int) {
	self.recursive(arr, 0, len(arr)-1)
}

func (self MergeSort) recursive(arr []int, l int, r int)  {
	if l >= r { return }
	m := (r + l) / 2
	self.recursive(arr, l, m)
	self.recursive(arr, m+1, r)
	self.merge(arr, l, m, r)
}

func (self MergeSort) merge(arr []int, l int, m int, r int)  {
	tmp := make([]int, 0, r - l + 1)
	i, j := l, m+1
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	for i <= m {
		tmp = append(tmp, arr[i])
		i++
	}
	for j <= r {
		tmp = append(tmp, arr[j])
		j++
	}
	for i, v := range tmp {
		arr[l+i] = v;
	}
}

type QuickSort struct{}

func (self QuickSort) sort(arr []int) {
	self.recursive(arr, 0, len(arr)-1)
}

func (self QuickSort) recursive(arr []int, l int, r int)  {
	if l >= r { return }
	m := self.partition(arr, l, r)
	self.recursive(arr, l, m-1)
	self.recursive(arr, m+1, r)
}

func (self QuickSort) partition(arr []int, l int, r int) int {
	m := l
	for i := l; i < r; i++ {
		if arr[i] < arr[r] {
			temp := arr[i]
			arr[i] = arr[m]
			arr[m] = temp
			m++
		}
	}
	temp := arr[r]
	arr[r] = arr[m]
	arr[m] = temp
	return m
}

type HeapSort struct{}

func (self HeapSort) sort(arr []int) {
	n := len(arr) - 1
	for i := (n / 2); i >= 0; i-- {
		self.sift(arr, i, n)
	}
	for i := n; i >= 1; i-- {
		temp := arr[i]
		arr[i] = arr[0]
		arr[0] = temp
		self.sift(arr, 0, i - 1)
	}
}

func (self HeapSort) sift(arr []int, l int, r int) {
	i, j, tmp := l, l * 2 + 1, arr[l]
	for j <= r {
		if j < r && arr[j] < arr[j + 1] {j++}
		if tmp < arr[j] {
			arr[i] = arr[j]
			i = j
			j = 2 * i + 1
		} else {break}
	}
	arr[i] = tmp
}

type RadixSort struct{}

func (self RadixSort) sort(arr []int) {
	n := len(arr)
	d, radix, tmp := self.maxbit(arr), 1, make([]int, n, n)
	for i := 1; i <= d; i++ {
		count := make([]int, 10, 10)
		for j := 0; j < n; j++ {
			count[(arr[j] / radix) % 10]++
		}
		for j := 1; j < 10; j++ {
			count[j] += count[j - 1]
		}
		for j := n-1; j >= 0; j-- {
			k := (arr[j] / radix) % 10
			count[k] -= 1
			tmp[count[k]] = arr[j]
		}
		for j := 0; j < n; j++ {
			arr[j] = tmp[j]
		}
		radix *= 10
	}
}

func (self RadixSort) maxbit(arr []int) int {
	d, p := 1, 10
	for _, i := range arr {
		for i >= p {
			p *= 10
			d++
		}
	}
	return d;
}

func main() {
	test(BubbleSort{})
	test(SelectSort{})
	test(InsertSort{})
	test(ShellSort{})
	test(MergeSort{})
	test(QuickSort{})
	test(HeapSort{})
	test(RadixSort{})
}

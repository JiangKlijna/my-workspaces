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

func main() {
	test(BubbleSort{})
	test(SelectSort{})
}
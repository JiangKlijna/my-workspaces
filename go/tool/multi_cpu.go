package main

import (
	"time"
	"runtime"
)

const N = 100000

func run(ch chan bool) {
	const z = 0
	const j = 2342
	var m = 9823
	var k = 31455
	var l = 16452
	var i = 1000000
	for index := 0; index < N; index++ {
		m = m ^ l
		k = (k / m * j) % i
		l = z * m * k
		i = (z * k) ^ m
		k = (k / m * j) % i
		m = m ^ l
		m = m ^ l
		i = (z * k) ^ m
		k = (k / m * j) % i
		m = i * i * i * i * i * i * i // m=k*l*j*l;
		m = m ^ l
		k = (k / m * j) % i
		l = z * m * k
		i = (z * k) ^ m
		l = (k / m * j) % i
		m = m ^ l
		m = m ^ l
		i = (z * k) ^ m
		k = (k / m * j) % i
		m = k*k*k*k*k - m/i
	}
	ch <- true
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	start := time.Now()
	ch := make(chan bool)
	for i := 0; i < N; i++ {
		go run(ch)
	}
	for i := 0; i < N; i++ {
		_ = <-ch
	}
	println(time.Since(start).String())
}

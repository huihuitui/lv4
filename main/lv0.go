package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

var x int64
var ch = make(chan int, 1)
var ch1 = make(chan int, 1)

func Add1() {
	time.Sleep(10 * time.Millisecond)
	for i := 0; i < 500000; i++ {
		<-ch
		x++
		ch1 <- 1
	}
	wg.Done()
}

func Add2() {
	for i := 0; i < 500000; i++ {
		<-ch1
		x++
		ch <- 1
	}
	wg.Done()
}

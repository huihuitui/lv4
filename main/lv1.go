package main

import (
	"fmt"
)

func odd() {
	for i := 1; i <= 100; i += 2 {
		<-ch
		fmt.Println(i)
		ch1 <- 1
	}
	wg.Done()
}
func even() {
	for i := 0; i <= 100; i += 2 {
		<-ch1
		fmt.Println(i)
		ch <- 1
	}
	wg.Done()
}

package main

import (
	"fmt"
	"sync"
)

var wgg sync.WaitGroup

func main() {
	wgg.Add(1)
	out := make(chan int, 1)
	out <- 2
	<-out
	go f1(out)
	wgg.Wait()
}

func f1(in chan int) {
	in <- 114514
	a := <-in
	fmt.Println(a)
	wgg.Done()
}

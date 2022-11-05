package main

import (
	"fmt"
	"sync"
)

// import (
//
//	"fmt"
//	"sync"
//
// )
//
// var x int64
// var wg sync.WaitGroup
//
//	func add1(temp chan int64) {
//		for i := 0; i < 50000; i++ {
//			temp <- 1 //将一个值传入通道。通道满了。上锁。
//			x++
//			<-temp //通道通了。解锁。
//		}
//		wg.Done()
//	}
//
// func add2(temp chan int64) { //同理
//
//		for i := 0; i < 50000; i++ {
//			temp <- 1
//			x++
//			<-temp
//		}
//		wg.Done()
//	}
//
//	func main() {
//		temp := make(chan int64, 1) //一个容量只有1的通道。
//		wg.Add(2)
//		go add1(temp)
//		go add2(temp)
//		wg.Wait()
//		fmt.Println(x)
//	}
var x int64
var wg sync.WaitGroup

func add(temp chan int64) { //没上面这么麻烦了。思路还是一样的。
	for i := 0; i < 50000; i++ {
		temp <- 1
		x = x + 1
		<-temp
	}
	wg.Done()
}
func main() {
	temp := make(chan int64, 1)
	wg.Add(2)
	go add(temp)
	go add(temp)
	wg.Wait()
	fmt.Println(x)
}

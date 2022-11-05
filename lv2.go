package main

import (
	"fmt"
	"sync"
)

// var sum = 0
// var wg1 sync.WaitGroup
// var temp chan int
//
//	func jishu() {
//		temp <- 1
//		sum++
//		fmt.Println(sum)
//		if sum == 99 {
//			wg1.Done()
//		}
//	}
//
//	func oushu() {
//		<-temp
//		sum++
//		fmt.Println(sum)
//		if sum == 100 {
//			wg1.Done()
//		}
//	}
//
//	func main() {
//		wg1.Add(2)
//		temp = make(chan int, 1)
//		go jishu()
//		go oushu()
//		wg1.Wait()
//	}
//
// 一个错误探索↑
var wg1 sync.WaitGroup
var temp1 = make(chan int, 1)
var temp2 = make(chan int, 1)

//两个管道交替使用，交替打印。

func main() {
	temp1 <- 1 //给temp1一个值，在第一次进入循环的时候被读掉。作用是使第一次协程运行时管道不为空，能被接收。
	wg1.Add(2)
	go func() {
		defer wg1.Done()
		for i := 1; i < 101; i += 2 {
			<-temp1
			fmt.Println(i)
			temp2 <- 1
		}
		<-temp1 //最后跳出循环的时候temp1还带着值，得把这个值导出去（确信
	}()
	go func() {
		defer wg1.Done()
		for i := 2; i < 101; i += 2 {
			<-temp2
			fmt.Println(i)
			temp1 <- 1
		}
	}()
	wg1.Wait()
}

package main

//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	var channel = make(chan int)
//	var send = 6666
//	// 关键字go后跟的是需要被并发执行的代码块，它由一个匿名函数代表
//	// 在这里，我们只要知道在花括号中的就是将要并发执行的代码就可以
//	go func() {
//		channel <- send
//		fmt.Println("数据已发送") // 这个语句在这个例子里能不能执行，为什么呢？
//	}()
//	// 这里让主线程休眠1秒钟
//	// 以使上面的goroutine有机会被执行
//	time.Sleep(1 * time.Second)
//	<-channel //另外一个思路
//	println("can can need")
//	//var receive int//小改动，能跑了
//	//receive = <-channel
//	//fmt.Println(receive)
//}

/*
func main() {
	var channel = make(chan int)
	var send = 6666
	var receive int
	go func() {
		// 向channel中传递sent的值
		channel <- send
		fmt.Println("数据已发送")
	}()
	time.Sleep(1 * time.Second)
	// 使用receive接受从channel中传来的值
	receive = <-channel
	fmt.Println(receive)
}
*/
/*
func main() {
	// 声明了一个带有3个缓存空间的channel
	var channel = make(chan int, 3)

	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		fmt.Println("我发送了3个数据")
		// 当发送第四个值的时候，goroutine阻塞
		var receive int
		receive = <-channel
		receive = <-channel
		fmt.Println(receive)
		channel <- 4
		fmt.Println("我发送了第4个数据")
	}()
	time.Sleep(time.Second)
}
*/
/*
func main() {
	st := time.Now()      // 返回当前的时间戳
	ch := make(chan bool) // 这里make了一个无缓存的channel
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true                          // 无缓冲，发送方阻塞直到接收方接收到数据。
	fmt.Println("cost", time.Since(st)) //计算从st声明到现在的时间
	time.Sleep(time.Second * 5)
}
*/
/*
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y //交换变量的便利写法
		if i == 5 {
			close(c) //关闭通道
		}
	}
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c) //cap(c)即为c的容量大小
	// 不断读取channel数据，直到channel被显式关闭
	for i := range c {
		fmt.Println(i)
	}
}
*/
//
//func main() {
//	for i := range random(100) {
//		fmt.Println(i)
//	}
//}
//
//func random(n int) <-chan int {
//	c := make(chan int)
//	go func() {
//		defer close(c)
//		for i := 0; i < n; i++ {
//			select {
//			case c <- 0:
//			case c <- 1:
//			}
//		}
//	}()
//	return c
//}

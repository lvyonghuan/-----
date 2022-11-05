package main

import (
	"fmt"
	"io"
	"os"
)

//	func main() {
//		os.Create("./plan.txt")
//		var str1 string
//		str1 = "I’m not afraid of difficulties and insist on learning programming"
//		//file := io.Writer()
//		_, err := io.WriteString(io.Writer([]byte(str1)), str1)
//		if err != nil {
//			return
//		}
//	}
//
// 失败作
func main() {
	var file *os.File
	file, _ = os.Create("H:/GoProjects/work3/plan.txt")
	var str string
	str = "I’m not afraid of difficulties and insist on learning programming"
	_, _ = io.WriteString(file, str)
	var ReAd []byte
	buf := make([]byte, 114514)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		ReAd = append(ReAd, buf[:n]...)
	}
	fmt.Println(string(ReAd))
	defer file.Close()
}

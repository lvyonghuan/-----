package main

//
//import (
//	"fmt"
//	"io/ioutil"
//	"os"
//)
//
//func main() {
//	fmt.Println("嗨客网(www.haicoder.net)")
//	os.Create("H:/GoProjects/work3/haicoder.txt")
//	var (
//		fileName = "H:/GoProjects/work3/haicoder.txt"
//		content  = "Hello HaiCoder"
//		file     *os.File
//		err      error
//	)
//	//使用追加模式打开文件
//	file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
//	if err != nil {
//		fmt.Println("Open file err =", err)
//		return
//	}
//	defer file.Close()
//	//写入文件
//	n, err := file.Write([]byte(content))
//	if err != nil {
//		fmt.Println("Write file err =", err)
//		return
//	}
//	fmt.Println("Write file success, n =", n)
//	//读取文件
//	fileContent, err := ioutil.ReadFile(fileName)
//	if err != nil {
//		fmt.Println("Read file err =", err)
//		return
//	}
//	fmt.Println("Read file success =", string(fileContent))
//}

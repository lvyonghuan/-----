package main

import (
	"bytes"
	"fmt"
	"os"
)

// import (
//
//	"fmt"
//	"io"
//	"os"
//	"strings"
//
// )
//
//	func main() {
//		reader := strings.NewReader("Clear is better than clever")
//		ReadStr(reader)
//	}
//
//	func ReadStr(reader io.Reader) {
//		p := make([]byte, 5)
//		for {
//			n, err := reader.Read(p)
//			if err != nil {
//				if err == io.EOF {
//					fmt.Println("EOF:", n)
//					break
//				}
//				fmt.Println(err)
//				os.Exit(1)
//			}
//			fmt.Println(n, string(p[:n]))
//		}
//	}
func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	var writer bytes.Buffer

	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}

	fmt.Println(writer.String())
}

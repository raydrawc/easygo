// 程序包名
package main

// import "fmt"
// import "time"

import (
	"fmt"
	"time"
)

func main() {
	// golang 中表达式尾加不加“；” 都可以，建议不加
	fmt.Println("hello world")

	time.Sleep(1 * time.Second)

	fmt.Println("hello world")
}

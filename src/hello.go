// 包头
package main

// 单条导入
// import "fmt"
// import "time"

import (
	"fmt"
	"time"
)

// 大括号必须跟函数名同一行
func main() {
	// 行尾结束符号可以不加";", 推荐不加
	fmt.Println("hello")

	time.Sleep(1 * time.Second)

	fmt.Println("hello")
}

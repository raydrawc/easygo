package main

import "fmt"

// const 来定义枚举类型
const (
	// 可以再const() 提娜佳一个关键字ioTa, 每行的 iota 的默认值是0
	// 按顺序上一个值来定义
	// 第一个值必须有定义 missing constant value
	BEIJING = iota
	SHANGHAI
	SHENZHEN
	FOSHAN
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	j, k                      // iota = 4, j = iota * 2, k = iota * 3, j = 8, h = 12
)

func main() {
	// 测试变量
	lenght := 10
	fmt.Println("lenght = ", lenght)

	const pi = 3.1415
	fmt.Println("pi = ", pi)

	// pi = 20 //常量不可修改
	fmt.Println("======================== ")

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
	fmt.Println("FOSHAN = ", FOSHAN)
	fmt.Println("======================== ")

	fmt.Println("======================== ")
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("j = ", j, "k = ", k)
	fmt.Println("======================== ")

	// iota 只能配合const()  一起使用， iota只有在const进行累加效果

	const p = iota
	const q = iota

	fmt.Println("p = ", p, "q = ", q)

	// undefine iota
	// var gg int = iota

}

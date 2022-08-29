package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// 多个返回值 匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 888
}

// 多个返回值 带形参
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// r1,r2 属于foo3的形参， 初始化默认值都是0
	// r1, r2 作用控件是 foo3 整个函数体的{} 空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000 // 可注释 返回默认值
	// 可当作赋值 r1, r2
	return 10, 20
}

// 多个返回值 带形参 (变量同类型可以放一起)
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("foo2", 999)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	ret1, ret2 = foo4("foo4", 333)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

}

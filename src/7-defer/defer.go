package main

import "fmt"

func deferFun() int {
	fmt.Println("deferFun::")
	return 0
}

func returnFun() int {
	fmt.Println("returnFun::")
	return 0
}

func returnDeferCheck() int {
	// 先执行return 在执行defer
	defer deferFun()
	return returnFun()
}

func main() {
	// 加上defer，则在函数生命周期结束执行(类似析构函数)，后进先出
	defer fmt.Println("defer End")
	fmt.Println("main::print1")
	fmt.Println("main::print2")
	returnDeferCheck()
}

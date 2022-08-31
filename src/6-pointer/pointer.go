package main

import (
	"fmt"
)

func swap(a, b int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
}

func swap1(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	var a int = 10
	var b int = 20

	swap(a, b)
	fmt.Println("a = ", a, " b = ", b)
	swap1(&a, &b)
	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	// 二级指针
	var pp **int
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}

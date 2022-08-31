package lib2

import "fmt"

// 当前lib包提供的API
func LibTest() {
	fmt.Println("lib2Test()...")
}

func init() {
	fmt.Println("lib2.init()....")
}

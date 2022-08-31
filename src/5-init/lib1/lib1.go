package lib1

import "fmt"

// 当前lib包提供的API
func LibTest() {
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1.init()....")
}

package main

import (
	// "easygo/src/6-init/lib1"
	// 在包前加入 _ 表示匿名导入 (不使用这个包不报错)
	// 给包起一个别名（匿名）无法使用当前包的方法，但是会执行当前包的init()方法
	_ "easygo/src/6-init/lib1"
	mylib2 "easygo/src/6-init/lib2"
	// [不建议使用]使用 “.” 表示把包全部方法导入到本包 方法可以直接使用
	// . "easygo/src/6-init/lib2"
)

func main() {
	// lib1.LibTest()
	// lib2.LibTest()
	mylib2.LibTest()
}

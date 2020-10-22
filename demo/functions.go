package main

import (
	"fmt"
	"golangStudy/stringutil"
)

func functions() {
	fmt.Println("***** 函数调用 *****")
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))
	fmt.Println()

	fmt.Println("***** 多值返回 *****")
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println()

	fmt.Println("***** 命名返回值 *****")
	fmt.Println(split(17))
	c, d := split(17)
	fmt.Println(d, c)
	fmt.Println()

	fmt.Println("***** 多值返回 + 命名返回值 + 类型转换 *****")
	fmt.Println(assemble("这是", "字符串", 42, 13))
	fmt.Println()

	fmt.Println("***** 编写库并引用 *****")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println()
}

// 函数可以没有参数或接受多个参数，注意类型在变量名之后
func add(x int, y int) int {
	return x + y
}

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
func sub(x, y int) int {
	return x - y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 没有参数的 return 语句返回已命名的返回值，也就是 直接 返回
}

// 多值返回 + 命名返回值 + 类型转换的组合
func assemble(x, y string, m, n int) (a string, b float64) {
	a = x + y
	b = float64(m) / float64(n) // 类型转换：表达式 T(v) 将值 v 转换为类型 T
	return
}

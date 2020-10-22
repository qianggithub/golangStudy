package main

import "fmt"

func loop() {
	fmt.Println("***** for 循环 *****")
	forTest1()
	forTest2()
	fmt.Println()

	fmt.Println("***** while *****")
	while()
	fmt.Println()
}

// for 循环
func forTest1() {
	// 基本的 for 循环由三部分组成，它们用分号隔开：
	// 初始化语句：在第一次迭代前执行
	// 条件表达式：在每次迭代前求值
	// 后置语句：在每次迭代的结尾执行

	// 初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见

	sum := 0
	for i := 0; i < 2; i++ {
		sum += i
	}
	fmt.Println(sum) // 返回 1
}

// for 循环 精简版
func forTest2() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum) // 返回 1024
}

// "while" 循环
func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

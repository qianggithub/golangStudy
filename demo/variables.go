package main

import "fmt"

// var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
// var 语句可以出现在包或函数级别
var aa, bb int

// 变量声明可以包含初始值，每个变量对应一个
// 如果初始化值已存在，则可以省略类型，变量会从初始值中获得类型
var xx, yy = true, "字符串"

func variables() {
	fmt.Println("***** 变量定义(var 关键字) & 变量初始化 *****")
	var cc float64 = 66.666666666123456
	fmt.Println(aa, bb, cc)
	fmt.Println(xx, yy)
	fmt.Println()

	fmt.Println("***** 变量类型 *****")
	fmt.Printf("Type: %T Value: %v\n", xx, xx)
	fmt.Printf("Type: %T Value: %v\n", yy, yy)
	fmt.Println()

	fmt.Println("***** 短变量声明 *****")
	shortVariables()
	fmt.Println()
}

// 短变量声明
func shortVariables() {
	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
	mm, nn := "双击", 666
	fmt.Println(nn, mm)
	fmt.Printf("Type: %T Value: %v\n", nn, nn)
	fmt.Printf("Type: %T Value: %v\n", mm, mm)
}

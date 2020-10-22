package main

import (
	"fmt"
	"math"
)

func if_() {
	fmt.Println("***** if 语句 *****")
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println()

	fmt.Println("***** if 简短语句 *****")
	fmt.Println(
		shortIf(3, 2, 10),
		shortIf(3, 3, 20),
	)
	fmt.Println()

	fmt.Println("***** if 和 else *****")
	fmt.Println(
		ifAndElse(3, 2, 10),
		ifAndElse(3, 3, 20),
	)
	fmt.Println()
}

// if 语句
func sqrt(x float64) string {
	// Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if 简短语句
func shortIf(x, n, lim float64) float64 {
	// 同 for 一样， if 语句可以在条件表达式前执行一个简单的语句
	// 该语句声明的变量作用域仅在 if 之内
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if 和 else
func ifAndElse(x, n, lim float64) float64 {
	// 在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

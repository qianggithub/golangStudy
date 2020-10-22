package main

import "fmt"

func array() {
	// 类型 [n]T 表示拥有 n 个 T 类型的值的数组
	// 数组的长度是其类型的一部分，因此数组不能改变大小
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	b := []string{"测", "试"}
	fmt.Println(b[1])
}

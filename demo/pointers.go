package main

import "fmt"

// 类型 *T 是指向 T 类型值的指针。指针保存了值的内存地址。其零值为 nil。
var p *int

func pointers() {
	fmt.Println(p)

	i, j := 42, 2701

	// & 操作符会生成一个指向其操作数的指针
	// * 操作符表示指针指向的底层值
	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j // 指向 j
	fmt.Println(p)
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

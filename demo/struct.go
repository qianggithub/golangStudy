package main

import "fmt"

// 一个结构体（struct）就是一组字段（field）
type Vertex struct {
	X int
	Y int
}

type Student struct {
	name, address string
	age, phone    int
}

// 结构体文法通过直接列出字段的值来新分配一个结构体
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	v4 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func struct_() {
	fmt.Println(Student{})
	fmt.Println(v1, v2, v3, v4)

	// 结构体字段使用点号来访问
	v := Vertex{6, 8}
	fmt.Println(v.Y)
	v.X = 10 // 设置结构体 X 的值
	fmt.Println(v.X)

	// 有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X
	// 也可以使用隐式间接引用，直接写 p.X 就可以
	v = Vertex{66, 88}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

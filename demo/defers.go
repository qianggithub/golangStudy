package main

func defers() {
	// defer 语句会将函数推迟到外层函数返回之后执行
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	// 如果有多个defer表达式，调用顺序类似于栈，越后面的defer表达式越先被调用

	// 参考：https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
	// 函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中

	// 使用defer时，用一个简单的转换规则改写一下，改写规则如下
	// 1.返回值 = xxx
	// 2.调用defer函数
	// 3.空的return

	println(deferTest1()) // 返回 1
	println(deferTest2()) // 返回 5
	println(deferTest3()) // 返回 1
	println(deferTest4()) // 返回 6
}

func deferTest1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func deferTest2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func deferTest3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func deferTest4() (r int) {
	defer func() {
		r = r + 5
	}()
	return 1
}

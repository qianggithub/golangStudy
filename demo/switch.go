package main

import (
	"fmt"
	"runtime"
	"time"
)

func switch_() {
	fmt.Println("***** switch 语句 *****")
	switchTest1()
	fmt.Println()

	fmt.Println("***** switch 的求值顺序 *****")
	switchTest2()
	fmt.Println()

	fmt.Println("***** 没有条件的 switch *****")
	switchTest3()
	fmt.Println()
}

// switch 语句
func switchTest1() {
	// Go 自动提供了在这些语言中每个 case 后面所需的 break 语句
	// 除非以 fallthrough 语句结束，否则分支会自动终止
	// switch 的 case 无需为常量，且取值不必为整数
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

// switch 的求值顺序
func switchTest2() {
	// switch 的 case 语句从上到下顺次执行，直到匹配成功时停止
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// 没有条件的 switch
func switchTest3() {
	// 没有条件的 switch 同 switch true 一样
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

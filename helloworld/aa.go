package main

import (
	"fmt"
	"math"
)

func main() {
	//cc()
	//变量()
	//常量()

}

func for循环() {

}

const 常量s string = "constant"

func 常量() {
	fmt.Println(常量s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}

func cc() {
	//字符串
	fmt.Println("go" + "java")
	//	整型,拼接字符串的时候需要使用，不同类型之间的拼接
	fmt.Println("1+1=", 1+1, "字符串拼接")
	//	浮点型
	fmt.Println("7.0/3.0=", 7.0/3.0)
	//	布尔运算
	fmt.Println(!true)
}

func 变量() {
	var str = "abc"
	fmt.Println(str)
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)
	//看一下e的初始值 0
	var e int
	fmt.Println(e) //0
	f := "short"   //已经确定是字符串的类型可以不使用var
	fmt.Println(f)
	g := 2 //同理，已经确定的类型可以不使用
	fmt.Println(g)
}

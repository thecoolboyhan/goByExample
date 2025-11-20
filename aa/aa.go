package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//cc()
	//变量()
	//常量()
	//for循环()
	//if或Else()
	//Switch分支结构()
}

func Switch分支结构() {
	var i = 2
	fmt.Print("write", i, " as ")
	//普通switch,可以看出go语言switch不需要必须指定default
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//选择函数的返回值
	switch time.Now().Weekday() {
	//这几个感觉像是java的枚举类
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	var t = time.Now()
	//甚至可以不预选确定需要选择变量是哪个？试试是否可以实现多变量选择
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	//	可以使用多个不同变量，不具备传递性只要命中最上面的一条case，就自动break，不向后执行
	case t.Weekday() == time.Wednesday:
		fmt.Println("测试是否可以选择多个不同的变量")
	default:
		fmt.Println("It's after noon")
	}

	//内部类
	var whatAmI = func(i interface{}) {
		//i是参数，.(type)有点向反射获取参数类型，由于go不需要指定参数类型，动态语言
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			//这里展示了也可以使用%填充变量给字符串，但是需要使用printf
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}

func if或Else() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//不同于java的地方，可以在本次ifelse中声明一个只属于本次if else判断的变量
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func for循环() {
	i := 1
	//循环到3
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	//i也被修改了，只能使用一次
	fmt.Println(i)
	//金典fori
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//没有变量j会报错
	//fmt.Println(j)

	//while(true)
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
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

package main

import "fmt"

func main() {
	//数组()
	//切片()
	哈希Map()
}

func 哈希Map() {
	//创建一个map，key类型为String，value类型为int
	var m = make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	var v1 = m["k1"]
	fmt.Println("v1:", v1)
	//len=2，go好像没有size的概念
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//	获取一个不存在的key
	_, prs := m["k2"]
	//false，不存在的值，如果是存在的值，返回true
	fmt.Println("prs:", prs)
	//创建一个有初始值的map,map.of
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func 切片() {
	//创建一个空切片
	var s = make([]string, 3)
	fmt.Println("emp:", s)

	//给每个元素赋值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	//这里有区别于java中的集合，每次append会返回一个新的切片，需要用一个新的切片来接收
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	//	下面演示看看老的切片是什么样的
	var st = append(s, "g")
	//新的切片有值，老的并没有更改（思考：每次append切片都被复制了一份吗？）
	fmt.Println("new:", st, "\nold:", s)

	var c = make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//有点类似于sub，切分了2，3，4（没有5）
	var l = s[2:5]
	fmt.Println("sl1:", l)
	//又被赋了新的值
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	var t = []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//二维切片
	var twoD = make([][]int, 3)
	for i := 0; i < 3; i++ {
		var innerLen = i + 1
		//创建二维切片中的小切片
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

func 数组() {
	//创建一个大小为5的数组
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//二维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

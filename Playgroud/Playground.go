package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// %d 表示整型数字，%s 表示字符串
	var number = 123456789 // 声明一个变量
	var deaddate = "2022-07-03"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, number, deaddate)
	fmt.Println(target_url)

	var a error
	fmt.Println(a)

	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v%v%v%q\n", i, f, b, s)

	t := "tVc2333"
	y := t
	fmt.Println(t, y)
	fmt.Println(&t, &y) // 打印变量的内存地址

	const L int = 7 //声明一个常量
	const W int = 5
	var area int
	const c, v, n = 1, false, "VTB"

	area = L * W
	fmt.Printf("area=%d", area)
	print(c, v, n)

	const (
		q = "123456789"
		e = len(q)
		r = unsafe.Sizeof(q)
	)
	println(q, e, r)

	const (
		o = iota //iota 为特殊常量，可以用作枚举值
		u
		u1
		u2
	)
	println(o, u, u1, u2)

}

package main

import (
	"fmt"
)

func main() {
	// %d 表示整型数字，%s 表示字符串
	var number = 123456789
	var deaddate = "2022-07-03"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, number, deaddate)
	fmt.Println(target_url)
}

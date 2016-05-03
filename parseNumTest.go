package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt能够解析出16进制的数字
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// 还可以使用ParseUint函数
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi是解析10进制整型的快捷方法
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 解析函数在遇到无法解析的输入时，会返回错误
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}

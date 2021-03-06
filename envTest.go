package main

import "os"
import "strings"
import "fmt"

func main() {
	// 为了设置一个key/value对，使用`os.Setenv`
	// 为了获取一个key的value，使用`os.Getenv`
	// 如果所提供的key在环境变量中没有对应的value，
	// 那么返回空字符串
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 使用`os.Environ`来列出环境变量中所有的key/value对
	// 你可以使用`strings.Split`方法来将key和value分开
	// 这里我们打印所有的key
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], "\t", pair[1:])
	}
}

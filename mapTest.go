package main

import "fmt"

type A struct {
	x, y int
	str  string
}

func main() {

	map1 := make(map[string]string)
	val := map1["xx"]

	if val == "" {
		fmt.Println("val == \"\"")
	} else {
		fmt.Println("val != \"\"")
	}

	fmt.Println(val)

	map2 := make(map[int]int)
	val2 := map2[1]
	fmt.Println(val2)

	map3 := make(map[A]A)
	a := A{}
	val3 := map3[a]
	fmt.Println(val3)

	val, ok := map1["k2"]
	fmt.Println("ok == ", ok)
	fmt.Println("val == ", val)

}

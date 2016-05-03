package main

import (
	"fmt"
)

//指向同一底层数组的slice之间copy时，允许存在重叠。copy数组时，受限于src和dst数组的长度最小值。

func copyTest() {

	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := make([]int, 3)
	fmt.Println(s2)

	var n int
	n = copy(s2, s1)
	fmt.Println(n, s2, len(s2), cap(s2))

	s3 := s1[1:9]
	fmt.Println(n, s3, len(s3), cap(s3))

	n = copy(s3, s1[1:8])
	fmt.Println(n, s3, len(s3), cap(s3))

}

func appendTes() {
	s1 := make([]int, 3, 6)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	s2 := append(s1, 1, 2, 3)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	fmt.Println("s2= ", s2, len(s2), cap(s2))
	s3 := append(s2, 4, 5, 6)
	fmt.Println("s1= ", s1, len(s1), cap(s1))
	fmt.Println("s2= ", s2, len(s2), cap(s2))
	fmt.Println("s3= ", s3, len(s3), cap(s3))

}

func slice() {
	s := make([]string, 3)
	fmt.Println("emp :", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:  ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}

func main() {

	appendTes()
}

package main

import "fmt"

func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func addBase(x int) func(y int) int {

	return func(y int) int {
		return x + y
	}

}

func lazyFucTest() {

	var fs []func() int
	for i := 0; i < 10; i++ {
		j := i
		fs = append(fs, func() int {
			return j
		})
	}

	for _, f := range fs {
		fmt.Printf("%p = %v\n", f, f())
	}

}

func sumer() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}

}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())

	add10 := addBase(10)

	fmt.Println(add10(10))
	fmt.Println(add10(5))
	fmt.Println(add10(6))

	lazyFucTest()

	result := sumer()
	for i := 0; i < 10; i++ {
		fmt.Println(result(i))
	}

}

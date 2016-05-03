package main

import "fmt"

func split(sum int) (x, y int) {

	x = sum * 4 / 9
	y = sum - x
	return
}

type Callback func(x, y int) int

func callbackTest(x, y int, callback Callback) int {
	return callback(x, y)
}

func add(x, y int) int {
	return x + y
}

func main() {

	fmt.Println(split(17))

	x, y := 1, 3
	fmt.Println(callbackTest(x, y, add))

}

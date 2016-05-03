package main

import "fmt"
import "os"

func sum(nums ...int) {

	fmt.Println(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
func main() {

	for i, arg := range os.Args {
		fmt.Println(i, "\t", arg)
	}

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}

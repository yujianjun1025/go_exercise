package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func shooting(msg_chan chan string) {
	var group = 1
	for {
		for i := 1; i <= 10; i++ {
			msg_chan <- strconv.Itoa(group) + ":" + strconv.Itoa(i)
		}
		group++
		time.Sleep(time.Second * 10)
	}
}

func countA(msg_chan chan string) {
	for {
		fmt.Println(<-msg_chan)
	}
}

func channelBufferTest() {
	var c = make(chan string, 20)
	go shooting(c)
	go countA(c)

	//var input string
	//fmt.Scanln(&input)

}

func sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("Math: negative number pass to sqrt")
	}
	z := 0.0

	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * x)

	}
	return z, nil

}

func swap(x, y string) (string, string) {

	return y, x

}

func getSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Shape interface {
	area() float64
}

type CircleSelf struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (rectangle Rectangle) area() float64 {

	return rectangle.width * rectangle.height
}

func (circle CircleSelf) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func arrayTest(param []int) {

}

func getAverage(arr []int, size int) float32 {

	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {

		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg

}

func pointTest() {

	var a int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	var ptr *int

	fmt.Printf("The value of ptr is : %x\n", ptr)

	if ptr == nil {
		fmt.Println("ptr == nil")
	} else {
		fmt.Println("ptr != nil")
	}
}

func pArrayTest() {

	const MAX int = 3
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func structTest() {

	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	/* book 2 specification */
	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	/* print Book1 info */
	printBook(&Book1)

	/* print Book2 info */
	printBook(&Book2)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func sliceTest() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* print the original slice */
	fmt.Println("numbers ==", numbers)

	/* print the sub slice starting from index 1(included) to index 4(excluded)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* missing lower bound implies 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* missing upper bound implies len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* print the sub slice starting from index 0(included) to index 2(excluded) */
	number2 := numbers[:2]
	printSlice(number2)

	/* print the sub slice starting from index 2(included) to index 5(excluded) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func rangeTest() {

	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	/* print the numbers */
	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}

	/* create a map*/
	coutryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range coutryCapitalMap {
		fmt.Println("Capital of", country, "is", coutryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range coutryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}
}

func list_elem(n int) {

	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func goRouteTest() {

	go list_elem(10)
	//var input string
	//fmt.Scanf("%s", &input)
	//fmt.Println(input)
}

func fixed_shooting(msg_chan chan string) {
	for {
		msg_chan <- "fixed shooting"
		time.Sleep(time.Second * 1)
	}
}

func count(msg_chan chan string) {
	for {
		msg := <-msg_chan
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func strTest() {

	strs := []string{"x y", "x         y z", "    x", "y    "}

	for _, conv_query := range strs {
		conv_query = strings.Join(strings.Fields(conv_query), " ")
		seg_word_array := strings.Split(conv_query, " ")
		q1 := strings.Replace(conv_query, " ", "", -1)
		q2 := strings.Join(seg_word_array[1:], "")
		fmt.Println("conv_query:", conv_query, "q1 ==@", q1, "@ q2 ==@", q2, "@")
	}

}
func main() {

	pArrayTest()
	circle := CircleSelf{x: 0, y: 0, radius: 5}
	fmt.Println(circle.area())

	rect := Rectangle{width: 10, height: 100}

	fmt.Print("rect area :", getArea(rect), "\n")
	fmt.Print("circle area :", getArea(circle), "\n")

	var name string
	name = "xx"

	fmt.Println(name)
	for i := -10; i < 10; i++ {
		result, err := sqrt(float64(i))
		if err == nil {
			fmt.Println("normal: ", i, " result :", result)
		} else {
			fmt.Println("exception: ", i, " err : ", err)
		}
	}
	fmt.Printf("HelloWorld!\n")

	numbers := [6]int{1, 2, 3, 5}

	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

	a, b := swap("aa", "bb")
	fmt.Print(a, " \t", b, "\n")

	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Print(getSqrt(100), "\n")

	nextNumber := getSeq()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	var n [10]int /* n is an array of 10 integers */
	var i, j int

	/* initialize elements of array n to 0 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var a2 = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a2[i][j])
		}
	}

	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* pass array as an argument */
	avg = getAverage(balance, 5)

	/* output the returned value */
	fmt.Printf("Average value is: %f ", avg)

	fmt.Println("avg address %p", &avg)

	pointTest()
	structTest()

	sliceTest()
	rangeTest()

	goRouteTest()

	str1 := []string{"a", "b", "c"}

	for index, str := range str1 {
		fmt.Print(index, "\t", str, "\n")
	}

	strTest()
}

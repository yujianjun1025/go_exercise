package main

import (
	"fmt"
)

func main() {
	p := Person{2, "张三"}

	p.test(1)
	var f1 func(int) = p.test
	f1(2)
	Person.test(p, 3)
	var f2 func(Person, int) = Person.test
	f2(p, 4)

	s := Student{Person{2, "张三"}, 25}
	s.test()

}

type Person struct {
	Id   int
	Name string
}

func (this Person) test(x int) {
	fmt.Println("Id:", this.Id, "Name", this.Name)
	fmt.Println("x=", x)
}

//匿名字段, 模拟继承
type Student struct {
	Person
	Score int
}

func (this Student) test() {
	fmt.Println("student test,  id: ", this.Id, " name: ", this.Name)
}

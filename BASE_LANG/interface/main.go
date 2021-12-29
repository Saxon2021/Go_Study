package main

import (
	"fmt"
)

type animal interface {
	move()
	eat()
}

type dog struct {
	name string
	age  int
}

func (d dog) move() {
	fmt.Println(d.name, "汪汪汪~")
}

func (d dog) eat() {
	fmt.Println(d.name, "正在吃饭", "今年：", d.age)
}

func test(x animal) {
	x.move()
	x.eat()
}


func main() {

	dog_a := dog{
		name: "张三", age: 15,
	}

	test(dog_a)

}

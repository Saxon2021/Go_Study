package main

import (
	"fmt"
)

func f1() {
	fmt.Println("I'm f1")
}

func f2() int {
	return 10
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

// func f5(int, int) (x func(int) int) {
	
// }

func main() {

	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)

	fmt.Printf("%T\n", f4)
}

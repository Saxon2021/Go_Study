package main

import "fmt"

// // defer
// func study_defer_1() int {

// 	x := 5

// 	defer func() {
// 		x++
// 	}()

// 	return x

// }

// func study_defer_2() (x int) {

// 	defer func() {
// 		x++
// 	}()

// 	return 5

// }

// func study_defer_3() (y int) {

// 	x := 5
// 	defer func() {
// 		x++
// 	}()

// 	return x
// }

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	// ret1 := study_defer_1()
	// fmt.Println(ret1)

	// ret2 := study_defer_2()
	// fmt.Println(ret2)

	// ret3 := study_defer_3()
	// fmt.Println(ret3)

	a := 1
	b := 2

	defer calc("1", a, calc("10", a, b))  //3
	a = 0
	defer calc("2", a, calc("20", a, b))  //2
	b = 1
}

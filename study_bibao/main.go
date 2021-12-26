package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main(){
	jpgFunc := makeSuffixFunc(".jpg")
	pngFunc := makeSuffixFunc(".png")

	fmt.Println(jpgFunc("123"))
	fmt.Println(pngFunc("test.jpg"))
}

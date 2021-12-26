package main

import (
	"fmt"
	"strings"
)

// 分割字符串
func string_count(str string) {
	// 分割字符串
	array_item := strings.Split(str, " ")

	map_item := make(map[string]int, 10)

	for _, item := range array_item {

		if _, ok := map_item[item]; !ok {
			map_item[item] = 1
		} else {
			map_item[item]++
		}
	}

	for key, value := range map_item {
		fmt.Println(key, value)
	}
}

// 回文判断
func palindrome_judgment(str string) {

	for _, i := range str {

		fmt.Printf("%v\n", i)

	}

}




func main() {

	// 需要统计的字符串
	str1 := "How do you do"
	string_count(str1)

	str2 := "上海自来水来自海上"
	palindrome_judgment(str2)




}

package main

import (
	"fmt"
	"strings"
)

/**
*分割字符串函数
 */
func split(s string) (ret []string) {
	x1 := 0
	index := 0
	for i, v := range s {
		if v == ' ' {
			ret = append(ret, s[x1:i])
			x1 = i
			index++
		}
	}
	ret = append(ret, s[x1:])
	return
}

func main() {
	var s map[string]int
	s = make(map[string]int, 100)
	tempString := "how do you do"
	tempArray := strings.Split(tempString, " ")
	tempArray2 := split(tempString)
	fmt.Println(tempArray2)
	for _, v := range tempArray {
		s[v]++
	}
	for index, value := range s {
		fmt.Println(index, value)
	}
}

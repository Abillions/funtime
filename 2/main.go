package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// str := "e3[a2[b]]"
	str := "ef3[a]2[bc]gh"
	str = unzip(str)
	fmt.Println("str:", str)

}

func unzip(str string) string {
	var sub string
	for i := 0; i < len(str); i++ {
		start_index := strings.Index(str, "[")
		end_index := strings.Index(str, "]")
		if start_index == -1 || end_index == -1 {
			return str
		}
		count, _ := strconv.Atoi(string(str[start_index-1]))
		// fmt.Println(start_index, end_index, count)

		if strings.Index(str[start_index+1:end_index], "[") != -1 {
			str = strings.Replace(str, str[start_index+1:], unzip(str[start_index+1:]), 1)
		} else {
			sub = expand(count, str[start_index+1:end_index])
			str = strings.Replace(str, string(str[start_index-1:end_index+1]), sub, 1)
			i = end_index + 1
		}
	}
	// fmt.Println("str:", str)
	return str
}
func expand(n int, str string) string {
	if strings.Index(str, "[") != -1 {
		return ""
	}
	var s string
	for i := 0; i < n; i++ {
		s = s + str
	}
	return s
}

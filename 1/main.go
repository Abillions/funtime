package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []int{2, 3, 1, 5, 2, 1, 2, 4, 3, 2, 3}
	k := 3
	find(str, k)
}

func find(str []int, d int) {
	num := make(map[int]int)
	out := make(map[int]int)
	var arr []int
	for _, i := range str {
		if _, ok := num[i]; ok {
			num[i] = num[i] + 1
		} else {
			num[i] = 1
		}
	}
	for k, v := range num {
		arr = append(arr, v)
		out[v] = k
	}
	sort.Ints(arr)
	for i := 1; i <= d; i++ {
		fmt.Println(out[arr[len(arr)-i]])
	}
}

package main

import (
	"fmt"
	"strings"
)

func Reverse(arr []int) []int {
	n := len(arr)
	for i := 0; i < (n-1)/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}

func CountWords(str string) map[string]int {
	arr_s := strings.Split(str, " ")
	m := make(map[string]int)
	for i := 0; i < len(arr_s); i++ {
		_, ok := m[arr_s[i]]
		if ok == true {
			m[arr_s[i]]++
		} else {
			m[arr_s[i]] = 1
		}
	}
	return m
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(Reverse(arr))

	str := "кот собака кот кот мышь собака кот мышь собака кот собака мышь кот собака"
	fmt.Println(CountWords(str))

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var str string
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &str)
		l := len(str)
		if l > 10 {
			fmt.Println(string(str[0]) + strconv.Itoa(l-2) + string(str[l-1]))
		} else {
			fmt.Println(str)
		}
	}
}

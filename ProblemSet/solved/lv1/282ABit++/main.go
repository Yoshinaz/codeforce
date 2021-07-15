package main

import "fmt"

func main() {
	var n int
	var input string
	fmt.Scanf("%d\n", &n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &input)
		if string(input[1]) == "-" {
			ans--
		} else {
			ans++
		}
	}
	fmt.Println(ans)
}

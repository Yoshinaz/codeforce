package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	flag := [26]bool{}
	count := 0
	for _, v := range input {
		idx := int(v) - int('a')
		if !flag[idx] {
			flag[idx] = true
			count++
		}
	}
	if count%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}

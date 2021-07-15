package main

import "fmt"

func main() {
	var n, t int
	var input string
	fmt.Scanf("%d %d\n", &n, &t)
	fmt.Scanf("%s\n", &input)
	b := []byte(input)
	for i := 0; i < t; i++ {
		for j := 0; j < n-1; j++ {
			if b[j] == 'B' && b[j+1] == 'G' {
				b[j+1] = 'B'
				b[j] = 'G'
				j++
			}
		}
	}
	fmt.Println(string(b))
}

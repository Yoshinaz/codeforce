package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &s)
	one := 0
	zero := 0
	for _, v := range s {
		if v == 'n' {
			one++
		} else if v == 'z' {
			zero++
		}
	}
	for i := 0; i < one; i++ {
		fmt.Printf("1 ")
	}
	for i := 0; i < zero; i++ {
		fmt.Printf("0 ")
	}
	fmt.Println()
}

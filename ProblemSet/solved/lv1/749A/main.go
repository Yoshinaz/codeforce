package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	fmt.Println(n / 2)
	for i := 0; i < n/2-1; i++ {
		fmt.Printf("2 ")
	}
	if n%2 == 0 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}

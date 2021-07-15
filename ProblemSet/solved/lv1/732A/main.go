package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, r int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &k, &r)
	for i := 1; i <= 10; i++ {
		if k*i%10 == 0 || (k*i-r)%10 == 0 {
			fmt.Println(i)
			break
		}
	}
}

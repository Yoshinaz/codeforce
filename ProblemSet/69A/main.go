package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x, y, z, sx, sy, sz int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y, &z)
		sx += x
		sy += y
		sz += z
	}
	if sx == 0 && sy == 0 && sz == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

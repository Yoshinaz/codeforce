package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var r, c int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &r, &c)
	s1 := strings.Repeat("#", c)
	s2 := strings.Repeat(".", c-1) + "#"
	s3 := "#" + strings.Repeat(".", c-1)
	toggle := true
	for i := 0; i < r; i++ {
		if i%2 == 0 {
			fmt.Println(s1)
		} else {
			if toggle {
				fmt.Println(s2)
			} else {
				fmt.Println(s3)
			}
			toggle = !toggle
		}
	}
}

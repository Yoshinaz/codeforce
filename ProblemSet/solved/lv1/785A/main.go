package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	face := map[string]int{"Tetrahedron": 4, "Cube": 6, "Octahedron": 8, "Dodecahedron": 12, "Icosahedron": 20}
	var n int
	var s string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		ans += face[s]
	}
	fmt.Println(ans)
}

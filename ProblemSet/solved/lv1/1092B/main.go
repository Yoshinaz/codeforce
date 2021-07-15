package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	students := make([]Students, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &students[i].pro)
		students[i].idx = i + 1
	}
	sort.Slice(students, func(i, j int) bool {
		return students[i].pro < students[j].pro
	})
	sum := 0
	for i := 0; i < n; i += 2 {
		sum += students[i+1].pro - students[i].pro
	}
	fmt.Println(sum)
}

type Students struct {
	pro int
	idx int
}

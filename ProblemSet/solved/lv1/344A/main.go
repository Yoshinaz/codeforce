package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//92 ms
func main() {
	var n int
	var input string
	var last strings.Builder

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)
	last.WriteString("")
	count := 0
	//fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		//fmt.Scanf("%s\n", &input)
		scanner.Scan()
		input = scanner.Text()
		if input != last.String() {
			count++
			last.Reset()
			last.WriteString(input)
		}
	}
	c := count
	fmt.Println(c)
}

//124 ms
// func main() {
// 	var n int
// 	var input string
// 	var last strings.Builder

// 	in := bufio.NewReader(os.Stdin)
// 	fmt.Fscan(in, &n)
// 	last.WriteString("")
// 	count := 0
// 	//fmt.Scanf("%d\n", &n)
// 	for i := 0; i < n; i++ {
// 		//fmt.Scanf("%s\n", &input)
// 		fmt.Fscan(in, &input)
// 		if input != last.String() {
// 			count++
// 			last.Reset()
// 			last.WriteString(input)
// 		}
// 	}
// 	c := count
// 	fmt.Println(c)
// }

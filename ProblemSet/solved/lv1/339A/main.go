package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	split := strings.Split(input, "+")
	bucket := [3]int{}
	for _, v := range split {
		bucket[int(v[0])-int('1')]++
	}
	var output strings.Builder
	for i := 0; i < 3; i++ {
		for bucket[i] > 0 {
			output.WriteString(fmt.Sprintf("%d+", i+1))
			bucket[i]--
		}
	}
	fmt.Println(output.String()[0 : output.Len()-1])
}

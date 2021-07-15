package main

import (
	"fmt"
	"strings"
)

func main(){
	var a,b string
	fmt.Scanf("%s\n",&a)
	fmt.Scanf("%s\n",&b)
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	ans := 0
	if a<b {
		ans = -1
	}else if a>b{
		ans = 1
	}
	fmt.Println(ans)
}
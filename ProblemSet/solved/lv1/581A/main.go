package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	var r,b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&r,&b)
	min,max := minmax(r,b)
	fmt.Println(min,(max-min)/2)
}
func minmax(a,b int)(int,int){
	if a<b {
		return a,b
	}
	return b,a
}
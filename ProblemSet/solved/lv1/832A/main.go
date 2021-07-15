package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var a,b int64
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&a,&b)
	tmp := a/b
	if tmp%2==1{
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}
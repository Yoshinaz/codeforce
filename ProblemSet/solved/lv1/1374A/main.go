package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var t,x,y,n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&t)
	for i:=0;i<t;i++{
		fmt.Fscan(in,&x,&y,&n)
		tmp := (n/x)*x
		if tmp+y <=n{
			fmt.Println(tmp+y)
		}else{
			fmt.Println(tmp+y-x)
		}
	}
}
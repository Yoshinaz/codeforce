package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	in := bufio.NewReader(os.Stdin)
	var t,m,n,a int
	fmt.Fscan(in,&t)
	for i:=0;i<t;i++{
		fmt.Fscan(in,&m,&n)
		flag := make([]bool,101)
		for j:=0;j<m;j++{
			fmt.Fscan(in,&a)
			flag[a] = true
		}
		count :=0
		for j:=0;j<n;j++{
			fmt.Fscan(in,&a)
			if flag[a]{
				count++
			}
		}
		fmt.Println(count)
	}
}
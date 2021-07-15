package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var l,r int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&l,&r)
	ans := -1
	for i:=l;i<=r;i++{
		if isAllDiff(i){
			ans = i
			break
		}
	}
	fmt.Println(ans)
}

func isAllDiff(a int)bool{
	var flag [10]bool
	for ;a>0;{
		if flag[a%10]{
			return false
		}
		flag[a%10]=true
		a/=10
	}
	return true
}
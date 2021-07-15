package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var s,v1,v2,t1,t2 int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in,&s,&v1,&v2,&t1,&t2)
	a1 := 2*t1+s*v1
	a2 := 2*t2+s*v2
	if a1<a2{
		fmt.Println("First")
	}else if a2<a1{
		fmt.Println("Second")
	}else{
		fmt.Println("Friendship")
	}
}
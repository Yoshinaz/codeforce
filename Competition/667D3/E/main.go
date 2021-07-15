package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("tmp")
	for i := 0; i < 300000; i++ {
		file.WriteString(fmt.Sprintf("%d ", 1))
	}
	file.WriteString("2\n")
}

package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Scanln(&s)
	sl:=len(s)
	if sl == 0 {
		os.Exit(2)
	}
	cl:=sl / 2
	if sl%2==1 {
		cl--
	}
	for i:=0; i<cl; i++ {
		if s[i]!=s[sl-i-1] {
			os.Exit(1)
		}
	}
	os.Exit(0)
}

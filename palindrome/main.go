package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Scan(&str)
	strLength := len(str)
	if strLength == 0 {
		os.Exit(2)
	}

	compareLength := strLength / 2
	if strLength % 2 == 1 {
		compareLength--
	}

	for i := 0; i < compareLength; i++ {
		if str[i] != str[strLength - i - 1] {
			os.Exit(1)
		}
	}
	os.Exit(0)
}

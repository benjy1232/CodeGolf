package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var m [][]int
	var b []byte
	fmt.Scanln(&b)
	e := json.Unmarshal(b, &m)
	if e != nil {
		fmt.Println("error: ", e)
		os.Exit(1)
	}
	h := make(map[int]int)
	for j := 0; j < len(m[0]); j++ {
		h[m[0][j]] = 1
	}
	for i := 1; i < len(m); i++ {
		if len(m[i]) == 0 {
			return
		}
		a := make(map[int]int)
		for j := 0; j < len(m[i]); j++ {
			a[m[i][j]] = 1
		}
		for k := range a {
			if h[k] != 0 {
				h[k]++
			}
		}
	}
	var f []int
	for k := range h {
		if h[k] == len(m) {
			f = append(f, k)
		}
	}

	for l := 0; l < len(f); l++ {
		if l == len(f) - 1 {
			fmt.Print(f[l])
			break
		}
		fmt.Printf("%d, ", f[l])
	}
	fmt.Printf("\n")
}

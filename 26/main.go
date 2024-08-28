package main

import (
	"fmt"
	strings2 "strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	s = strings2.ToLower(s)

	set := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := set[r]; ok {
			fmt.Println("false")
			return
		}
		set[r] = struct{}{}
	}
	fmt.Println("true")
}

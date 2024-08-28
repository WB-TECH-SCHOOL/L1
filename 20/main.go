package main

import (
	"bufio"
	"fmt"
	"os"
	strings2 "strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	bytes, _, err := in.ReadLine()
	if err != nil {
		panic(err)
	}

	strings := strings2.Split(string(bytes), " ")
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}

	fmt.Println(strings2.Join(strings, " "))
}

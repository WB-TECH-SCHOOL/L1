package main

import "fmt"

func wantType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	var value interface{}
	wantType(value)

	value = 10
	wantType(value)

	value = "kk"
	wantType(value)

	value = true
	wantType(value)

	value = make(chan interface{})
	wantType(value)
}

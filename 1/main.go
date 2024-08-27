package main

import "fmt"

type Human struct {
	Name string
	Age  uint
}

func (h Human) String() string {
	return fmt.Sprintf("%s: %d", h.Name, h.Age)
}

type Action struct {
	Human
}

func main() {
	human := Human{
		Name: "Test",
		Age:  20,
	}
	fmt.Println(human.String())

	action := Action{
		Human: human,
	}
	fmt.Println(action.String())
}

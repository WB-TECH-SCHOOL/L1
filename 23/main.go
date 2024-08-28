package main

import (
	"errors"
	"fmt"
)

func removeByIndex[T any](slice []T, index int) ([]T, error) {
	if index > len(slice) || index < 0 {
		return nil, errors.New("incorrect index")
	}
	return append(slice[:index], slice[index+1:]...), nil
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var target int
	fmt.Scan(&target)

	fmt.Println(removeByIndex(nums, target))
}

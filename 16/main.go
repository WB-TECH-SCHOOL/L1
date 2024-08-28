package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	p := arr[0]

	var lower, upper []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < p {
			lower = append(lower, arr[i])
		} else {
			upper = append(upper, arr[i])
		}
	}
	lower = quickSort(lower)
	upper = quickSort(upper)
	return append(append(lower, p), upper...)
}

func main() {
	var nums = []int{15, 21, -88, 52, 28, -2, 37, -13, 14}

	fmt.Println(quickSort(nums))
}

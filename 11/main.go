package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	nums2 := []int{3, 5, -14, 8, -9, 11, 31, 18}

	num1Set := make(map[int]struct{})

	for _, num := range nums1 {
		num1Set[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok := num1Set[num]; ok {
			fmt.Println(num)
		}
	}
}

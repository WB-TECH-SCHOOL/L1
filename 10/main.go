package main

import "fmt"

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	tempsMap := make(map[int][]float64)

	for _, temp := range temps {
		key := int(temp/10) * 10
		tempsMap[key] = append(tempsMap[key], temp)
	}

	for key, val := range tempsMap {
		fmt.Printf("%d: %v\n", key, val)
	}
}

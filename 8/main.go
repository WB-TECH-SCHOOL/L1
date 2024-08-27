package main

import "fmt"

// Программа, устанавливающая конкретному биту значение с помощью XOR
// Если данный бит уже установлен, ничего не меняется
func main() {
	fmt.Println("Введите переменную, номер бита (с 1) и его желаемое значение (0/1)")
	var n, i, v int64
	fmt.Scan(&n, &i, &v)

	var mask int64 = 1 << (i - 1)
	fmt.Println("Введенное число в двоичной СС")
	fmt.Printf("%064b\n", n)
	fmt.Println("Маска в двоичной СС")
	fmt.Printf("%064b\n", mask)

	result := n ^ mask
	if v == 0 && result > n || v == 1 && result < n {
		result = n
	}

	fmt.Println("Полученное число в десятичной и двоичной СС")
	fmt.Printf("%064b %d", result, result)
}

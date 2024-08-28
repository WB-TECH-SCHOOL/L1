package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Проблема заключается в ссылке слайса на большой кусок памяти, в котором нет необходимости
// GC не будет чистить оставшийся кусок памяти, так как на него есть ссылка
// Чтобы её избежать нужно сделать новый объект типа строки и не ссылаться на

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Printf("%#v\n", *(*reflect.StringHeader)(unsafe.Pointer(&v)))
	// Проверка тупой теории о копировании байтов facepalm
	s := v[0:100]
	fmt.Printf("%#v\n", *(*reflect.StringHeader)(unsafe.Pointer(&s)))
	// Базовый подход - аналог копирования слайсов
	justString = string(append([]byte{}, v[0:100]...))
	fmt.Printf("%#v\n", *(*reflect.StringHeader)(unsafe.Pointer(&justString)))
}

func main() {
	someFunc()
}

package main

import "fmt"

const (
	a     = 42
	b int = 42
)

func main() {
	fmt.Println(a, "is untyped -- 컴파일러가 유동적으로 타입 지정")
	fmt.Println(b, "is typed")
}

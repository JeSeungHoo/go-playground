package main

import "fmt"

func main() {
	x := `원시 문자열 리터럴
엔터도 가능하다! 
\n 이런 거 인식 안한다!`
	fmt.Println(x)
}

package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
}

func sum(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func division(a int, b int) int {
	return a / b
}

func bitCounter(n uint) int {
	counter := 0
	for n != 0 {
		if n&1 == 1 {
			counter++
		}
		n >>= 1
	}
	return counter
}

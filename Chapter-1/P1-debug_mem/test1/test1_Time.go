package main

import "fmt"

func main() {
	fmt.Println(fibonacci(2))
}

func fibonacci(i int) int {
	if i < 2 {
		return 1
	}
	return fibonacci(i-2) + fibonacci(i-1)
}

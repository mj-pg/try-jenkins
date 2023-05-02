package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 == 1
}

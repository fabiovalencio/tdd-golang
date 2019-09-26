package main

import "fmt"

func sum(x, y int) int {
	result := x + y
	return result
}
func main() {
	result := sum(5, 5)
	fmt.Println("Sum :", result)
}

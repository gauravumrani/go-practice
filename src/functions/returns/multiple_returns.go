package main

import "fmt"

func getDouble(num int) (int, int) {
	return num, num * 2
}

func main() {
	original, doubled := getDouble(10)
	fmt.Println("The original is ", original)
	fmt.Println("The doubled is ", doubled)

	// if we don't want to use any of values we can bypass it with `_` operator
	_, do := getDouble(20)
	fmt.Println("The do is ", do)

	or, _ := getDouble(20)
	fmt.Println("The or is ", or)
}

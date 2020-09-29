package main

import "fmt"

func main() {
	fmt.Println("Hi")
	x := 19
	if x > 18 {
		fmt.Println("Eligible for vote")
	} else if x == 18 {
		fmt.Println("This is 18 just one more year")
	} else {
		fmt.Println("Not eligible")
	}
}

package main

import "fmt"

func main() {
	/* 	a := [5]int{7, 3, 23, 44, 35}
	   	numToFind := 3
			 searchNum(a, numToFind) */
	a := 4
	b := 5
	fmt.Println(addNum(a, b))
	fmt.Println(a)
	fmt.Println(b)
}

/* func searchNum(arr [5]int, numToFind int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == numToFind {
			fmt.Println(arr[i])
		}
	}
} */

func addNum(number1, number2 int) (int, int) {
	return number2, number1
}

package package1

import (
	"fmt"
)

// prints hello
func Hello() {
	fmt.Println("Hello")
}

func init() {
	fmt.Println("Initialize pacakge1!!!")
}

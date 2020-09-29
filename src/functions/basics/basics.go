package main

import "fmt"

func hello(name string) string {
	return "Hello " + name + " World"
}

func main() {
	fmt.Println(hello("Gaurav"))
}

package main

import (
	"fmt"
	"import-pkg/sub"
);

func main() {
	fmt.Println("sub.Constant\t", sub.FirstConstant)
	fmt.Println("sub.Constant\t", sub.Hot())
}
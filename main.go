package main

import "fmt"

var x int

type ownInt int

var y ownInt

func main() {

	x = 23
	fmt.Printf("%v\t", x)
	fmt.Printf("%T\n", x)

	y = 42
	fmt.Printf("%v\t", y)
	fmt.Printf("%T\n", y)
}

package main

import "fmt"

func main() {
	const pi float64 = 3.1415
	const (
		A int = 25
		B
		C float32 = 25.99
		D
	)

	const Hello = "Hello"
	const typedHello string = "Hello"
	fmt.Println(pi, A, B, C, D, Hello, typedHello)

	//iota
	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	fmt.Println(c0, c1, c2)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		_ //skip a day
		Add
	)
	fmt.Println(Sunday, Wednesday, Add)

	const (
		x         = iota * 42
		y float32 = iota * 42
		z         = iota * 42
	)
	fmt.Println(x, y, z)
}

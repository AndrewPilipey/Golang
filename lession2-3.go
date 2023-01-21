package main

import "fmt"

func main() {
	i := 3
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown number")
	}

	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("Default")
	}

	var c uint32
	fmt.Println("Enter a number from 1 to 6000")
	fmt.Scan(&c)
	switch {
	case 1 <= c && c <= 99:
		fmt.Println("from 1 to 99")
	case 100 <= c && c <= 999:
		fmt.Println("from 100 to 999")
	case 1000 <= c && c <= 6000:
		fmt.Println("from 1000 to 6000")
	}

	//next part is a task
	var a int
	fmt.Scan(&a)
	switch {
	case a == 0:
		fmt.Println("Ноль")
	case a < 0:
		fmt.Println("Число отрицательное")
	case a > 0:
		fmt.Println("Число положительное")
	}
}

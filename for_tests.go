package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if 10 <= b && b%8 == 0 && b <= 99 {
			sum += b
		}
	}
	fmt.Println(sum)
}

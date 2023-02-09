package main

import "fmt"

func main() {
	var sum int = 0

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)

	var sum2 = 0

	for i := 0; i <= 9; i++ {
		if i > 4 {
			break
		}
		sum2 += i
	}
	fmt.Println(sum2)
}

//just a lesson

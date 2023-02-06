package main

import "fmt"

func main() {
	var A int
	var B int
	var C int

	for fmt.Scan(&A, &B); A < B; A++ {
		C += A
	}
	fmt.Println(C + B)
}

/*
Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B).
Вывести сумму всех чисел от A до B  включительно.

Sample Input:
1 5
Sample Output:
15
*/

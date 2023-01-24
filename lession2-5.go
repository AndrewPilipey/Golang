package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Println("Enter a number. Not more than 10000.")
	fmt.Scan(&a)
	convlen := len(strconv.Itoa(a))

	switch {
	case convlen == 1:
		fmt.Println(a)
	case convlen == 2:
		fmt.Println(a / 10)
	case convlen == 3:
		fmt.Println(a / 100)
	case convlen == 4:
		fmt.Println(a / 1000)
	case a == 10000:
		fmt.Println(1)
	}
}

/*Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:
1234
Sample Output:
1*/

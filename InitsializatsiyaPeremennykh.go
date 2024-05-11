package main

import (
	fmt "fmt"
)

func main() {
	fmt.Println("Введите значение для переменной x")
	var x, y, z int
	fmt.Scan(&x)
	fmt.Println("Введите значение для переменной y")
	fmt.Scan(&y)
	fmt.Println("Введите значение для переменной z")
	fmt.Scan(&z)
	fmt.Println(x)
	fmt.Println("Вы ввели  ", "x:", x, "y:", y, "z:", z)
}

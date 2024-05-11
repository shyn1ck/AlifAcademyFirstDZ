package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var x, y int
	fmt.Println("Введите значение x")
	fmt.Scan(&x)
	fmt.Println("Введите значение y")
	fmt.Scan(&y)
	fmt.Println("Начальные значение: x =", x, ", y =", y)
	x = x + y
	y = x - y
	x = x - y
	fmt.Println("Конечные значение: x =", x, ", y =", y)
}

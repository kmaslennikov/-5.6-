package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Вводите коэффициенты a, b и c квадратного уравнения:")
	var a, b, c float64
	var x1, x2 float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	D := math.Pow(b, 2) - (4 * a * c)

	if D < 0 {
		fmt.Println("Дескриминант меньше 0")
		fmt.Println("Корней нет")
	} else if D == 0 {
		x1 = b / (2 * a)
		fmt.Println("Дескриминант равен 0")
		fmt.Println("Один корень уравнения:", x1)
	} else {
		x1 = (-b - math.Sqrt(D)) / (2 * a)
		x2 = (-b + math.Sqrt(D)) / (2 * a)
		fmt.Println("Дескриминант больше 0")
		fmt.Println("Первый корень уравнения:", x1)
		fmt.Println("Второй корень уравнения:", x2)
	}

}

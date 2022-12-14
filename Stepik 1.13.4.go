/*
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/

package main

import "fmt"

func main() {
	// Инициализация переменных
	var a, b, c int
	// Присвоение значений
	fmt.Scan(&a, &b, &c)

	switch {
	// в случае если a*a+b*b == c*c
	case a*a+b*b == c*c:
		// Вывод
		fmt.Println("Прямоугольный")
	case a*a+b*b >= c*c:
		// Вывод
		fmt.Println("Тупоугольный")
	case a*a+b*b <= c*c:
		// Вывод
		fmt.Println("Остроугольный")
	case a*a == b*b && b*b == c*c && a*a == c*c:
		// Вывод
		fmt.Println("Равнобедренный")
	// в противном случае
	default:
		// Вывод
		fmt.Println("Неизвестный в природе")
	}
}

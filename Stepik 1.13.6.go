/*
Даны два числа. Найти их среднее арифметическое.

Формат входных данных
На вход дается два целых положительных числа a и b.

Формат выходных данных
Программа должна вывести среднее арифметическое чисел a и b (ответ может быть целым числом или дробным)
*/

package main

import "fmt"

func main() {
	// Инициализация переменных
	var a, b, average float64
	// Присвоение значений
	fmt.Scan(&a, &b)
    // Вычисление среднего арифмитического
	average = (a + b) / 2
	// Вывод
    fmt.Println(average)
}

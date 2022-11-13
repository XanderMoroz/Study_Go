/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.
*/

package main

import "fmt"

func main() {
	// Инициализация переменных
	var N, dig, count int
	// Присвоение значений
	fmt.Scan(&N)
	// В цикле:
	for i := 0; i < N; i++ {
		// Извлекаем число
		fmt.Scan(&dig)
		// Если число равно нулю
		if dig == 0 {
			// Инкремент счетчика
			count++
		}
	}
	// Выводим количество нулей
	fmt.Println(count)
}

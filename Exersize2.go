// Перед вами классическое задание для собеседований на любом языке программирования — FizzBuzz.
// Напишите программу, которая выводит на экран числа от 1 до 100.
// При этом вместо чисел, кратных трём, программа должна выводить слово Fizz, а вместо кратных пяти — Buzz.
// Если число кратно и трём, и пяти, программа должна выводить слово FizzBuzz.

package main

import "fmt"

func main() {
	for elem := 1; elem <= 100; elem++ {
		switch {
		case elem%3 == 0:
			fmt.Println("Fizz")
		case elem%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(elem)
		}
	}
}

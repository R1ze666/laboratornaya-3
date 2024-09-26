package main

import (
	"fmt"
	"laba3/mathutils"
	"laba3/stringutils"
)

func main() {

	var num int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&num)
	result := mathutils.Factorial(num)
	fmt.Printf("Факториал числа %d = %d\n", num, result)

	var input string
	fmt.Print("Введите строку для переворота: ")
	fmt.Scan(&input)
	reversed := stringutils.Reverse(input)
	fmt.Printf("Перевёрнутая строка: %s\n", reversed)

	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Массив из 5 целых чисел:")
	for _, num := range arr {
		fmt.Println(num)
	}

	slice := arr[:]
	fmt.Println("Исходный срез:", slice)
	slice = append(slice, 60)
	fmt.Println("Срез после добавления элемента:", slice)
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Срез после удаления второго элемента:", slice)

	strings := []string{"1234", "12345q", "52525qqq", "pine123e", "qwfqgabsbmm78"}
	var longest string
	for _, str := range strings {
		if len(str) > len(longest) {
			longest = str
		}
	}
	fmt.Printf("Самая длинная строка: %s\n", longest)
}

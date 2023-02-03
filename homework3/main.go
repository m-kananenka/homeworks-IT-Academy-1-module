package main

import (
	"fmt"
	"homework/homework3/internal"
)

func main() {
	getNum(11)        //Да
	getString("шест") //я не знаю
	getOdd(13)        //1 3 5 7 9 11 13
	fmt.Println()
	internal.Recursion(1, 10)

}

/*
1.Написать функцию, принимающую аргумент типа int. Выводить на экран "Да", если аргумент больше 10 и меньше 20,
"Нет" если аргумент меньше 10, "Наверное" если аргумент больше 20.
*/
func getNum(i int) {
	if i >= 10 && i <= 20 {
		fmt.Println("Да")
	} else if i < 10 {
		fmt.Println("Нет")
	} else if i > 20 {
		fmt.Println("Наверное")
	}
}

/*
2.Написать функцию, принимающую string в качестве аргумента. Значение аргумента это написанные на русском языке
числа "один", "два" и так далее, до "десять". Выводить на экран соответствующее число. аргумент не совпадает с
описанными значениями, выводить "Не знаю"
*/
func getString(s string) {
	switch s {
	case "один":
		fmt.Println(1)
	case "два":
		fmt.Println(2)
	case "три":
		fmt.Println(3)
	case "четыре":
		fmt.Println(4)
	case "пять":
		fmt.Println(5)
	case "шесть":
		fmt.Println(6)
	case "семь":
		fmt.Println(7)
	case "восемь":
		fmt.Println(8)
	case "девять":
		fmt.Println(9)
	case "десять":
		fmt.Println(10)
	default:
		fmt.Println("Я не знаю")
	}
}

// 3.Написать функцию, выводящую на экран все нечетные числа от 0 до аргумента типа int, переданного в функцию.
func getOdd(v int) {
	for i := 0; i <= v; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}

	}
}

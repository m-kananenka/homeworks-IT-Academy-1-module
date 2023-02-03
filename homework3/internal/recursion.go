package internal

import "fmt"

/*4.Написать функцию, принимающие 2 аргумента типа int. Вывести на экран все числа от A до B если A<B, иначе вывести
все числа от B до A. Использовать рекурсию.*/
//5.Создать новый пакет и перенести в него любую из написанных функций. Вызывать функцию в main.
func Recursion(a, b int) {
	if a < b {
		if a >= b {
			return
		}
		fmt.Print(a, " ")
		a++
		Recursion(a, b)
	} else if a > b {
		if a <= b {
			return
		}
		fmt.Print(b, " ")
		b++
		Recursion(a, b)
	}
}

//I would solve this task this way :)

//	if a < b {
//		for i := a; i <= b; i++ {
//			fmt.Print(i, " ")
//		}
//	} else {
//		for i := a; i >= b; i-- {
//			fmt.Print(i, " ")
//		}
//	}

package main

import (
	"fmt"
)

func main() {
	sortArray([]int{1, 3, 4, 5, 2}) //[1 2 3 4 5]

	ints := []int{1, 2, 3, 4, 5}
	search, b := binarySearch(ints, 5)
	fmt.Println(search, b)

	c := сompareString("aaab", "aaac")
	fmt.Println(c)

	printStairs(10)
}

// 1.Реализовать сортировку массива целых чисел любым методом
func sortArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				var tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp

			}

		}

	}
	fmt.Println(arr)
}

// 2. Реализовать двоичный поиск в отсортированном массиве целых чисел
func binarySearch(a []int, target int) (int, bool) {
	l, r := 0, len(a)-1
	for l <= r {
		m := (l + r) / 2
		if a[m] > target {
			r = m - 1
		} else if a[m] < target {
			l = m + 1
		} else {
			return m, true
		}
	}
	return 0, false
}

/*
3.Даны 2 строки, состоящие из букв латинского алфавита. Проверить, можно ли получить одну строку из другой,
перестановкой букв. (Задача очень легко решается с помощью мапы)
*/
func сompareString(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[byte]int)

	for v := range s1 {
		target := s1[v]
		_, ok := m[target]
		if ok {
			i := m[target]
			i++
			m[target] = i
		} else {
			m[target] = 1
		}
	}

	for v := range s2 {
		target := s2[v]
		_, ok := m[target]
		if ok {
			i := m[target]
			i--
			m[target] = i
			if i < 0 {
				return false
			}
		} else {
			return false

		}
	}
	return true
}

/*
4.Реализовать функцию, принимающую целое число n. Функция создает слайс слайсов размером nxn. Элементами слайса
являются слайсы, элементами которых являются числа от 0 до n. Вывести на экран слайс в виде лестницы.
*/
func printStairs(n int) {
	s := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		s2 := make([]int, 0)
		for j := 0; j < n; j++ {
			s2 = append(s2, j)
		}
		s = append(s, s2)
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

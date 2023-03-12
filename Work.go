package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	splittedString := strings.Split(text, " ")
	if len(splittedString) <= 1 {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		os.Exit(1)
	}
	if len(splittedString) != 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		os.Exit(1)
	}
	operator := "+ - / *"
	if !strings.Contains(operator, splittedString[1]) {
		fmt.Println("Такой оператор не поддерживается")
		os.Exit(1)
	}

	myMap := make(map[int]string)
	myMap[1] = "I"
	myMap[2] = "II"
	myMap[3] = "III"
	myMap[4] = "IV"
	myMap[5] = "V"
	myMap[6] = "VI"
	myMap[7] = "VII"
	myMap[8] = "VIII"
	myMap[9] = "IX"
	myMap[10] = "X"
	var r int = 0
	var c int = 0
	x := "I II III IV V VI VII VIII IX X"
	newx := strings.Split(x, " ")
	a := "1 2 3 4 5 6 7 8 9 10"
	ch1, _ := strconv.Atoi(splittedString[0])
	ch2, _ := strconv.Atoi(splittedString[2])

	if (strings.Contains(a, splittedString[0]) && strings.Contains(a, splittedString[2])) || (strings.Contains(x, splittedString[0]) && strings.Contains(x, splittedString[2])) {
		if strings.Contains("+", splittedString[1]) && (strings.Contains(a, splittedString[0]) && strings.Contains(a, splittedString[2])) {
			fmt.Println(ch1 + ch2)
		} else if strings.Contains("-", splittedString[1]) && (strings.Contains(a, splittedString[0]) && strings.Contains(a, splittedString[2])) {
			fmt.Println(ch1 - ch2)
		} else if strings.Contains("/", splittedString[1]) && (strings.Contains(a, splittedString[0]) && strings.Contains(a, splittedString[2])) {
			fmt.Println(float64(ch1) / float64(ch2))
		} else if strings.Contains("*", splittedString[1]) && (strings.Contains(a, splittedString[0]) && strings.Contains(a, splittedString[2])) {
			fmt.Println(ch1 * ch2)
		}

		if (splittedString[0] == "I" || splittedString[0] == "II" || splittedString[0] == "III" || splittedString[0] == "IV" || splittedString[0] == "V" || splittedString[0] == "VI" || splittedString[0] == "VII" || splittedString[0] == "VIII" || splittedString[0] == "IX" || splittedString[0] == "X") && (splittedString[2] == "I" || splittedString[2] == "II" || splittedString[2] == "III" || splittedString[2] == "IV" || splittedString[2] == "V" || splittedString[2] == "VI" || splittedString[2] == "VII" || splittedString[2] == "VIII" || splittedString[2] == "IX" || splittedString[2] == "X") {
			for i := range newx {
				if newx[i] == splittedString[0] {
					r = i + 1
				}
			}
			for i := range newx {
				if newx[i] == splittedString[2] {
					c = i + 1
				}
			}
			if strings.Contains("+", splittedString[1]) {
				summ := r + c
				if summ > 10 {
					summ1 := summ / 10
					summ2 := summ % 10
					fmt.Println(strings.Repeat(myMap[10], summ1) + myMap[summ2])
				} else {
					fmt.Println(myMap[summ])
				}
			} else if strings.Contains("*", splittedString[1]) {
				umn := r * c
				if umn > 10 {
					umn1 := umn / 10
					umn2 := umn % 10
					fmt.Println(strings.Repeat(myMap[10], umn1) + myMap[umn2])
				} else {
					fmt.Println(myMap[umn])
				}
			} else if strings.Contains("/", splittedString[1]) {
				del := r / c
				if del < 1 {
					fmt.Println("Исключение, ответ меньше единицы")
					os.Exit(1)
				} else {
					fmt.Println(myMap[del])
				}
			} else if strings.Contains("-", splittedString[1]) {
				min := r - c
				if r < c {
					fmt.Println("Ошибка, так как в римской системе нет отрицательных чисел.")
					os.Exit(1)
				} else if r == c {
					fmt.Println("Исключение, ответ меньше единицы.")
					os.Exit(1)
				} else {
					fmt.Println(myMap[min])
				}
			}
		}
	} else {
		fmt.Println("Не подходящее число, вводить можно только целые числа от 1 до 10 включительно, либо арабской системы счисления либо римской!")
		os.Exit(1)
	}
}

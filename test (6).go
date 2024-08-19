package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s := strings.Split(text, " ")
	if strings.ContainsAny(text, "."){ //Проверка на то, чтобы не было десятичных дробей, либо не было где-то лишних точек.
		fmt.Println("Паника. Операции проводятся только с целыми числами или неправильный формат ввода")
	} else if len(s) == 3{
		number_1 := s[0] //первое введенное число
		operation := s[1] // операция между числами
		number_2 := s[2] // второе введенное число
		number_2 = strings.TrimSpace(number_2)// убрал \n в конце
		if strings.ContainsAny(number_1, "0123456789") && !strings.ContainsAny(number_1, "IVXLC") && //проверка на то, чтобы в первом и во втором числе были только арабские цифры. чтобы не было ввода по типу: 1V + 39XX
		 strings.ContainsAny(number_2, "0123456789") && !strings.ContainsAny(number_2, "IVXLC"){
			toNumber_1, _ := strconv.Atoi(number_1) //перевод первого числа из string в int
			toNumber_2, _ := strconv.Atoi(number_2) //перевод второго числа из string в int
			if toNumber_1 >= 1 && toNumber_1 <= 10 && toNumber_2 >= 1 && toNumber_2 <= 10{ //проверка чтобы числа были в диапазоне от 1 до 10
				switch operation{
				case "+":
					fmt.Println(toNumber_1, "+", toNumber_2, "=", toNumber_1 + toNumber_2)
				case "-":
					fmt.Println(toNumber_1, "-", toNumber_2, "=", toNumber_1 - toNumber_2)
				case "*":
					fmt.Println(toNumber_1, "*", toNumber_2, "=", toNumber_1 * toNumber_2)
				case "/":
					fmt.Println(toNumber_1, "/", toNumber_2, "=", toNumber_1 / toNumber_2)
				default:
					fmt.Println("Паника. Нет такой арифметической операции или формат математической операции не удовлетворяет заданию.")
				}
			} else {
				fmt.Println("Паника. Введенные цифры должны находиться в диапазоне от 1 до 10 включительно.")
			}
		 } else if !strings.ContainsAny(number_1, "0123456789") && strings.ContainsAny(number_1, "IVXLC") && //проверка на то, чтобы в первом и во втором числе были только римские цифры. чтобы не было ввода по типу: 1V2 + X4X
		 !strings.ContainsAny(number_2, "0123456789") && strings.ContainsAny(number_2, "IVXLC"){
			romeToArab := map[string]int{
				"I":  1,
				"II": 2,
				"III": 3, 
				"IV": 4,
				"V": 5,
				"VI": 6,
				"VII": 7,
				"VIII": 8,
				"IX": 9,
				"X": 10,
				"XII": 12,
				"XIV": 14,
				"XV": 15,
				"XVI": 16,
				"XVIII": 18,
				"XX": 20,
				"XXI": 21,
				"XXIV": 24,
				"XXV": 25,
				"XXVII": 27,
				"XXVIII": 28,
				"XXX": 30,
				"XXXII": 32,
				"XXXV": 35,
				"XXXVI": 36,
				"XL": 40,
				"XLII": 42,
				"XLV": 45,
				"XLVIII": 48, 
				"XLIX": 49,
				"L": 50,
				"LIV": 54,
				"LVI": 56,
				"LX": 60,
				"LXIII": 63,
				"LXIV": 64,
				"LXX": 70,
				"LXXII": 72,
				"LXXX": 80,
				"LXXXI": 81, 
				"XC": 90, 
				"C": 100,
			}
			arabToRome := map[int]string{
				1: "I",
				2: "II",
				3: "III", 
				4: "IV",
				5: "V",
				6: "VI",
				7: "VII",
				8: "VIII",
				9: "IX",
				10: "X",
				12: "XII",
				14: "XIV",
				15: "XV",
				16: "XVI",
				18: "XVIII",
				20: "XX",
				21: "XXI",
				24: "XXIV",
				25: "XXV",
				27: "XXVII",
				28: "XXVIII",
				30: "XXX",
				32: "XXXII",
				35: "XXXV",
				36: "XXXVI",
				40: "XL",
				42: "XLII",
				45: "XLV",
				48: "XLVIII", 
				49: "XLIX",
				50: "L",
				54: "LIV",
				56: "LVI",
				60: "LX",
				63: "LXIII",
				64: "LXIV",
				70: "LXX",
				72: "LXXII",
				80: "LXXX",
				81: "LXXXI", 
				90: "XC", 
				100: "C",
			}
			arabNumber_1 := romeToArab[number_1]
			arabNumber_2 := romeToArab[number_2]
			if arabNumber_1 >= 1 && arabNumber_1 <= 10 && arabNumber_2 >= 1 && arabNumber_2 <= 10{ // проверка чтобы числа были в диапазоне от 1 до 10
				switch operation{
				case "+":
					fmt.Println(number_1, "+", number_2, "=", arabToRome[arabNumber_1 + arabNumber_2])
				case "-":
					if arabNumber_1 - arabNumber_2 > 0{ // проверка чтобы результат был положительным
						fmt.Println(number_1, "-", number_2, "=", arabToRome[arabNumber_1 - arabNumber_2])
					} else {
						fmt.Println("Паника. В римской системе счисления нет отрицательных чисел и ноля.")
					}
				case "*":
					fmt.Println(number_1, "*", number_2, "=", arabToRome[arabNumber_1 * arabNumber_2])
				case "/":
					if arabNumber_1 / arabNumber_2 > 0{ //проверка чтобы результат не был меньше единицы например: [V / X]
						fmt.Println(number_1, "/", number_2, "=", arabToRome[arabNumber_1 / arabNumber_2])
					} else {
						fmt.Println("Паника. В римской системе счисления нет чисел меньше единицы")
					}
				default:
					fmt.Println("Паника. Нет такой арифметической операции или формат математической операции не удовлетворяет заданию")
				}
			} else {
				fmt.Println("Паника. Введенные цифры должны находиться в диапазоне от 1 до 10 включительно.")
			}			
		 } else if strings.ContainsAny(number_1, "0123456789") && !strings.ContainsAny(number_1, "IVXLC") &&//проверка на то чтобы в одном выражении были разные системы счисления.
		 !strings.ContainsAny(number_2, "0123456789") && strings.ContainsAny(number_2, "IVXLC") ||
		  !strings.ContainsAny(number_1, "0123456789") && strings.ContainsAny(number_1, "IVXLC") &&
		 strings.ContainsAny(number_2, "0123456789") && !strings.ContainsAny(number_2, "IVXLC"){
			fmt.Println("Паника. Одновременно разные системы счисления использоваться не могут.")
		 } else{
			fmt.Println("Неправильный формат ввода.")
		 }			
	} else if len(s) < 3 && !strings.ContainsAny(text, "+-*/"){
		fmt.Println("Паника. Строка не является математической операцией.")
	} else if len(s) > 3 {
		fmt.Println("Паника. Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	} else{
		fmt.Println("Неправильный формат ввода.")
	}
}
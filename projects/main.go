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
	fmt.Println("Input:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	var numb []string = strings.Split(text, " ")

	Scaner(numb)
}

func divide(x []string) (int, int) {
	a, _ := strconv.Atoi(x[0])
	b, _ := strconv.Atoi(x[2])
	return a, b
}

func calculet(a int, b int, c string) int {
	var d int
	if c == "*" {
		d = a * b
	} else if c == "/" {
		d = (a - (a % b)) / b
	} else if c == "+" {
		d = a + b
	} else if c == "-" {
		d = a - b
	} else {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	return d

}

func RimNumb(a string, b string, c string) string {
	var a1, b1, res, res1, res2 int
	var Result string

	rim := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	ResRim := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		20:  "XX",
		30:  "XXX",
		40:  "XL",
		50:  "L",
		60:  "LX",
		70:  "LXX",
		80:  "LXXX",
		90:  "XC",
		100: "C",
	}

	if value, ok := rim[a]; ok {
		a1 = value

	} else {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	if value, ok := rim[b]; ok {
		b1 = value
	} else {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	res = calculet(a1, b1, c)
	if res < 0 {
		panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
	}
	if res > 10 {
		res1 = res - (res % 10)
		res2 = res % 10
		Result = ResRim[res1] + ResRim[res2]
	} else {
		Result = ResRim[res]
	}

	return Result
}

func Scaner(x []string) {
	if len(x) > 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if len(x) < 3 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	rim := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arab := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, value := range rim {
		if value == x[0] {
			for _, value := range rim {
				if value == x[2] {
					fmt.Println(RimNumb(x[0], x[2], x[1]))

				}
				for _, value := range arab {
					if value == x[2] {
						panic("Выдача паники, так как используются одновременно разные системы счисления.")
					}

				}
			}

		}

	}

	for _, value := range arab {
		if value == x[0] {
			for _, value := range arab {
				if value == x[2] {
					a, b := divide(x)
					fmt.Println(calculet(a, b, x[1]))

				}

				for _, value := range rim {
					if value == x[2] {
						panic("Выдача паники, так как используются одновременно разные системы счисления.")
					}
				}
			}
		}
	}
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("input your number and operator")
	for {
		var input1 string
		var input2 string
		var operator string

		fmt.Println("Enter first number:")
		fmt.Scan(&input1)

		num1, err1 := strconv.Atoi(input1)
		if err1 != nil {
			fmt.Println("error: must be a number")
			continue
		}

		fmt.Println("Enter Second number:")
		fmt.Scan(&input2)

		num2, err2 := strconv.Atoi(input2)
		if err2 != nil {
			fmt.Println("error: must be a number")
			continue
		}

		fmt.Println("operator: [1] Add [2] Substract [3] Multiply [4] Division [5] Quit [6]Help")
		fmt.Scan(&operator)
		switch operator {
		case "1":
			fmt.Println(num1 + num2)

		case "2":
			fmt.Println(num1 - num2)
		case "3":
			fmt.Println(num1 * num2)
		case "4":
			if num2 == 0 {
				fmt.Println("error: try again")
				continue
			}
			fmt.Println(num1 / num2)

		case "5":
			fmt.Println("Thanks for your time")
			continue

		case "help":
			fmt.Println("available command")
			fmt.Println("1 ---> add two number")
			fmt.Println("2 ---> substract two number")
			fmt.Println("3 ---> multiply two number")
			fmt.Println("4 ---> divide two number")
			fmt.Println("5 ---> quit")
			fmt.Println("Example:")
			fmt.Println("  Enter: 10, 5, then choose 1 → Result: 15")

		default:
			fmt.Println("Not a valid input, type 'help' for guidance")

		}

	}

}

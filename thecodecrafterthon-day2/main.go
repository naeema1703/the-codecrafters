package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("\nSelect an operation to continue:")
		fmt.Println("1. Hex to Decimal")
		fmt.Println("2. Binary to Decimal")
		fmt.Println("3. Decimal to Binary/Hex")
		fmt.Println("4. Exit")
		fmt.Print("code: ")

		var code string
		fmt.Scanln(&code)

		switch code {
		case "1":
			fmt.Print("Enter hexadecimal number: ")
			var input string
			fmt.Scanln(&input)

			decimal, err := strconv.ParseInt(input, 16, 64)
			if err != nil {
				fmt.Println("Invalid hexadecimal input. Try again.")
				continue
			}
			fmt.Println("Decimal result:", decimal)

		case "2":
			fmt.Print("Enter binary number: ")
			var input string
			fmt.Scanln(&input)

			decimal, err := strconv.ParseInt(input, 2, 64)
			if err != nil {
				fmt.Println("Invalid binary input. Try again.")
				continue
			}
			fmt.Println("Decimal result:", decimal)

		case "3":
			fmt.Print("Enter decimal number: ")
			var input string
			fmt.Scanln(&input)

			decimal, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Invalid decimal input. Try again.")
				continue
			}

			binary := strconv.FormatInt(decimal, 2)
			hexadecimal := strings.ToUpper(strconv.FormatInt(decimal, 16))

			fmt.Println("Binary:", binary)
			fmt.Println("Hexadecimal:", hexadecimal)

		case "4":
			fmt.Println("Goodbye")
			return

		default:
			fmt.Println("Invalid code. Please enter 1, 2, 3, or 4.")
		}
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	var num1, num2 int
	var ch string
	fmt.Println("Choose operation to perform (+,-,*,/)")
	fmt.Scan(&ch)
	if ch == "+" || ch == "-" || ch == "*" || ch == "/" {
		_, err := fmt.Scan(&num1, &num2)
		if err != nil {
			fmt.Println("Invalid number format;", err)
			os.Exit(0)
		}
		switch ch {
		case "+":
			fmt.Println(num1 + num2)
			break
		case "-":
			fmt.Println(num1 - num2)
			break
		case "*":
			fmt.Println(num1 * num2)
			break
		case "/":
			fmt.Println(num1 / num2)
			break
		default:
			fmt.Println("invalid input\n use appropriate values")

		}
	} else {
		fmt.Println("use valid operations")
	}
}

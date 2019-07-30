package main

import "fmt"
func main(){
var num1,num2 float64
fmt.Println("enter the first number")
fmt.Scan(&num1)
fmt.Println("enter the second number")
fmt.Scan(&num2)
var ch string 
fmt.Println("Choose operation to perform (+,-,*,/)")
fmt.Scan(&ch)
switch ch{ 	
	case "+":
		fmt.Println(num1+num2)
		break
	case "-":
		fmt.Println(num1-num2)
		break
	case "*":
		fmt.Println(num1*num2)
		break
	case "/":
		fmt.Println(num1/num2)
		break	
	default:
		fmt.Println("error")
	
}

}
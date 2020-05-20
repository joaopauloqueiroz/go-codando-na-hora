package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Qual operação deseja realizar? \n \n ")
	fmt.Printf("1: Adição \t \t 2: Subtração \n 3: Multiplicação \t 4: Divisão \n ")

	operation, _ := reader.ReadString('\n')
	operationNumber, _ := strconv.ParseInt(strings.TrimSpace(operation), 10, 64)

	fmt.Printf("Qual o primeiro valor para a operação? ")
	x, _ := reader.ReadString('\n')

	fmt.Printf("Qual o segundo valor para a operação? ")
	y, _ := reader.ReadString('\n')

	xNumber, _ := strconv.ParseFloat(strings.TrimSpace(x), 64)
	yNumber, _ := strconv.ParseFloat(strings.TrimSpace(y), 64)

	if operationNumber == 1 {
		result := sum(xNumber, yNumber)
		fmt.Printf("O resultado da operação de adição é: %.2f", result)
	} else if operationNumber == 2 {
		result := sub(xNumber, yNumber)
		fmt.Printf("O resultado da operação de subtração é: %.2f", result)
	} else if operationNumber == 3 {
		result := multiply(xNumber, yNumber)
		fmt.Printf("O resultado da operação de multiplicação é: %.2f", result)
	} else if operationNumber == 4 {
		result := division(xNumber, yNumber)
		fmt.Printf("O resultado da operação de divisão é: %.2f", result)
	} else {
		fmt.Printf("Operação invalida")
	}	
}

func sum(num1 float64, num2 float64) float64 {
	return (num1 + num2)
}

func sub(num1 float64, num2 float64) float64 {
	return (num1 - num2)
}

func division(num1 float64, num2 float64) float64 {
	return (num1 / num2)
}

func multiply(num1 float64, num2 float64) float64 {
	return (num1 * num2)
}
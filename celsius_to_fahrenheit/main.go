package main

import "fmt"

func main() {
	var celsius int

	fmt.Print("Digite uma temperatura em Calsius: ")
	fmt.Scan(&celsius)

	var fahrenheit = (celsius * 9/5) + 32

	fmt.Printf("Convertido para fahrenheit: %d\n", fahrenheit)
}
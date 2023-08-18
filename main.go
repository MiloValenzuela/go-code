package main

import "fmt"


func main() {
	saludo := hello("Milo")
	fmt.Println(saludo)
	sum, mul := calc(4, 5)
	fmt.Println("La suma es: ", sum)
	fmt.Println("La multipicacion es: ", mul)
}

func hello(name string) string {
	return "Hola, " + name
}

func calc(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}
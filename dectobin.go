package main

import "fmt"

func main() {

	var number int

	fmt.Println("Please enter number: ")
	fmt.Scan(&number)

	fmt.Printf("The number %d in binary = %b", number, number)
}

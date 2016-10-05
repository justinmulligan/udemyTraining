package main

import "fmt"

// Global/Package level variables
var pi = 3.16
var name = "Justin"

func main() {
	// Block level variables
	a := 3

	fmt.Printf("%v \n", a)
	fmt.Println("Hello", name)
	fmt.Println("Pi =", pi)
	fmt.Println("Pi x 2 =", pi*2)
	fmt.Println("Pi + 5 =", pi+5)
	fmt.Println("Pi / 2 =", pi/2)

	// Calling sub routine called submain
	fmt.Printf("\n")
	fmt.Println("***Calling sub routine***")
	submain()

}

func submain() {
	// Block level variables
	a := 2.2

	fmt.Println("Hello", name, "You are", a*pi)
}

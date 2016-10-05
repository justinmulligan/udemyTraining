package main

import	"fmt"

func main() {
	a := 10
	b := "Justin"
	c := 60.2
	d := true
	name := "Justin "

// The %v will just print out the vaiable and work out what type it is
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)


// The %T will print out what type of value the is assogned to the variable
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

// Print my name
	fmt.Println("Hello", name)

}

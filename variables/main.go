package main

import (
	"fmt"
	"reflect"
)

func main() {

	//  Declaration with initialization
	/*
		If you know beforehand what a variable's value will be,
		you can declare variables and assign them values on the same line.
	*/
	{
		fmt.Println("Declaration with initialization")

		var i int = 10
		var s string = "Canada"

		fmt.Println(i)
		fmt.Println(s)

	}

	// Declaration without initialization
	/*
		The keyword var is used for declaring variables followed by
		the desired name and the type of value the variable will hold.
	*/
	{
		fmt.Println("\nDeclaration without initialization")

		var i int
		var s string

		i = 10
		s = "Canada"

		fmt.Println(i)
		fmt.Println(s)
	}

	// Declaration with type inference
	/*
		You can omit the variable type from the declaration,
		when you are assigning a value to a variable at the time of declaration.
		The type of the value assigned to the variable will be used as the type of that variable.
	*/
	{
		fmt.Println("\nDeclaration with type inference")

		var i = 10
		var s = "Canada"

		fmt.Println(reflect.TypeOf(i))
		fmt.Println(reflect.TypeOf(s))
	}

	// Declaration of multiple variables
	/*
		Golang allows you to assign values to multiple variables in one line.
	*/
	{
		fmt.Println("\nDeclaration of multiple variables")

		var fname, lname string = "John", "Doe"
		m, n, o := 1, 2, 3
		item, price := "Mobile", 2000

		fmt.Println(fname + lname)
		fmt.Println(m + n + o)
		fmt.Println(item, "-", price)
	}

	// Short Variable Declaration in Golang
	/*
		The := short variable assignment operator indicates that short variable declaration is being used.
		There is no need to use the var keyword or declare the variable type.
	*/
	{
		fmt.Println("\nShort Variable Declaration in Golang")

		name := "John Doe"
		fmt.Println(reflect.TypeOf(name))
	}

	// Zero Values
	/*
		If you declare a variable without assigning it a value,
		Golang will automatically bind a default value (or a zero-value) to the variable.
	*/
	{
		fmt.Println("\nZero Values")

		var quantity float32
		fmt.Println(quantity)

		var price int16
		fmt.Println(price)

		var product string
		fmt.Println(product)

		var inStock bool
		fmt.Println(inStock)
	}

	// Golang Variable Declaration Block
	/*
		Variables declaration can be grouped together into blocks for greater readability and code quality.
	*/
	{
		fmt.Println("\nGolang Variable Declaration Block")

		var (
			product  = "Mobile"
			quantity = 50
			price    = 50.50
			inStock  = true
		)

		fmt.Println(quantity)
		fmt.Println(price)
		fmt.Println(product)
		fmt.Println(inStock)
	}

}

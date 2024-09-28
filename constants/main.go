package main

import "fmt"

const PRODUCT string = "Canada"
const PRICE = 500

const (
	QUANTITY = 50
	STOCK    = true
)

const (
	a = iota
	b
	c
)

const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {

	fmt.Println(QUANTITY)
	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(STOCK)

	fmt.Println("\niota")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	filesize := 4000000.00
	fmt.Printf("\n%.2fGB\n", filesize/GB)

}

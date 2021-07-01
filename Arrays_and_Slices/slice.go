package main

import (
	"fmt"
	"reflect"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func itemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func main() {

	/* Create Empty Slice */
	{
		fmt.Printf("\nCreate Empty Slice\n")
		var intSlice []int
		var strSlice []string

		fmt.Println(reflect.ValueOf(intSlice).Kind())
		fmt.Println(reflect.ValueOf(strSlice).Kind())
	}

	/* Declare Slice using Make */
	{
		fmt.Printf("\nDeclare Slice using Make\n")

		var intSlice = make([]int, 10)        // when length and capacity is same
		var strSlice = make([]string, 10, 20) // when length and capacity is different

		fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
		fmt.Println(reflect.ValueOf(intSlice).Kind())

		fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
		fmt.Println(reflect.ValueOf(strSlice).Kind())
	}

	/* Initialize Slice with values using a Slice Literal */
	{
		fmt.Printf("\nInitialize Slice with values using a Slice Literal\n")

		var intSlice = []int{10, 20, 30, 40}
		var strSlice = []string{"India", "Canada", "Japan"}

		fmt.Printf("%v intSlice \tLen: %v \tCap: %v\n", intSlice, len(intSlice), cap(intSlice))
		fmt.Printf("%v strSlice \tLen: %v \tCap: %v\n", strSlice, len(strSlice), cap(strSlice))
	}

	/* Remove Item from Slice */
	{
		fmt.Printf("\nRemove Item from Slice\n")
		var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
		fmt.Println(strSlice)

		strSlice = RemoveIndex(strSlice, 3)
		fmt.Println(strSlice)
	}

	/* Copy a Slice */
	{
		fmt.Printf("\nCopy a Slice\n")

		a := []int{5, 6, 7} // Create a slice
		fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))

		b := make([]int, 5, 10) // Create a bigger slice
		fmt.Println("Slice B before copying:", b)
		fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

		copy(b, a) // Copy function

		fmt.Println("Slice B after copying:", b)
		fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))

		b[3] = 8
		b[4] = 9
		fmt.Println("Slice B after adding elements:", b)
	}

	/* Loop Through a Slice */
	{
		fmt.Printf("\nLoop Through a Slice\n")

		var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}

		fmt.Println("\n---Example 1 ---\n")
		for index, element := range strSlice {
			fmt.Println(index, "--", element)
		}

		fmt.Println("\n---Example 2 ---\n")
		for _, value := range strSlice {
			fmt.Println(value)
		}

		j := 0
		fmt.Println("\n---Example 3 ---\n")
		for range strSlice {
			fmt.Println(strSlice[j])
			j++
		}
	}

	/* Check if Item Exists */
	{
		var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
		fmt.Println(itemExists(strSlice, "Canada"))
		fmt.Println(itemExists(strSlice, "Africa"))
	}

}

package main

import (
	"fmt"
	"reflect"
)

/* Note: In Golang it is not recommended to use Pointer to an Array as an Argument
to Function as the code become difficult to read. Also, it is not considered a good way
to achieve this concept. To achieve this you can use slice instead of passing pointers. */
func UpdateArray(a *[3]string) {
	a[0] = "Zero@Array"
}

func UpdateSlice(s []string) {
	s[0] = "Zero@Slice"
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		fmt.Printf("Invalid data-type")
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func main() {

	/* Initializing an Array with Array Literal & ellipses... */
	{
		fmt.Printf("\n\nInitializing an Array\n")

		array1 := [3]int{1, 2, 3}
		array2 := [...]int{1, 2, 3}

		fmt.Printf("array1 : %v Type : %T\n", array1, array1)
		fmt.Printf("array2 : %v Type : %T\n", array2, array2[0])
	}

	/* Pointer to an Array as Function Argument */
	{
		fmt.Printf("\n\nPointer to an Array as Function Argument\n")

		var array1 [3]string
		fmt.Printf("array1: %v\n", array1)
		array1[0] = "one"
		array1[1] = "two"
		array1[2] = "three"

		fmt.Printf("array1 : %v\n", array1)

		UpdateArray(&array1)
		fmt.Printf("array1 : %v\n", array1)
		UpdateSlice(array1[:])
		/* Array1 still be an array */
		fmt.Printf("array1 : %v Type : %T\n", array1, array1)

		/* Array2 is Slice*/
		array2 := array1[:]
		fmt.Printf("array2 : %v Type : %T\n", array2, array2)
	}

	/* Loop Through an Indexed Array */
	{
		fmt.Printf("\n\nLoop Through an Indexed Array\n")

		intArray := [5]int{10, 20, 30, 40, 50}

		fmt.Printf("\n---Example 1---\n")
		for i := 0; i < len(intArray); i++ {
			fmt.Println(intArray[i])
		}

		fmt.Printf("---Example 2---\n")
		for index, element := range intArray {
			fmt.Println(index, "=>", element)

		}

		fmt.Printf("---Example 3---\n")
		for _, value := range intArray {
			fmt.Println(value)
		}

		j := 0
		fmt.Printf("---Example 4---\n")
		for range intArray {
			fmt.Println(intArray[j])
			j++
		}

	}

	/* Copy Array */
	{
		fmt.Printf("\n\nCopy Array\n")

		strArray1 := [3]string{"Japan", "Australia", "Germany"}

		strArray2 := strArray1  // data is passed by value
		strArray3 := &strArray1 // data is passed by refrence

		fmt.Printf("strArray1: %v\n", strArray1)
		fmt.Printf("strArray2: %v\n", strArray2)

		strArray1[0] = "Canada"

		fmt.Printf("strArray1: %v\n", strArray1)
		fmt.Printf("strArray2: %v\n", strArray2)
		fmt.Printf("*strArray3: %v\n", *strArray3)
	}

	/* Check if Element Exists */
	{
		fmt.Printf("\n\nCheck if Element Exists\n")
		strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
		fmt.Println(itemExists(strArray, "Canada"))
		fmt.Println(itemExists(strArray, "Africa"))
	}
}

package main

import "fmt"

func binarySearch(target int, list []int, low int, high int) bool {

	var mid int = (high + low) / 2

	if high >= low {
		if list[mid] == target {
			return true
		} else if list[mid] > target {
			return binarySearch(target, list, low, mid-1)
		} else {
			return binarySearch(target, list, mid+1, high)
		}
	} else {
		return false
	}
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(1, items, 0, (len(items) - 1)))
}

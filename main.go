package main

import (
	search "data-struct/search"
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	searchValue := 4

	binarySearchRecursive := execBinarySearchRecursive(list, searchValue)
	binarySearch := execBinarySearch(list, searchValue)

	if binarySearchRecursive == -1 || binarySearch == -1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(list[binarySearchRecursive])
	fmt.Println(list[binarySearch])

}

func execBinarySearch(list []int, searchValue int) int {
	position, exists := search.BinarySearch(list, searchValue)

	if !exists {
		fmt.Println("\nNot Found >:0")
		return -1
	}

	fmt.Println("Found at position: ", position)
	return position
}

func execBinarySearchRecursive(list []int, searchValue int) int {
	low, high := 0, len(list)-1
	position, exists := search.BinarySearchRecursive(list, searchValue, low, high)

	if !exists {
		fmt.Println("\nNot Found >:0")
		return -1
	}

	fmt.Println("Found at position: ", position)
	return position
}

package main

import "fmt"

func main() {
	// Declaring and initializing an array
	numbers := [5]int{1, 2, 3, 4, 5}

	// Printing the array
	fmt.Println("Array:", numbers)

	// Accessing individual elements of the array
	fmt.Println("Element at index 2:", numbers[2])

	// Modifying an element of the array
	numbers[2] = 10
	fmt.Println("Modified array:", numbers)

	// Slicing the array
	slicedArray := numbers[1:4]
	fmt.Println("Sliced array:", slicedArray)

	// Appending to a slice
	slicedArray = append(slicedArray, 6, 7, 8)
	fmt.Println("Appended slice:", slicedArray)

	// Length and capacity of the slice
	fmt.Println("Length of slice:", len(slicedArray))
	fmt.Println("Capacity of slice:", cap(slicedArray))
}

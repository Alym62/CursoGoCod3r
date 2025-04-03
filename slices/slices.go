package main

import "fmt"

// Function to add elements in slice using - append()
func addFiveElementsInSlice(slice []int) []int {
	slice = append(slice, 1, 2, 3, 4, 5)
	return slice
}

func copyElementsSlice(slice []int) (newSlice []int) {
	newSlice = make([]int, 2)
	copy(newSlice, slice)

	return newSlice
}

func main() {
	arrayOne := [3]int{1, 2, 3} // Array fixed
	slice := []int{1, 2, 3}     // Slice with height variable

	fmt.Println(arrayOne, slice)

	array := [5]int{1, 2, 3, 4, 5}

	// Slice no is array!! Slice is a part of array
	sliceArr := array[1:3]
	fmt.Println(array, sliceArr)

	sliceArrTwo := array[:2]
	fmt.Println(sliceArrTwo)

	sliceArrThree := sliceArrTwo[:1]
	fmt.Println(sliceArrThree)

	sliceWithMake := make([]int, 10)
	sliceWithMake[9] = 10
	fmt.Println(sliceWithMake)

	sleiceWithMakeAndArr := make([]int, 10, 20)
	fmt.Println(sleiceWithMakeAndArr, len(sleiceWithMakeAndArr), cap(sleiceWithMakeAndArr))

	sleiceWithMakeAndArr = append(sleiceWithMakeAndArr, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sleiceWithMakeAndArr, len(sleiceWithMakeAndArr), cap(sleiceWithMakeAndArr))

	// Access intern array
	newSlice := make([]int, 10, 20)
	newSliceTwo := append(newSlice, 1, 2, 3)

	fmt.Println(newSlice)
	fmt.Println(newSliceTwo)

	newSlice[0] = 10
	fmt.Println(newSlice)
	fmt.Println(newSliceTwo)

	// Understand methods of copy and append
	var sliceToAdd []int
	sliceToAdd = addFiveElementsInSlice(sliceToAdd)

	sliceCopy := copyElementsSlice(sliceToAdd)
	fmt.Println(sliceCopy)
}

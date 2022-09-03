package main

import "fmt"

func addNum(slice []int) []int {
	slice = append(slice, 4)

	// fmt.Println("addNum : ", *slice)
	return slice
}

func pointerAddNum(slice *[]int) {
	*slice = append(*slice, 4)

	// fmt.Println("addNum : ", *slice)
	// return slice
}

func main() {

	slice := []int{1, 2, 3}

	fmt.Println("first slice : ", slice)

	slice = addNum(slice)
	// pointerAddNum(&slice)

	fmt.Println("addNum slice : ", slice)
	// fmt.Println("pointerAddNum slice : ", slice)
}

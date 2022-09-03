package main

import "fmt"

// func changeArray(array2 [5]int) {
// 	array2[2] = 200
// }

// func changeSlice(slice2 []int) {
// 	slice2[2] = 200
// }

func main() {

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// slice1[1] = 100

	slice1 = append(slice1, 500)
	fmt.Println(slice1)
	fmt.Println(slice2)

	// fmt.Println(slice1)
	// fmt.Println(slice2)

	// fmt.Println(slice1[cap(slice1)-1])
	// fmt.Println(slice2[len(slice2)-1])
	// fmt.Println(slice2[cap(slice2)])

	// array := [5]int{1, 2, 3, 4, 5}
	// slice := []int{1, 2, 3, 4, 5}

	// changeArray(array)
	// changeSlice(slice)

	// fmt.Println(array)
	// fmt.Println(slice)

	// var slice2 = make([]int, 3, 5)

	// fmt.Println(slice2)

	// slice := []int{1, 2, 3}

	// slice2 := append(slice, 4)

	// fmt.Println(slice)
	// fmt.Println(slice2)

	// var sumSlice = []int{1, 2, 3}
	// var multiplySlice = []int{1, 2, 3}

	// for i := 0; i < len(sumSlice); i++ {
	// 	sumSlice[i] += 10
	// }

	// fmt.Println(sumSlice)

	// for i := 0; i < len(multiplySlice); i++ {
	// 	multiplySlice[i] = i * 2
	// }
	// for i, v := range multiplySlice {
	// 	multiplySlice[i] = v * 2
	// }
	// fmt.Println(multiplySlice)

	// var makeSlice = make([]int, 3) // 초기화 방법1
	// makeSlice2 := []int{0, 0, 0}   // 초기화 방법2

	// if len(slice) == 0 {
	// 	fmt.Println("slice is empty", slice)
	// }

	// slice[1] = 10
	// fmt.Println(slice)

	// fmt.Println(makeSlice)
	// fmt.Println(makeSlice2)

}
